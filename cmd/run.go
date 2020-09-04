package cmd

import (
	"fmt"
	"log"
	"nm5/utils/cli"
	request "nm5/utils/request"
	"os"
	"strings"

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

		var message string

		prompt := promptui.Select{
			Label: "message",
			Items: []string{strings.ReplaceAll(defaultMessage, "\n", "\\n"), "Custom"},
		}

		index, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if index == 1 {
			// var configPath string = home + "/" + configName + "." + configType
			messageBytes, err := cli.CaptureInputFromEditor(defaultMessage)
			message = string(messageBytes)

			if err != nil {
				log.Fatalln("Error editing file", err)
			}

			viper.Set("message", message)
			viper.WriteConfig()
			if message == "" {
				fmt.Println("Message cannot be empty")
				os.Exit(1)
			}
			fmt.Printf("Message: %v\n", message)
		} else {
			message = result
		}

		request.SendMessage(message)

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
