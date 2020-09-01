package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCookieCmd represents the setCookie command
var setCookieCmd = &cobra.Command{
	Use:   "set-cookie",
	Short: "Set cookie to use when send-chat API is called",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) == 0 {
			fmt.Println("Invalid number of arguments. Expect: 1")
			return
		}
		viper.Set("cookie", args[0])
		viper.WriteConfig()
		fmt.Println(viper.AllSettings())
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
