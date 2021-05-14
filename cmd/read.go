package cmd

import (
	"fmt"
	"strconv"

	"github.com/jasonblanchard/thecodelessctl/pkg/codeless"
	"github.com/jasonblanchard/thecodelessctl/pkg/store"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a single story at stdout",
	Long: `Read a single story at stdout

For a more enjoyable reading experience, pipe it to less:

	thecodelessctl read 42 | less
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id int
		var err error
		bookmark, err := cmd.Flags().GetBool("bookmark")
		config, err := cmd.Flags().GetString("config")

		if err != nil {
			return err
		}

		if len(args) == 0 {
			id = store.GetBookmark()
		} else {
			idString := args[0]
			id, err = strconv.Atoi(idString)

			if err != nil {
				return err
			}
		}

		if bookmark == true {
			err := store.WriteBookmark(config, id)
			if err != nil {
				return err
			}
		}

		story := codeless.GetStoryById(id)
		fmt.Println(codeless.DecorateStory(story))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	readCmd.Flags().BoolP("bookmark", "b", false, "Set bookmark value to this story")
}
