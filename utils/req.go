package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

var apiURL string = "https://www.chatwork.com/gateway/send_chat.php"
var roomID string = "195481599"

// SendMessage send message to specific channel on chatwork
func SendMessage(message string) {
	baseURL, err := url.Parse(apiURL)

	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	params := url.Values{}
	params.Add("room_id", roomID)

	baseURL.RawQuery = params.Encode()

	var replacedString = strings.ReplaceAll(message, "\\n", "\n")
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

	fmt.Println(formData.Encode())

	req, err := http.NewRequest(
		http.MethodPost,
		baseURL.String(),
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("cookie", viper.GetString("cookie"))

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln("Error making request")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error parsing response %v\n", err)
		return
	}

	fmt.Printf("%s\n", string(body))
}
