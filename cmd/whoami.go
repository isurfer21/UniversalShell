/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

const (
	i18nWhoamiCmdTitle  = "Print the current user id and name"
	i18nWhoamiCmdDetail = `
Print the current user id and name

It displays your effective user ID as a name.

It has been obsoleted by the id utility, and is equivalent to 'ush id -u -n'.
`
)

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: i18nWhoamiCmdTitle,
	Long:  i18nWhoamiCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		activeUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(activeUser.Username)
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
