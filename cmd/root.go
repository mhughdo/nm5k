package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var configName string = ".no-more-5k"
var configType string = "yaml"
var defaultMessage string
var roomID string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nm5",
	Short: "5k 5k 5k",
	Long:  `nm5 is a tool that make sure you wont lose any money anymore`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.no-more-5k.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".no-more-5k" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.OnConfigChange(func(e fsnotify.Event) {
			// fmt.Println("Config file changed")
		})
		viper.WatchConfig()

		var configPath string = home + "/" + configName + "." + configType
		if err := viper.SafeWriteConfigAs(configPath); err != nil {
			if os.IsNotExist(err) {
				err = viper.WriteConfigAs(configPath)
			}
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		defaultMessage = "[To:4001758]Le Tuan Hiep (nick chính thức)\nToday plan: Làm task trong sprint 4\nTomorrow plan: Tiếp tục làm các task trong sprint 4"
		roomID = "195481599"

		if !viper.IsSet("message") {
			viper.Set("message", defaultMessage)
			viper.WriteConfig()
		} else {
			defaultMessage = viper.GetString("message")
		}

		if !viper.IsSet("roomID") {
			viper.Set("roomID", roomID)
			viper.WriteConfig()
		} else {
			roomID = viper.GetString("roomID")
		}
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
