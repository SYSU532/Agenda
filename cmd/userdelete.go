/*
*  CMD -- User Delete
*/

package cmd

import (
	"fmt"
	"os"
	"github.com/SYSU532/agenda/entity"

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
		// Delete the current user, do not need any inpit
		name, err := entity.DeleteUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when deleting user: %v\n", err)
		} else {
			fmt.Printf("Successfully delete current user: %v!\n", name)
		}
	},
}
