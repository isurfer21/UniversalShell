/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	i18nPwdCmdTitle  = "Path of working directory"
	i18nPwdCmdDetail = `
Path of working directory

Writes the absolute pathname of the current working directory to the standard 
output.
`
)

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: i18nPwdCmdTitle,
	Long:  i18nPwdCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(dir)
	},
}

func init() {
	rootCmd.AddCommand(pwdCmd)
}
