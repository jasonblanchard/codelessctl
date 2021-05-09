/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jasonblanchard/thecodelessctl/codeless"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "List all stories and case numbers",
	Long: `For a more enjoyably readin experience, pipe it to less:

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
