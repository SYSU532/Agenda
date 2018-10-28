/*
*  CMD -- User Delete
*/

package cmd

import (
	"fmt"
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
		// Delete the user by current user info
		err := entity.DeleteCurrUser()
		if err != nil {
			fmt.Println("FAIL to delete current user!")
		}else {
			fmt.Println("Successfully delete current user!")
		}
	},
}
