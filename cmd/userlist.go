/*
*  CMD -- User List
 */

package cmd

import (
	"fmt"
	"github.com/SYSU532/agenda/entity"
	"github.com/SYSU532/agenda/log"
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
		// Write init lOG
		log.WriteLog("Invoke user list command to list all users in agenda system", 1)
		result, err := entity.GetUserList(targetUserName, targetEmail)
		if err != nil {
			fmt.Println("FAIL to print user list!")
			log.WriteLog("Fail to print user list", 0)
		} else if len(result) == 0 {
			if targetUserName == "" && targetEmail == "" {
				fmt.Println("There is no user in the DataBase!")
				log.WriteLog("No any user in the database", 0)
			} else {
				fmt.Println("There is no user satisfying your searching conditions!")
				log.WriteLog("No any user satisfies your searching conditions", 0)
			}
		} else {
			fmt.Printf("|%-20v|%-20v|%-20v|\n", "Username", "Email", "Phone")
			fmt.Println("|--------------------|--------------------|--------------------|")
			for _, ele := range result {
				fmt.Printf("|%-20v|%-20v|%-20v|\n", ele.Username, ele.Email, ele.Phone)
			}
			log.WriteLog("List users done!Print target users info....", 1)
		}
	},
}
