/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"../lib"
	"github.com/spf13/cobra"
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
	i18nRmCmdConfirmationMsg = "Delete %s \nAre you sure? (yes/no) "
)

type RmLib struct {
}

func (rm *RmLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (rm *RmLib) isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func (rm *RmLib) readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (rm *RmLib) logVerbose(s string) {
	if rmFlg.verbose {
		fmt.Println(s)
	}
}

func (rm *RmLib) isWritable(s string) {
	writable := false
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		writable = true
	}
	return writable
}

func (rm *RmLib) delAll(path string) {
	rm.handleError(os.RemoveAll(path))
	rm.logVerbose(path)
}

func (rm *RmLib) del(path string) {
	if rmFlg.interactive {
		fmt.Printf(i18nRmCmdConfirmationMsg, path)
		if askForConfirmation() {
			rm.handleError(os.Remove(path))
			rm.logVerbose(path)
		}
	} else {
		rm.handleError(os.Remove(path))
		rm.logVerbose(path)
	}
}

func (rm *RmLib) delFile(path string) {
	if rm.isWritable(path) || rmFlg.force {
		rm.del(path)
	} else {
		rm.logVerbose("Permission denied to delete", path)
	}
}

func (rm *RmLib) delFolder(path string) {
	vacant, err := rm.isDirEmpty(path)
	rm.handleError(err)
	if vacant {
		rm.del(path)
	} else {
		if rmFlg.Recursive || rmFlg.recursive {
			if rmFlg.force {
				rm.delAll(path)
			} else {
				rm.delContent(path)
			}
		} else {
			fmt.Println(path, "is a non-empty directory.")
		}
	}
}

func (rm *RmLib) delContent(path string) {
	items, err := rm.readDir(path)
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
			path, pathErr := os.Stat(args[i])
			rmLib.handleError(pathErr)
			if path.IsDir() {
				rmLib.delFolder(args[i])
			} else {
				rmLib.delFile(args[i])
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
