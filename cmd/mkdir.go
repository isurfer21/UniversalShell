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
	i18nMkdirCmdTitle  = "Make directories"
	i18nMkdirCmdDetail = `
Make directories

It creates the directories named as operands, in the order specified, using 
mode rwxrwxrwx (0777). 
`
)

type MkdirFlag struct {
	path bool
}

var mkdirFlg MkdirFlag

// mkdirCmd represents the mkdir command
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: i18nMkdirCmdTitle,
	Long:  i18nMkdirCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			if mkdirFlg.path {
				pathErr := os.MkdirAll(args[i], 0777)
				if pathErr != nil {
					fmt.Println(pathErr)
					os.Exit(1)
				}
			} else {
				fileErr := os.Mkdir(args[i], 0777)
				if fileErr != nil {
					fmt.Println(fileErr)
					os.Exit(1)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mkdirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// mkdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	mkdirCmd.Flags().BoolVarP(&mkdirFlg.path, "path", "p", false, "create intermediate directories as required")
}
