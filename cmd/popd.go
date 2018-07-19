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

	"../lib"
)

const (
	i18nPopdCmdTitle  = "Restore the previous value of the current directory"
	i18nPopdCmdDetail = `
Restore the previous value of the current directory

Remove the top entry from the directory stack, and cd to the new top directory.

When no arguments are given, popd removes the top directory from the stack and
performs a cd to the new top directory. 

The elements are numbered from 0 starting at the first directory listed with
dirs; i.e., popd is equivalent to popd +0.
`
)

type PopdLib struct {
	dirStack lib.DirStack
}

func (popd *PopdLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (popd *PopdLib) flush() string {
	output := strings.Join(popd.dirStack.Reverse(popd.dirStack.Short()), " ")
	return output
}

type PopdFlag struct {
	nochange bool
}

var (
	popdFlg PopdFlag
	popdLib PopdLib
)

// popdCmd represents the popd command
var popdCmd = &cobra.Command{
	Use:   "popd",
	Short: i18nPopdCmdTitle,
	Long:  i18nPopdCmdDetail,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		popdLib.dirStack.Load()
		path, err := popdLib.dirStack.Pop()
		popdLib.handleError(err)
		popdLib.dirStack.Save()
		if !popdFlg.nochange {
			os.Chdir(path)
			fmt.Printf("<- cd %s\n", path)
		}
		fmt.Println(popdLib.flush())
	},
}

func init() {
	rootCmd.AddCommand(popdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// popdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	popdCmd.Flags().BoolVarP(&popdFlg.nochange, "no-change", "n", false, "suppresses the normal change of directory")
}
