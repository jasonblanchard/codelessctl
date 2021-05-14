package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// Version updated by ldflags
var Version = "development"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thecodelessctl",
	Short: "A CLI reader for The Codeless Code :: Fables and Koans for the Software Engineer https://thecodelesscode.com/",
	Long: `Set up to enable bookmark state:

	thecodelessctl init

This will create a config file at your config path (default is $HOME/.thecodelessctl.yaml).

Start reading with:

	thecodelessctl read

Then read through each story using your bookmark state by calling next:

	thecodelessctl next

Pipe story output to less for a nicer reading experience:

	thecodelessctl read 42 | less`,
	Version: Version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", fmt.Sprintf("%s/.thecodelessctl.yaml", home), "config file (default is $HOME/.thecodelessctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".thecodelessctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".thecodelessctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
