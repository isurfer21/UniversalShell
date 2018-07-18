/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	i18nPushdCmdTitle  = "Save and then change the current directory"
	i18nPushdCmdDetail = `
Save and then change the current directory

With no arguments, pushd exchanges the top two directories.
`
)

type PushdLib struct {
	stack []string
}

func (pushd *PushdLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (pushd *PushdLib) load() {
	str := viper.GetString("DIRSTACK")
	if len(str) > 0 {
		pushd.stack = strings.Split(str, ",")
	} else {
		pushd.stack = []string{}
	}
}

func (pushd *PushdLib) push(path string) {
	pushd.stack = append(pushd.stack, path)
}

func (pushd *PushdLib) save() {
	str := strings.Join(pushd.stack, ",")
	viper.Set("DIRSTACK", str)
	viper.WriteConfig()
}

func (pushd *PushdLib) flush() string {
	str := strings.Join(pushd.stack, " ")
	return str
}

type PushdFlag struct {
	nochange bool
}

var (
	pushdFlg PushdFlag
	pushdLib PushdLib
)

// pushdCmd represents the pushd command
var pushdCmd = &cobra.Command{
	Use:   "pushd",
	Short: i18nPushdCmdTitle,
	Long:  i18nPushdCmdDetail,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pushdLib.load()
		absPath, err := filepath.Abs(args[0])
		pushdLib.handleError(err)
		pushdLib.push(absPath)
		pushdLib.save()
		os.Chdir(absPath)
		fmt.Println(pushdLib.flush())
	},
}

func init() {
	rootCmd.AddCommand(pushdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// pushdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	pushdCmd.Flags().BoolVarP(&pushdFlg.nochange, "no-change", "n", false, "suppresses the normal change of directory")
}
