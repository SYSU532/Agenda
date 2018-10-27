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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help messages of Agenda.",
	Long: fmt.Sprintf(`Agenda is an agenda management cli app.

Usage: %v {help|register|cm}`, os.Args[0]),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("help called")
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

}
