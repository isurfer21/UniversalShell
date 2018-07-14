/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

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

Syntax:
  dirname pathname

Example:
  Extract the path from the variable 'pathnamevar' and store in the variable 
  result using parameter expansion
`
)

type DirnameLib struct {
}

func (dirname *DirnameLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (dirname *DirnameLib) getParent(pathname string) string {
	p := strings.Split(pathname, "/")
	r := p[len(p)-2]
	if r == "" {
		r = "/"
	}
	return r
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
		if strings.Count(args[0], "/") > 0 {
			fmt.Println(dirnameLib.getParent(args[0]))
		} else {
			fmt.Println(".")
		}
	},
}

func init() {
	rootCmd.AddCommand(dirnameCmd)
}
