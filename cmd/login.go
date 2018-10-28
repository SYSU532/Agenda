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
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

var loginUsername, loginPassword string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login User.",
	Long: `Login user in order to perform operations like creating meetings, view current meetings, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		if loginUsername == "" {
			fmt.Print("Enter username: ")
			loginUsername, _ = reader.ReadString('\n')
			//trim \n
			loginUsername = loginUsername[:len(loginUsername)-1]
		}
		if loginPassword == "" {
			fmt.Print("Enter password: ")
			bytePass, _ := terminal.ReadPassword(int(syscall.Stdin))
			loginPassword = string(bytePass)
		}
		fmt.Println()
		err := entity.LoginUser(loginUsername, loginPassword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to login: %v\n", err)
			os.Exit(0)
		}
		err = entity.SetCurrentUser(loginUsername, loginPassword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to login: %v\n", err)
			os.Exit(0)
		}
		fmt.Printf("Login as user %v succeeded.\n", loginUsername)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&loginUsername, "username", "u", "", "The username of the user.")
	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "The password of the user.")

}
