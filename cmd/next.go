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
	"strconv"

	"github.com/jasonblanchard/thecodelessctl/codeless"
	"github.com/jasonblanchard/thecodelessctl/store"
	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Read the next story",
	Long: `Reads the bookmark in your config file, finds the next story, then displays it and sets that as your current bookmark.

You can page through each story by running this command multiple times.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id int
		var err error
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

		nextStoryId := codeless.GetNextStoryId(id)
		store.WriteBookmark(config, nextStoryId)

		story := codeless.GetStoryById(nextStoryId)
		fmt.Println(codeless.DecorateStory(story))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
