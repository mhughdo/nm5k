package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiURL string = "https://www.chatwork.com/gateway/send_chat.php"
var roomID string = "195722902"

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var defaultMessage = "[To:4001758]Le Tuan Hiep (nick chính thức) \\n Today plan: Làm task trong sprint 3 \\n Tomorrow plan: tiếp tục làm sprint 3"

		validate := func(message string) error {
			if len(message) < 3 {
				return errors.New("Message must have more than 3 characters")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "message",
			Validate: validate,
			Default:  defaultMessage,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		baseURL, err := url.Parse(apiURL)

		if err != nil {
			fmt.Println("Malformed URL: ", err.Error())
			return
		}

		params := url.Values{}
		params.Add("room_id", roomID)

		baseURL.RawQuery = params.Encode()

		var replacedString = strings.ReplaceAll(result, "\\n", "\n")
		jsonValues, err := json.Marshal(map[string]string{
			"text": replacedString,
			"_t":   "32b56ad427eeb67111a4cf8147f2643d899b738b5f4c667575738",
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

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// response, err := http.PostForm(baseURL.String(), formData)

// if err != nil {
// 	fmt.Printf("Error sending message %v\n", err)
// 	return
// }

// defer response.Body.Close()
// body, err := ioutil.ReadAll(response.Body)

// if err != nil {
// 	fmt.Printf("Error parsing response %v\n", err)
// 	return
// }

// fmt.Printf("%s\n", string(body))

// var finalRes map[string]interface{}

// json.NewDecoder(response.Body).Decode(&finalRes)

// log.Println(finalRes)
