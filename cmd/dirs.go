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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	i18nDirsCmdTitle  = "Display list of remembered directories"
	i18nDirsCmdDetail = `
Display list of remembered directories

By default files are listed in columns, sorted vertically, and special
characters are represented by backslash escape sequences.
`
)

type DirsLib struct {
	stack []string
}

func (dirs *DirsLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (dirs *DirsLib) load() {
	str := viper.GetString("DIRSTACK")
	if len(str) > 0 {
		dirs.stack = strings.Split(str, ",")
	} else {
		dirs.stack = []string{}
	}
}

func (dirs *DirsLib) clear() error {
	if len(dirs.stack) > 0 {
		dirs.stack = []string{}
		return nil
	}
	return errors.New("Directory stack is empty!")
}

func (dirs *DirsLib) save() {
	str := strings.Join(dirs.stack, ",")
	viper.Set("DIRSTACK", str)
	viper.WriteConfig()
}

func (dirs *DirsLib) flush() string {
	home, err := homedir.Dir()
	dirs.handleError(err)
	str := strings.Join(dirs.stack, " ")
	str = strings.Replace(str, home, "~", -1)
	return str
}

func (dirs *DirsLib) long() string {
	str := strings.Join(dirs.stack, " ")
	return str
}

func (dirs *DirsLib) list() string {
	str := strings.Join(dirs.stack, "\n")
	return str
}

func (dirs *DirsLib) vertical() string {
	stack := []string{}
	for i := 0; i < len(dirs.stack); i += 1 {
		stack = append(stack, fmt.Sprintf("%d %s", i, dirs.stack[i]))
	}
	str := strings.Join(stack, "\n")
	return str
}

type DirsFlag struct {
	clear    bool
	long     bool
	perline  bool
	vertical bool
}

var (
	dirsFlg DirsFlag
	dirsLib DirsLib
)

// dirsCmd represents the dirs command
var dirsCmd = &cobra.Command{
	Use:   "dirs",
	Short: i18nDirsCmdTitle,
	Long:  i18nDirsCmdDetail,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		dirsLib.load()
		if dirsFlg.clear {
			err := dirsLib.clear()
			dirsLib.handleError(err)
			dirsLib.save()
		} else if dirsFlg.long {
			fmt.Println(dirsLib.long())
		} else if dirsFlg.perline {
			fmt.Println(dirsLib.list())
		} else if dirsFlg.vertical {
			fmt.Println(dirsLib.vertical())
		} else {
			fmt.Println(dirsLib.flush())
		}
	},
}

func init() {
	rootCmd.AddCommand(dirsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// dirsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	dirsCmd.Flags().BoolVarP(&dirsFlg.clear, "clear", "c", false, "clears the directory stack by deleting all of the elements")
	dirsCmd.Flags().BoolVarP(&dirsFlg.long, "long", "l", false, "produces a longer listing")
	dirsCmd.Flags().BoolVarP(&dirsFlg.perline, "perline", "p", false, "print the directory stack with one entry per line")
	dirsCmd.Flags().BoolVarP(&dirsFlg.vertical, "vertical", "v", false, "print the directory stack with one entry per line with its index")
}
