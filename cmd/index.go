package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jasonblanchard/thecodelessctl/pkg/codeless"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "List all stories and case numbers",
	Long: `List all stories and case numbers
	
For a more enjoyable reading experience, pipe it to less:

	thecodelessctl index | less
`,
	Run: func(cmd *cobra.Command, args []string) {
		stories := codeless.GetAllStories()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		for i := 0; i < len(stories); i++ {
			story := stories[i]
			row := fmt.Sprintf("%s\t%s\t", story.CaseNumber, story.Title)
			fmt.Fprintln(w, row)
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
