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
	"os"

	"github.com/SYSU532/agenda/entity"
	"github.com/SYSU532/agenda/Log"
	"github.com/spf13/cobra"
)

// clearmCmd represents the clearm command
var clearmCmd = &cobra.Command{
	Use:   "clearm",
	Short: "Clear all meetings",
	Long: fmt.Sprintf(`Use this command to clear all meetings using a already logged in user.

Usage: %v clearm `, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		Log.WriteLog("Invoke clear meeting command to stop all the meetings you create", 1)
		userinfo, err := entity.GetCurrentUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to cancel meeting: %v\n", err)
			Log.WriteLog("Error when geeting current user, maybe you are not logged in", 0)
			return
		}
		err = entity.ClearMeeting(userinfo.Username)
		if err == nil {
			fmt.Println("Successfully cleared all the meetings")
			Log.WriteLog(fmt.Sprintf("user %s successfully cleared all the meetings", userinfo.Username), 1)
		} else {
			fmt.Fprintf(os.Stderr, "Fail to clear all meetings: %v", err)
			Log.WriteLog(fmt.Sprintf("user %s fail to clear all meetings", userinfo.Username), 0)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(clearmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
