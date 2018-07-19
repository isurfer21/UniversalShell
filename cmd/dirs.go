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
	i18nDirsCmdTitle  = "Display list of remembered directories"
	i18nDirsCmdDetail = `
Display list of remembered directories

By default files are listed in columns, sorted vertically, and special
characters are represented by backslash escape sequences.
`
)

type DirsLib struct {
	dirStack lib.DirStack
}

func (dirs *DirsLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (dirs *DirsLib) flush() string {
	stack := dirs.dirStack.Stack
	if !dirsFlg.long {
		stack = dirs.dirStack.Short()
	}
	output := strings.Join(dirs.dirStack.Reverse(stack), " ")
	return output
}

func (dirs *DirsLib) list() string {
	stack := dirs.dirStack.Stack
	if !dirsFlg.long {
		stack = dirs.dirStack.Short()
	}
	output := strings.Join(dirs.dirStack.Reverse(stack), "\n")
	return output
}

func (dirs *DirsLib) vertical() string {
	stack := dirs.dirStack.Stack
	if !dirsFlg.long {
		stack = dirs.dirStack.Short()
	}
	stack = dirs.dirStack.Reverse(stack)
	for i := 0; i < len(stack); i += 1 {
		stack[i] = fmt.Sprintf("%d %s", i, stack[i])
	}
	output := strings.Join(stack, "\n")
	return output
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
		dirsLib.dirStack.Load()
		if dirsFlg.clear {
			err := dirsLib.dirStack.Clear()
			dirsLib.handleError(err)
			dirsLib.dirStack.Save()
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
	dirsCmd.Flags().BoolVarP(&dirsFlg.long, "long", "l", false, "produces a longer listing, avoids '~' for home directory")
	dirsCmd.Flags().BoolVarP(&dirsFlg.perline, "perline", "p", false, "print the directory stack with one entry per line")
	dirsCmd.Flags().BoolVarP(&dirsFlg.vertical, "vertical", "v", false, "print the directory stack with one entry per line with its index")
}
