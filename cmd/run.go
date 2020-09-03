package cmd

import (
	"errors"
	"fmt"
	"log"

	request "nm5/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Send message",
	Long:  `Send message to channel on chatwork, you must set cookie and set token for it to work`,
	Run: func(cmd *cobra.Command, args []string) {
		if !viper.IsSet("token") || !viper.IsSet("cookie") {
			log.Fatalln("Token or Cookie is not set!")
		}
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

		request.SendMessage(result)

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
