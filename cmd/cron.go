package cmd

import (
	"fmt"
	"log"
	"nm5/utils/cli"
	request "nm5/utils/request"
	"os"
	"strings"
	"sync"

	"github.com/manifoldco/promptui"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func worker(message string, wg *sync.WaitGroup) {
	request.SendMessage(message)
	defer wg.Done()
}

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Set a cron job to send report at 16:47",
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

		var wg sync.WaitGroup
		wg.Add(1)
		c := cron.New()
		// c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 30 16 * * *", func() { worker(&wg) })
		c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 47 16 * * *", func() { worker(message, &wg) })
		c.Start()
		fmt.Println("Cron job running...")
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
