/*
* CMD -- Login
 */

package cmd

import (
	"bufio"
	"syscall"
	"fmt"
	"github.com/SYSU532/agenda/entity"
	"github.com/SYSU532/agenda/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

var loginUsername, loginPassword string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login User.",
	Long:  fmt.Sprintf(`Login user in order to perform operations like creating meetings, view current meetings, etc.

Usage: %v login [-u username] [-p password]`, os.Args[0]),

	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		log.WriteLog("Invoke log in command to log user with username and password", 1)
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
			log.WriteLog(fmt.Sprintf("Fail to login: %v\n", err), 0)
			os.Exit(0)
		}
		err = entity.SetCurrentUser(loginUsername, loginPassword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to login: %v\n", err)
			log.WriteLog(fmt.Sprintf("Fail to login: %v\n", err), 0)
			os.Exit(0)
		}
		logMess := fmt.Sprintf("Login as user %v succeeded", loginUsername)
		fmt.Println(logMess)
		log.WriteLog(logMess, 1)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&loginUsername, "username", "u", "", "The username of the user.")
	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "The password of the user.")
}

