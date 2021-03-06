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

// rCmd represents the r command
var rCmd = &cobra.Command{
	Use:   "r",
	Short: "register",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Register called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("cellphone")
		if username == "" || password == "" || email == "" || phone == "" {
			fmt.Println("Error: some information has not been input")
			return
		}
		pass, err := service.UserRegister(username, password, email, phone)
		if pass == false {
			fmt.Println("Username existed!")
			return
		} else {
			if err != nil {
				fmt.Println("Error! Please read error.log for detail")
				return
			} else {
				fmt.Println("Successfully register!")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rCmd)
	rCmd.Flags().StringP("username", "u", "", "username")
	rCmd.Flags().StringP("password", "p", "", "password")
	rCmd.Flags().StringP("email", "e", "", "email")
	rCmd.Flags().StringP("cellphone", "c", "", "cellphone")
}
