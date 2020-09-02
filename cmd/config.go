package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get all configs",
	Long:  `Get all configs`,
	Run: func(cmd *cobra.Command, args []string) {
		c := viper.AllSettings()
		bs, err := yaml.Marshal(c)
		if err != nil {
			log.Fatalf("Unable to marshal config to YAML: %v", err)
		}
		fmt.Println(string(bs))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
