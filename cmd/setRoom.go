package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setRoomCmd represents the setRoom command
var setRoomCmd = &cobra.Command{
	Use:     "set-room",
	Short:   "Set roomID (channelID) of chatwork to send message ",
	Long:    `Specify roomID to send message to that room`,
	Aliases: []string{"sr"},
	Example: "nm5 sr 195481599",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) == 0 {
			fmt.Printf("Invalid number of arguments. Expect: 1, Given: %v\n", len(args))
			return
		}

		trimRoomID := strings.Trim(args[0], " ")
		viper.Set("roomID", trimRoomID)
		viper.WriteConfig()
		fmt.Printf("Set room successfully! Room ID: %v\n", trimRoomID)
	},
}

func init() {
	rootCmd.AddCommand(setRoomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setRoomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setRoomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
