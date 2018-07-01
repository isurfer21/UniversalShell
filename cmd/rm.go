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
	i18nRmCmdTitle  = "Remove files and directories"
	i18nRmCmdDetail = `
Remove files and directories

It attempts to remove the non-directory type files specified on the 
command line. If the permissions of the file do not permit writing,
and the standard input device is a terminal, the user is prompted
(on the standard error output) for confirmation.
`
)

type rmFlag struct {
}

var rmFlg rmFlag

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: i18nRmCmdTitle,
	Long:  i18nRmCmdDetail,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			path, pathErr := os.Stat(args[0])
			if pathErr == nil {
				if path.IsDir() {
					dirErr := os.RemoveAll(args[0])
					if dirErr != nil {
						fmt.Println(dirErr)
						os.Exit(1)
					}
				} else {
					fileErr := os.Remove(args[0])
					if fileErr != nil {
						fmt.Println(fileErr)
						os.Exit(1)
					}
				}
			} else {
				fmt.Println(pathErr)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command	is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
