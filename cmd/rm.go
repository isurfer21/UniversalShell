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

	"../lib"
)

const (
	i18nRmCmdTitle  = "Remove files and directories"
	i18nRmCmdDetail = `
Remove files and directories

It attempts to remove the non-directory type files specified on the command 
line. If the permissions of the file do not permit writing, and the standard 
input device is a terminal, the user is prompted (on the standard error output) 
for confirmation.
`
	i18nRmTplConfirmationMsg = "Delete %s \nAre you sure? (yes/no) "
	i18nRmTplPermDeniedToDel = "Permission denied to delete "
	i18nRmTplIsNonEmptyDir   = "%s is a non-empty directory.\n"
)

type RmLib struct {
	evoke   lib.Confirm
	dossier lib.Dossier
	folder  lib.Folder
}

func (rm *RmLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (rm *RmLib) catchError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (rm *RmLib) logVerbose(s string) {
	if rmFlg.verbose {
		fmt.Println(s)
	}
}

func (rm *RmLib) delAll(path string) {
	if rmFlg.interactive {
		fmt.Printf(i18nRmTplConfirmationMsg, path)
		if rm.evoke.AskForConfirmation() {
			rm.handleError(os.RemoveAll(path))
			rm.logVerbose(path)
		}
	} else {
		rm.handleError(os.RemoveAll(path))
		rm.logVerbose(path)
	}
}

func (rm *RmLib) del(path string) {
	if rmFlg.interactive {
		fmt.Printf(i18nRmTplConfirmationMsg, path)
		if rm.evoke.AskForConfirmation() {
			rm.handleError(os.Remove(path))
			rm.logVerbose(path)
		}
	} else {
		rm.handleError(os.Remove(path))
		rm.logVerbose(path)
	}
}

func (rm *RmLib) delFile(path string) {
	// writable, err := rm.dossier.IsWritable(path)
	fi, err := os.Lstat(path)
	if err == nil {
		fm := fi.Mode()
		if fm.IsRegular() {
			rm.del(path)
		} else if fm&os.ModeSymlink != 0 {
			rm.del(path)
		} else if fm&os.ModeNamedPipe != 0 {
			rm.del(path)
		} else {
			rm.logVerbose(i18nRmTplPermDeniedToDel + path)
		}
	}
}

func (rm *RmLib) delFolder(path string) {
	vacant, err := rm.folder.IsDirEmpty(path)
	rm.handleError(err)
	if vacant {
		rm.del(path)
	} else {
		if rmFlg.Recursive || rmFlg.recursive {
			rm.delContent(path)
		} else {
			fmt.Printf(i18nRmTplIsNonEmptyDir, path)
		}
	}
}

func (rm *RmLib) forceDelete(path string) {
	err := os.Remove(path)
	if err != nil {
		if strings.Contains(err.Error(), "directory not empty") {
			if rmFlg.Recursive || rmFlg.recursive {
				rm.handleError(os.RemoveAll(path))
				rm.logVerbose(path)
			} else {
				fmt.Printf(i18nRmTplIsNonEmptyDir, path)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		rm.logVerbose(path)
	}
}

func (rm *RmLib) delContent(path string) {
	items, err := rm.folder.ReadDir(path)
	rm.handleError(err)
	for _, item := range items {
		if item.IsDir() {
			rm.delFolder(filepath.Join(path, item.Name()))
		} else {
			rm.delFile(filepath.Join(path, item.Name()))
		}
	}
}

type RmFlag struct {
	force       bool
	interactive bool
	Recursive   bool
	recursive   bool
	verbose     bool
}

var (
	rmLib RmLib
	rmFlg RmFlag
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: i18nRmCmdTitle,
	Long:  i18nRmCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			path := args[i]
			if rmFlg.force {
				rmLib.forceDelete(path)
			} else {
				fi, err := os.Lstat(args[i])
				rmLib.handleError(err)
				if fi.IsDir() {
					rmLib.delFolder(path)
				} else {
					rmLib.delFile(path)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command	is called directly, e.g.:
	rmCmd.Flags().BoolVarP(&rmFlg.force, "force", "f", false, "ignore nonexistent files and arguments, never prompt")
	rmCmd.Flags().BoolVarP(&rmFlg.interactive, "interactive", "i", false, "prompt before every removal")
	rmCmd.Flags().BoolVarP(&rmFlg.Recursive, "Recursive", "R", false, "remove directories and their contents recursively")
	rmCmd.Flags().BoolVarP(&rmFlg.recursive, "recursive", "r", false, "equivalent to -R.")
	rmCmd.Flags().BoolVarP(&rmFlg.verbose, "verbose", "v", false, "explain what is being done")
}
