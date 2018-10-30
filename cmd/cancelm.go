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
	"github.com/SYSU532/agenda/Log"
	"github.com/spf13/cobra"
)

var cancelmTitle string

// cancelmCmd represents the cancelm command
var cancelmCmd = &cobra.Command{
	Use:   "cancelm",
	Short: "Cancel a meeting",
	Long: fmt.Sprintf(`Use this command to cancel a meeting using a already logged in user.

Usage: %v cancelm [-t title]`, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		Log.WriteLog("Invoke cancel meeting command to stop the meetings you create", 1)
		userinfo, err := entity.GetCurrentUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to cancel meeting: %v\n", err)
			Log.WriteLog("Error when geeting current user, maybe you are not logged in", 0)
			return
		}
		reader := bufio.NewReader(os.Stdin)
		if cancelmTitle == "" {
			fmt.Printf("Enter the meeting title: ")
			title, _ := reader.ReadString('\n')
			cancelmTitle = title[0 : len(title)-1]
		}

		err = entity.CancelMeeting(cancelmTitle, userinfo.Username)
		if err == nil {
			fmt.Println("Successfully canceled meeting")
			Log.WriteLog(fmt.Sprintf("Successfully canceled meeting %v", cancelmTitle), 1)
		} else {
			fmt.Fprintf(os.Stderr, "Fail to cancel meeting: %v", err)
			Log.WriteLog(fmt.Sprintf("Fail to cancel meeting: %v", err), 0)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(cancelmCmd)

	cancelmCmd.Flags().StringVarP(&cancelmTitle, "title", "t", "", "The title of the meeting to be canceled")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
