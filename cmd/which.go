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
	i18nWhichCmdTitle  = "Locate program file path"
	i18nWhichCmdDetail = `
Locate program file path

Locate a program file in the user's path. If no program filename is provided, 
it will return current program file's location.
`
)

// whichCmd represents the which command
var whichCmd = &cobra.Command{
	Use:   "which",
	Short: i18nWhichCmdTitle,
	Long:  i18nWhichCmdDetail,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			path, err := os.Executable()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(path)
		} else {

		}
	},
}

func init() {
	rootCmd.AddCommand(whichCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
