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

func (rm *RmLib) logVerbose(s string) {
	if rmFlg.verbose {
		fmt.Println(s)
	}
}

func (rm *RmLib) delAll(path string) {
	rm.handleError(os.RemoveAll(path))
	rm.logVerbose(path)
}

func (rm *RmLib) del(path string) {
	rm.handleError(os.Remove(path))
	rm.logVerbose(path)
}

func (rm *RmLib) delHierarchy(path string) {
	dirVoid, dirVoidErr := rm.isDirEmpty(path)
	rm.handleError(dirVoidErr)
	if dirVoid {
		rm.del(path)
	} else {
		if rmFlg.force {
			items, itemErr := readDir(path)
			rm.handleError(itemErr)
			for _, item := range items {
				if item.IsDir() {
					dirPath := filepath.Join(path, item.Name())
					if rmFlg.dir {
						rm.delAll(dirPath)
					} else if rmFlg.hierarchy {
						rm.delHierarchy(dirPath)
						// todo: rm.del(dirPath)
					}
				} else {
					rm.del(filepath.Join(path, item.Name()))
				}
			}
		} else {
			fmt.Println(path, "is a non-empty directory.")
		}
	}
}

type RmFlag struct {
	dir       bool
	force     bool
	confirm   bool
	overwrite bool
	hierarchy bool
	recursive bool
	verbose   bool
	whiteout  bool
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
				if rmFlg.recursive && rmFlg.force {
					rmLib.delAll(args[i])
				} else {
					rmLib.delHierarchy(args[i])
				}
			} else {
				rmLib.del(args[i])
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
	rmCmd.Flags().BoolVarP(&rmFlg.dir, "dir", "d", false, "Attempt to remove directories as well as other types of files.")
	rmCmd.Flags().BoolVarP(&rmFlg.force, "force", "f", false, "Attempt to remove the files without prompting for confirmation, regardless of the file's permissions.")
	// rmCmd.Flags().BoolVarP(&rmFlg.confirm, "confirm", "i", false, "Request confirmation before attempting to remove each file, regardless of the file's permissions.")
	// rmCmd.Flags().BoolVarP(&rmFlg.overwrite, "overwrite", "P", false, "Overwrite regular files before deleting them.")
	rmCmd.Flags().BoolVarP(&rmFlg.hierarchy, "hierarchy", "R", false, "Attempt to remove the file hierarchy rooted in each file argument.")
	rmCmd.Flags().BoolVarP(&rmFlg.recursive, "recursive", "r", false, "Equivalent to -R.")
	rmCmd.Flags().BoolVarP(&rmFlg.verbose, "verbose", "v", false, "Be verbose when deleting files, showing them as they are removed.")
	// rmCmd.Flags().BoolVarP(&rmFlg.whiteout, "whiteout", "W", false, "Attempt to undelete the named files and recover files covered by whiteouts.")
}
