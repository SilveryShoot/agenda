// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)

// auCmd represents the au command
var auCmd = &cobra.Command{
	Use:   "au",
	Short: "add a user to a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Add User called")
		users, _ := cmd.Flags().GetStringSlice("user")
		title, _ := cmd.Flags().GetString("title")
		if len(users) == 0 || title == "" {
			fmt.Println("Please input title and user(s)(input like \"name1,name2\")")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login")
		} else {
			flag := service.AddMeetingParticipator(user.Name, title, users)
			if flag != true {
				fmt.Println("Unexpected error. Check error.log for detail")
			} else {
				fmt.Println("Successfully add")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(auCmd)
	auCmd.Flags().StringSliceP("user", "u", nil, "user(s) you want to add, input like \"name1,name2\"")
	auCmd.Flags().StringP("title", "t", "", "the title of meeting")
}
