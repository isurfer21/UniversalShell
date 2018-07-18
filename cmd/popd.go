/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	stack []string
}

func (popd *PopdLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (popd *PopdLib) load() {
	str := viper.GetString("DIRSTACK")
	if len(str) > 0 {
		popd.stack = strings.Split(str, ",")
	} else {
		popd.stack = []string{}
	}
}

func (popd *PopdLib) pop() (string, error) {
	if len(popd.stack) > 0 {
		lastItem := popd.stack[len(popd.stack)-1]
		popd.stack = popd.stack[:len(popd.stack)-1]
		return lastItem, nil
	}
	return "", errors.New("Directory stack is empty!")
}

func (popd *PopdLib) save() {
	str := strings.Join(popd.stack, ",")
	viper.Set("DIRSTACK", str)
	viper.WriteConfig()
}

func (popd *PopdLib) flush() string {
	str := strings.Join(popd.stack, " ")
	return str
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
		popdLib.load()
		path, err := popdLib.pop()
		popdLib.handleError(err)
		popdLib.save()
		os.Chdir(path)
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
