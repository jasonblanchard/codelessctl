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

	"github.com/jasonblanchard/thecodelessctl/store"
	"github.com/spf13/cobra"
)

// setBookmarkCmd represents the setBookmark command
var setBookmarkCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current bookmark value",
	Args:  cobra.ExactArgs(1),
	Long: `Set the bookmark in your config to story 42:

	thecodelessctl bookmark set 42
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cmd.Flags().GetString("config")
		idString := args[0]
		id, err := strconv.Atoi(idString)
		err = store.WriteBookmark(config, id)

		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("Bookmark set to %v in %v", id, config))

		return nil
	},
}

func init() {
	bookmarkCmd.AddCommand(setBookmarkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setBookmarkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setBookmarkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
