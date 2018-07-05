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
	i18nTmpdirCmdTitle  = "Path of default directory to use for temporary files"
	i18nTmpdirCmdDetail = `
Returns the path of default directory to use for temporary files

On Unix systems, it returns $TMPDIR if non-empty, else /tmp. On Windows, it 
uses GetTempPath, returning the first non-empty value from %TMP%, %TEMP%, 
%USERPROFILE%, or the Windows directory. On Plan 9, it returns /tmp.

The directory is neither guaranteed to exist nor have accessible permissions.
`
)

// tmpdirCmd represents the tmpdir command
var tmpdirCmd = &cobra.Command{
	Use:   "tmpdir",
	Short: i18nTmpdirCmdTitle,
	Long:  i18nTmpdirCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(os.TempDir())
	},
}

func init() {
	rootCmd.AddCommand(tmpdirCmd)
}
