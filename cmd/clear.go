/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
)

const (
	i18nClearCmdTitle  = "Clear terminal screen"
	i18nClearCmdDetail = `
Clear terminal screen

It clears your screen if this is possible. It looks in the environment for the
terminal type and then in the terminfo database to figure out how to clear the
screen.

It ignores any command-line parameters that may be present.
`
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: i18nClearCmdTitle,
	Long:  i18nClearCmdDetail,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
