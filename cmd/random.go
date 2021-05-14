package cmd

import (
	"fmt"

	"github.com/jasonblanchard/thecodelessctl/pkg/codeless"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Display a random story",
	Run: func(cmd *cobra.Command, args []string) {
		story := codeless.GetRandomStory()
		fmt.Println(codeless.DecorateStory(story))
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
