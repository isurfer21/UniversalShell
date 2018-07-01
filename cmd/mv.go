/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	i18nMvCmdTitle  = "Move files and directories"
	i18nMvCmdDetail = `
Move files and directories

It mv moves each file named by a source operand to a
destination file in the existing directory named by the 
directory operand.
`
)

type mvFlag struct {
	path bool
}

var mvFlg mvFlag

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: i18nMvCmdTitle,
	Long:  i18nMvCmdDetail,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			fmt.Println("mv called")
		}
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// mvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
