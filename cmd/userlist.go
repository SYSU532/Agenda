/*
*  CMD -- User List
 */

package cmd

import (
	"fmt"
	"github.com/SYSU532/agenda/entity"
	"os"

	"github.com/spf13/cobra"
)

var targetUserName, targetEmail string

func init() {
	rootCmd.AddCommand(userlistCmd)

	userlistCmd.Flags().StringVarP(&targetUserName, "username", "u", "", "The username for searching.")
	userlistCmd.Flags().StringVarP(&targetEmail, "email", "e", "", "The email for searching.")
}

var userlistCmd = &cobra.Command{
	Use:   "userlist",
	Short: "Show all user list, or print user with special username or email",
	Long: fmt.Sprintf(`Print user list depending on the input conditions.

Usage: %v userlist [-uUserName] [-eEmail]`, os.Args[0]),

	Run: func(cmd *cobra.Command, args []string) {
		result, err := entity.GetUserList(targetUserName, targetEmail)
		if err != nil {
			fmt.Println("FAIL to print user list!")
		} else if len(result) == 0 {
			if targetUserName == "" && targetEmail == "" {
				fmt.Println("No any user in the DataBase!")
			} else {
				fmt.Println("No any user satisfies your searcing conditions!")
			}
		} else {
			for _, ele := range result {
				fmt.Printf("UserName: %v  Email: %v \n", ele.Username, ele.Email)
			}
		}
	},
}
