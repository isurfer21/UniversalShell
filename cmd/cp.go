/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/isurfer21/UniversalShell/lib"
)

const (
	i18nCpCmdTitle  = "Copy files and directories"
	i18nCpCmdDetail = `
Copy files and directories

It copies a file from src to dst. If src and dst files exist, and are the same, 
then return success. Otherise, attempt to create a hard link between the two 
files. If that fail, copies the contents of the file named src to the file 
named by dst. The file will be created if it does not already exist. If the 
destination file exists, all it's contents will be replaced by the contents of 
the source file.
`
)

type CpLib struct {
	dossier lib.Dossier
	folder  lib.Folder
}

func (cp *CpLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (cp *CpLib) copy(src string, dest string) {
	path, err := os.Stat(src)
	cp.handleError(err)
	if path.IsDir() {
		cp.handleError(cp.folder.CopyDir(src, dest))
	} else {
		cp.handleError(cp.dossier.CopyFile(src, dest))
	}
}

type CpFlag struct {
}

var (
	cpFlg CpFlag
	cpLib CpLib
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: i18nCpCmdTitle,
	Long:  i18nCpCmdDetail,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			cpLib.copy(args[0], args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
