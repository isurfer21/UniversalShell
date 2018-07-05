/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

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

type rmFlag struct {
	dir       bool
	file      bool
	confirm   bool
	overwrite bool
	hierarchy bool
	subdir    bool
	verbose   bool
	whiteout  bool
}

var rmFlg rmFlag

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: i18nRmCmdTitle,
	Long:  i18nRmCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			path, pathErr := os.Stat(args[i])
			if pathErr == nil {
				if path.IsDir() {
					/*if rmFlg.directories || rmFlg.files {
						items, err := readDir(path)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						} else {
							for _, file := range items {
								if file.IsDir() {
									isDir = "Dir"
								} else {
									if rmFlg.files {
										fileErr := os.Remove(file)
										if fileErr != nil {
											fmt.Println(fileErr)
											os.Exit(1)
										}
									}
								}
							}
						}
					}*/
					dirErr := os.RemoveAll(args[i])
					if dirErr != nil {
						fmt.Println(dirErr)
						os.Exit(1)
					}
				} else {
					fileErr := os.Remove(args[i])
					if fileErr != nil {
						fmt.Println(fileErr)
						os.Exit(1)
					}
				}
			} else {
				fmt.Println(pathErr)
				os.Exit(1)
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
	rmCmd.Flags().BoolVarP(&rmFlg.file, "file", "f", false, "Attempt to remove the files without prompting for confirmation, regardless of the file's permissions.")
	rmCmd.Flags().BoolVarP(&rmFlg.confirm, "confirm", "i", false, "Request confirmation before attempting to remove each file, regardless of the file's permissions.")
	rmCmd.Flags().BoolVarP(&rmFlg.overwrite, "overwrite", "P", false, "Overwrite regular files before deleting them.")
	rmCmd.Flags().BoolVarP(&rmFlg.hierarchy, "hierarchy", "R", false, "Attempt to remove the file hierarchy rooted in each file argument.")
	rmCmd.Flags().BoolVarP(&rmFlg.subdir, "subdir", "r", false, "Equivalent to -R.")
	rmCmd.Flags().BoolVarP(&rmFlg.verbose, "verbose", "v", false, "Be verbose when deleting files, showing them as they are removed.")
	rmCmd.Flags().BoolVarP(&rmFlg.whiteout, "whiteout", "W", false, "Attempt to undelete the named files and recover files covered by whiteouts.")
}
