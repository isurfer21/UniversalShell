/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	i18nBasenameCmdTitle  = "Strip directory and suffix from filenames"
	i18nBasenameCmdDetail = `
Strip directory and suffix from filenames

It will print NAME with any leading directory components removed. If specified,
it will also remove a trailing SUFFIX (typically a file extention).
`
)

type BasenameLib struct {
}

type BasenameFlag struct {
}

var (
	basenameFlg BasenameFlag
	basenameLib BasenameLib
)

// basenameCmd represents the basename command
var basenameCmd = &cobra.Command{
	Use:   "basename",
	Short: i18nBasenameCmdTitle,
	Long:  i18nBasenameCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(filepath.Base(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(basenameCmd)
}
