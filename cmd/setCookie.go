package cmd

import (
	"fmt"
	"nm5/utils/request"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCookieCmd represents the setCookie command
var setCookieCmd = &cobra.Command{
	Use:     "set-cookie [cookie]",
	Short:   "Set cookie to use when send-chat API is called",
	Aliases: []string{"sc"},
	Example: "nm5 set-cookie up8ri7rfmqoabgpa829efi3q90",
	Long:    "Open dev tools, click on tab Application (Chrome) or Storage(firefox) > Cookies, copy cookie with key named: cwssid, then use set-cookie [cookie] to set cookie. ex: nm5 set-cookie n9sse6jqobe91bp7um7jn7j21c;",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) == 0 {
			fmt.Printf("Invalid number of arguments. Expect: 1, Given: %v\n", len(args))
			return
		}

		trimCookie := strings.Trim(args[0], " ")
		viper.Set("cookie", trimCookie)
		viper.WriteConfig()
		fmt.Printf("Set cookie successfully! Cookie: %v\n", trimCookie)
		fmt.Println("Getting token...")
		token := request.GetToken()
		if token != "" {
			viper.Set("token", token)
			viper.WriteConfig()
			fmt.Printf("Set token successfully! Token: %v\n", token)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCookieCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCookieCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCookieCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
