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
	"github.com/SYSU532/agenda/entity"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var(
	cmTitle string
	cmParticipators []string
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "Create Meeting",
	Long: fmt.Sprintf(`Use this command to create a meeting using a already logged in user.

Usage: %v cm [-t title -p participator1, participator2, ...]`, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		userInfo, err := entity.GetCurrentUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to create meeting: %v\n", err)
			return
		}
		reader := bufio.NewReader(os.Stdin)
		if cmTitle == "" {
			fmt.Print("Enter the meeting title: ")
			title, _ := reader.ReadString('\n')
			cmTitle = title[:len(title)-1]
		}
		if len(cmParticipators) == 1 && cmParticipators[0] == "" {
			cmParticipators = []string{}
			fmt.Print("Enter the number of participators: ")
			var partNum uint
			fmt.Scan(&partNum)
			for i := uint(0); i < partNum; i++ {
				var part string
				fmt.Printf("Enter participator %d: ", i)
				fmt.Scan(&part)
				cmParticipators = append(cmParticipators, part)
			}
		}

		format := "2006-01-02 15:04"
		fmt.Print("Enter the start time of the meeting (format: YYYY-mm-dd hh:mm): ")
		start, _ := reader.ReadString('\n')
		start = start[:len(start)-1]
		startTime, err := time.Parse(format, start)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to parse start time\n")
			return
		}
		fmt.Print("Enter the end time of the meeting: ")
		end, _ := reader.ReadString('\n')
		end = end[:len(end)-1]
		endTime, err := time.Parse(format, end)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to parse end time\n")
			return
		}

		err = entity.AddMeeting(cmTitle, userInfo.Username, startTime, endTime)
		if err == nil {
			fmt.Println("Successfully created meeting,")
		} else {
			fmt.Fprintf(os.Stderr, "Fail to create meeting: %v\n", err)
			return
		}

		for _, part := range cmParticipators {
			err = entity.AddPaticipant(cmTitle, part)
			if err != nil {
				fmt.Printf("Fail to add participant %v: %v\n", part, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)

	cmCmd.Flags().StringVarP(&cmTitle, "title", "t", "", "The title of the meeting to be created")
	cmCmd.Flags().StringArrayVarP(&cmParticipators, "participators", "p", []string{""}, "All the participators of the meeting to be created")
}
