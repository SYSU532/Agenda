/*
* CMD -- Login
 */

package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/SYSU532/agenda/entity"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var loginUsername, loginPassword string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login User.",
	Long:  `Login user in order to perform operations like creating meetings, view current meetings, etc.`,
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
		//LoginAgendaTerminal(loginUsername)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&loginUsername, "username", "u", "", "The username of the user.")
	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "The password of the user.")
}

func LoginAgendaTerminal(name string) {
	var (
		input  string
		output bytes.Buffer
		items  []string
	)
	s := fmt.Sprintf("%s@Agenda~: ", name)
	for {
		fmt.Printf(s)
		fmt.Scanln(&input)
		items = strings.Split(input, " ")
		if items[0] == "exit" {
			//entity.ClearCurrentUser()
			break
		} else if items[0] == "login" {
			fmt.Println("Do not login again in Agenda Termial!")
		} else {
			cmd := exec.Command(os.Args[0], input)
			cmd.Stdout = &output
			cmd.Run()
			fmt.Printf(output.String())
			output.Reset()
		}
	}
}


