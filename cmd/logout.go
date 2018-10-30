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
	"github.com/SYSU532/agenda/entity"
	"github.com/SYSU532/agenda/Log"
	"github.com/spf13/cobra"
	"os"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from the current user",
	Long: fmt.Sprintf(`Log out from the current user.

Usage: %v logout`, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		Log.WriteLog("Invoke log out command to sign out current user", 1)
		info, err := entity.GetCurrentUser()
		if err != nil {
			fmt.Println("You are currently not login!")
			Log.WriteLog("You are currently not login!", 0)
			return
		}
		err = entity.ClearCurrentUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when logging out: %v\n", err)
			Log.WriteLog(fmt.Sprintf("Error when logging out: %v", err), 0)
			return
		}
		logMess := fmt.Sprintf("Successfully log out from %v", info.Username)
		fmt.Println(logMess)
		Log.WriteLog(logMess, 1)
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
