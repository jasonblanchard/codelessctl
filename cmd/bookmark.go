package cmd

import (
	"fmt"

	"github.com/jasonblanchard/thecodelessctl/pkg/store"
	"github.com/spf13/cobra"
)

// bookmarkCmd represents the bookmark command
var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Show current bookmark value",
	Run: func(cmd *cobra.Command, args []string) {
		id := store.GetBookmark()
		fmt.Println(id)
	},
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookmarkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookmarkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
