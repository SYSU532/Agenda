package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

var createUserName, createUserPass, createUserEmail string

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&createUserName, "username", "u", "user", "The username of the new user.")
	registerCmd.Flags().StringVarP(&createUserPass, "password", "p", "pass", "The password of the new user.")
	registerCmd.Flags().StringVarP(&createUserEmail, "email", "e", "email", "The email of the new user.")

}

func formatCheck() (result bool, errMess string) {
	var (
		userCheck,emailCheck bool
		checkMess string
	)
	userFormat, _ := regexp.Compile("^[A-Za-z0-9]+$");
	emailFormat, _ := regexp.Compile("^([A-Za-z0-9]+)@([a-z0-9]+).([a-z]+)$");
	// Checking UserName start format
	userCheck = userFormat.MatchString(createUserName);
	if userCheck != true {
		checkMess = "UserName Format Error! Must be sequence of numbers and letters!\n";
	}
	// Checking Email standard format
	emailCheck = emailFormat.MatchString(createUserEmail);
	if emailCheck != true {
		checkMess += "Email Format Error! Must be like XXX@XXX.XXX!\n";
	}

	return (userCheck && emailCheck), checkMess;
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Long:  fmt.Sprintf(`Register a new user with the input username, password and email.

Usage: %v register {help | -uUserName –password pass –email=a@xxx.com}`, os.Args[0]),

	Run: func(cmd *cobra.Command, args []string) {
		// Format checking
		checkRes, errMess := formatCheck();
		if checkRes != true {
			fmt.Println(errMess, "User create fail....");
			os.Exit(2);
		}
		/*
		// Duplication checking
		dcheckRes, err1 := entity.dupCheck(createUserName, createUserPass, createUserEmail);
		if dcheckRes != true {

		}*/
		// Double Check Success, create the user
		fmt.Printf("Successfully regist a new user: %s, email: %s\n", createUserName, createUserEmail);
		
	},


}

