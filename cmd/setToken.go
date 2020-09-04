package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setTokenCmd represents the setToken command
var setTokenCmd = &cobra.Command{
	Use:     "set-token [token]",
	Short:   "Set token in config file to use when send-chat API is called.",
	Aliases: []string{"st"},
	Example: "nm5 st c2aef82685a644d2bdecfc9357bda0cb7c8ce2905f519a47420f2",
	Long: `Open dev tools on chatwork.com site, run the script code below to get token and run nm5 set-token [token]:
			let scripts = document.getElementsByTagName('script')
			for (const s of scripts) {
     		if (s.innerText.includes('ACCESS_TOKEN')) {
         for (const line of s.text.split('\n')) {
             if (line.includes('ACCESS_TOKEN')) {
                console.log(line.match(/\'(.*)\'/)[1])
            }

    			}
    		}
			}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) == 0 {
			fmt.Println("Invalid number of arguments. Expect: 1")
			return
		}

		trimToken := strings.Trim(args[0], " ")
		viper.Set("token", trimToken)
		viper.WriteConfig()
		fmt.Printf("Set token successfully! Token: %v\n", trimToken)
	},
}

func init() {
	rootCmd.AddCommand(setTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
