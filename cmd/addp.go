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

var (
	addpTitle         string
	addpParticipators []string
)

// addpCmd represents the addp command
var addpCmd = &cobra.Command{
	Use:   "addp",
	Short: "Add participator",
	Long: fmt.Sprintf(`Use this command to add participators to a meeting
	using a already logged in user.
	Usage: %v addp [-t title -p participator1, participator2, ...]`, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		Log.WriteLog("Invoke add participant command to add persons in meeting", 1)
		userinfo, err := entity.GetCurrentUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to add participator: %v\n", err)
			Log.WriteLog("Error when geeting current user, maybe you are not logged in", 0)
			return
		}
		reader := bufio.NewReader(os.Stdin)
		if addpTitle == "" {
			fmt.Print("Enter the meeting title: ")
			title, _ := reader.ReadString('\n')
			addpTitle = title[:len(title)-1]
		}
		if len(cmParticipators) == 1 && cmParticipators[0] == "" {
			addpParticipators = []string{}
			fmt.Print("Enter the number of participators: ")
			var partNum uint
			fmt.Scan(&partNum)
			for i := uint(0); i < partNum; i++ {
				var part string
				fmt.Printf("Enter participator %d: ", i)
				fmt.Scan(&part)
				addpParticipators = append(addpParticipators, part)
			}
		}
		err = entity.CheckBeforeModP(addpTitle, userinfo.Username)
		if err == nil {
			for _, part := range addpParticipators {
				err = entity.CheckDupPart(addpTitle, part)
				if err != nil {
					tmp := fmt.Sprintf("Meeting %v Fail to add participant %v: %v", addpTitle, part, err)
					fmt.Println(tmp)
					Log.WriteLog(tmp, 0)
					continue
				}
				err = entity.AddPaticipant(addpTitle, part)
				if err != nil {
					tmp := fmt.Sprintf("Meeting %v Fail to add participant %v: %v", addpTitle, part, err)
					fmt.Println(tmp)
					Log.WriteLog(tmp, 0)
				}
			}
			if err == nil {
				fmt.Println("Successfully added participant(s)")
				Log.WriteLog(fmt.Sprintf("Successfully add participant into Meeting %v", addpTitle), 1)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Fail to add participant: %v\n", err)
			Log.WriteLog(fmt.Sprintf("Fail to add participant: %v", err), 0)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addpCmd)

	addpCmd.Flags().StringVarP(&addpTitle, "title", "t", "", "The title of the meeting")
	addpCmd.Flags().StringArrayVarP(&addpParticipators, "participators", "p", []string{""}, "All the participators to be added")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
