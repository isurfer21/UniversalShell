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
	i18nRenCmdTitle  = "Rename files and directories"
	i18nRenCmdDetail = `
Rename files and directories

It renames (moves) oldpath to newpath. If newpath already 
exists and is not a directory, Rename replaces it. 
OS-specific restrictions may apply when oldpath and 
newpath are in different directories.
`
)

type renFlag struct {
}

var renFlg renFlag

// renCmd represents the ren command
var renCmd = &cobra.Command{
	Use:   "ren",
	Short: i18nRenCmdTitle,
	Long:  i18nRenCmdDetail,
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
	rootCmd.AddCommand(renCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// renCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// renCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
