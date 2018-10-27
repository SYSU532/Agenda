package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createUserName, createUserPass, createUserEmail string

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&createUserName, "username", "u", "", "The username of the new user.")
	registerCmd.Flags().StringVarP(&createUserPass, "password", "p", "", "The password of the new user.")
	registerCmd.Flags().StringVarP(&createUserEmail, "email", "e", "", "The email of the new user.")
	registerCmd.MarkFlagRequired("username")
	registerCmd.MarkFlagRequired("password")
	registerCmd.MarkFlagRequired("email")
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Long:  fmt.Sprintf(`Register a new user with the input username, password and email.

Usage: %v register -uUserName –password pass –email=a@xxx.com`, os.Args[0]),

	Run: func(cmd *cobra.Command, args []string) {
		

	},


}

