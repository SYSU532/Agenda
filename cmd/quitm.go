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
	"bufio"
	"fmt"
	"os"

	"github.com/SYSU532/agenda/entity"
	"github.com/spf13/cobra"
)

var quitmTitle string

// quitmCmd represents the quitm command
var quitmCmd = &cobra.Command{
	Use:   "quitm",
	Short: "Quit a meeting",
	Long: fmt.Sprintf(`Use this command to quit a meeting using a already logged in user.

Usage: %v quitm [-t title]`, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		userinfo, err := entity.GetCurrentUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to cancel meeting: %v\n", err)
			return
		}
		reader := bufio.NewReader(os.Stdin)
		if quitmTitle == "" {
			fmt.Printf("Enter the meeting title: ")
			title, _ := reader.ReadString('\n')
			quitmTitle = title[0 : len(title)-1]
		}
		err = entity.QuitMeeting(quitmTitle, userinfo.Username)
		if err == nil {
			fmt.Printf("Successfully quited the meeting: %v\n", quitmTitle)
		} else {
			fmt.Fprintf(os.Stderr, "Fail to quit the meeting: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(quitmCmd)

	quitmCmd.Flags().StringVarP(&quitmTitle, "title", "t", "", "The title of the meeting you want to quit")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
