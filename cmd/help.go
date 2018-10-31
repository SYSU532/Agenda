//Deprecated Since Cobra automatically add help subcommand


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



//import (
//	"fmt"
//	"github.com/spf13/cobra"
//	"os"
//)
//
//// helpCmd represents the help command
//var helpCmd = &cobra.Command{
//	Use:   "help",
//	Short: "Display help messages of Agenda.",
//	Long: fmt.Sprintf(`An agenda management cli app.
//
//Usage: %v {help|register|login|logout|userlist|userdelete|cm|fm|cancelm|quitm|clearm|addp|rp}`, os.Args[0]),
//	Run: func(cmd *cobra.Command, args []string) {
//		if len(os.Args) <= 2 {
//			// Print help messages for using agenda
//			fmt.Println("Agenda -- A local meeting management cli app")
//			fmt.Println("\nUsage: agenda [command]")
//			// Print availbale  commands
//			fmt.Printf("\nAvailable Commands:\n  %-50v  Register a new user\n", "register -uUsername -pPassword -eEmail -oPhone")
//			fmt.Printf("  %-50v  Add participants into the meeting you create\n", "addp -tTitle -p Participant1 Participant2 .....")
//			fmt.Printf("  %-50v  Remove participants from the meeting you create\n", "rp -tTitle -p Participant1 Participant2 .....")
//			fmt.Printf("  %-50v  Create meeting with special conditions\n", "cm -tTitle -p Participant1 Participant2 .....")
//			fmt.Printf("  %-50v  List all the conditional users in the database\n", "userlist [-uUsername] [-eEmail]")
//			fmt.Printf("  %-50v  Find and list all meetings satisfy the conditions\n", "fm [-tTitle] [-sStartTime -eEndTime]")
//			fmt.Printf("  %-50v  Login user\n", "login -uUsername -pPassword")
//			fmt.Printf("  %-50v  Quit the meeting that you take part in\n", "quitm -tTile")
//			fmt.Printf("  %-50v  Cancel the meeting that you create\n", "cancelm -tTitle")
//			fmt.Printf("  %-50v  Clear all the meetings that you create\n", "clearm")
//			fmt.Printf("  %-50v  Logout current user\n", "logout")
//			fmt.Printf("  %-50v  Delete current user\n", "userdelete")
//			// Print tips
//			fmt.Println("\nTips:\n  Use \"agenda [command] --help\" for more information about a command")
//		}else {
//
//		}
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(helpCmd)
//}
