package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config file",
	Long: `Initialize config file
	
Noop if file already exists`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		if fileExists(config) {
			fmt.Println(fmt.Sprintf("Config file already exists at %s", config))
			return nil
		}

		viper.GetViper().Set("bookmark", 1)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("Config file initialized at %s", config))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
