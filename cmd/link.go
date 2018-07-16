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
	i18nLinkCmdTitle  = "Create a link to a file"
	i18nLinkCmdDetail = `
Create a link to a file

Create a link named FILE2 to an existing FILE1.

Syntax:
  link FILE1 FILE2
  link OPTION
`
)

type LinkLib struct {
}

func (link *LinkLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type LinkFlag struct {
}

var (
	linkFlg LinkFlag
	linkLib LinkLib
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: i18nLinkCmdTitle,
	Long:  i18nLinkCmdDetail,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		linkLib.handleError(os.Link(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
