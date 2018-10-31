/*
*  CMD -- User Delete
 */

package cmd

import (
	"fmt"
	"github.com/SYSU532/agenda/entity"
	"github.com/SYSU532/agenda/log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userdeleteCmd)
	// Do not need any pflags
}

var userdeleteCmd = &cobra.Command{
	Use:   "userdelete",
	Short: "Delete the user itself",
	Long: fmt.Sprintf(`Delete the user if the user choose to do so.

Usage: %v userdelete`, os.Args[0]),

	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		log.WriteLog("Invoke delete user command to delete user itself", 1)
		// Delete the current user, do not need any inpit
		name, err := entity.DeleteUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when deleting user: %v\n", err)
			log.WriteLog(fmt.Sprintf("Error with deleting current user %s", name), 0)
		} else {
			fmt.Printf("Successfully delete current user: %v!\n", name)
			log.WriteLog(fmt.Sprintf("Successfully delete current user %s", name), 1)
		}
		entity.ClearCurrentUser()
	},
}
