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
	i18nMvCmdTitle  = "Move or rename files and directories"
	i18nMvCmdDetail = `
Move or rename files and directories

It moves each file named by a source operand to a destination file in the 
existing directory named by the directory operand.

It renames (moves) oldpath to newpath. If newpath already exists and is not a 
directory, Rename replaces it. OS-specific restrictions may apply when oldpath 
and newpath are in different directories.
`
)

type MvFlag struct {
	path bool
}

var mvFlg MvFlag

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: i18nMvCmdTitle,
	Long:  i18nMvCmdDetail,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			err := os.Rename(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
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
