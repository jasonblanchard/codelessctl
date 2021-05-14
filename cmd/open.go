package cmd

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/jasonblanchard/thecodelessctl/pkg/store"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open story in the browser at http://thecodelesscode.com",
	Long: `Open story in the browser at http://thecodelesscode.com
	
Pass argument to open a specific story:

	thecodelessctl open 42

Or omit argument to use bookmark key in config:

	thecodelessctl open
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id int
		var err error

		if len(args) == 0 {
			id = store.GetBookmark()
		} else {
			idString := args[0]
			id, err = strconv.Atoi(idString)

			if err != nil {
				return err
			}
		}

		url := fmt.Sprintf("http://thecodelesscode.com/case/%v", id)

		err = exec.Command("open", url).Start()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
