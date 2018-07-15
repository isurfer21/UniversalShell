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
	i18nDirnameCmdTitle  = "Convert a full pathname to just a path"
	i18nDirnameCmdDetail = `
Convert a full pathname to just a path

Prints all but the final slash-delimited component of a string (presumably a 
filename).

If PATHNAME is a single component, dirname prints . (meaning the current 
directory)
`
)

type DirnameLib struct {
}

type DirnameFlag struct {
}

var (
	dirnameFlg DirnameFlag
	dirnameLib DirnameLib
)

// dirnameCmd represents the dirname command
var dirnameCmd = &cobra.Command{
	Use:   "dirname",
	Short: i18nDirnameCmdTitle,
	Long:  i18nDirnameCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(filepath.Dir(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(dirnameCmd)
}
