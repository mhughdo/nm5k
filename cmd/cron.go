package cmd

import (
	"errors"
	"fmt"
	"log"
	request "no-more-5k/utils"
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		var wg sync.WaitGroup
		wg.Add(1)
		c := cron.New()
		// c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 30 16 * * *", func() { worker(&wg) })
		c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 55 16 * * *", func() { worker(result, &wg) })
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
