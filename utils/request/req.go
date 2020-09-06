package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
)

var homePage string = "https://www.chatwork.com"
var apiURL string = homePage + "/gateway/send_chat.php"
var cookieName string = "cwssid"

// var roomID string = "195722902"

// SendMessage send message to specific channel on chatwork
func SendMessage(message string) {
	if message == "" {
		fmt.Println("Message cannot be empty")
		os.Exit(1)
	}

	type Status struct {
		Success bool
		Message interface{}
	}

	type Result struct {
		NewMessageID string `json:"new_message_id"`
	}
	type Response struct {
		Status Status `json:"status"`
		Result interface{}
	}

	baseURL, err := url.Parse(apiURL)

	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	roomID := viper.GetString("roomID")
	params := url.Values{}
	params.Add("room_id", roomID)

	baseURL.RawQuery = params.Encode()

	var replacedString = strings.ReplaceAll(message, "\\n", "\n")
	// re := regexp.MustCompile(`\r?\n`)
	// replacedString := re.ReplaceAllString(message, `\n`)
	// fmt.Printf("fd;fhkhfkhf %v", replacedString)
	jsonValues, err := json.Marshal(map[string]string{
		"text": replacedString,
		"_t":   viper.GetString("token"),
	})

	if err != nil {
		log.Fatalln(err)
	}

	formData := url.Values{
		"pdata": {string(jsonValues)},
	}

	req, err := http.NewRequest(
		http.MethodPost,
		baseURL.String(),
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("cookie", fmt.Sprintf("%v=%v", cookieName, viper.GetString("cookie")))

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making send-message request, %v\n", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error parsing response %v\n", err)
		return
	}

	var parsedBody Response
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		log.Fatalln("Error parsing response", err)
	}

	if parsedBody.Status.Success != true {
		var errMsg string
		switch v := parsedBody.Status.Message.(type) {
		case string:
			errMsg = v
		case []interface{}:
			// for _, val := range v {
			// 	fmt.Println(val)
			// }
			errMsg = v[0].(string)
		}
		fmt.Printf("❌  Send report failed: %v\n", errMsg)

	} else {
		fmt.Println("✅  Send report successfully")
	}
	// fmt.Printf("%s\n", string(body))

}

// GetToken get token from chatwork and set in config file
func GetToken() string {
	if !viper.IsSet("cookie") {
		log.Fatalln("Cookie not found, run nm5 sc [cookie] to set cookie")
	}
	var token string
	cookie := fmt.Sprintf("%v=%v", cookieName, viper.GetString("cookie"))

	req, err := http.NewRequest(http.MethodGet, homePage, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("cookie", cookie)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making get-token request, %v\n", err)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	html, err := doc.Html()

	if !strings.Contains(html, "ACCESS_TOKEN") {
		log.Fatalln("Error getting token: Invalid or Expired cookie")
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		tagText := s.Text()

		if strings.Contains(tagText, "ACCESS_TOKEN") {
			tagTextArr := strings.Split(tagText, "\n")
			for _, line := range tagTextArr {
				if strings.Contains(line, "ACCESS_TOKEN") {
					re := regexp.MustCompile(`\'(.*)\'`)
					matches := re.FindStringSubmatch(line)
					token = matches[1]
					return
				}
			}
		}
	})

	return token
}
