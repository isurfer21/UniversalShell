/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

const (
	i18nUnrarCmdTitle  = "Extract files from a rar archive"
	i18nUnrarCmdDetail = `
Extract files from a rar archive

It could unpack rar archives and extract files from archives of the .rar format. 
The destination path is optional; default is current directory.

Syntax:
  ush unrar <archive-filename> <output-directory>

Example:
  # Extracts in 'test' directory
  ush unrar test.rar test

  # Extracts in current working directory
  ush unrar test.rar ""
  ush unrar test.rar
`
)

type UnrarLib struct {
}

func (unrar *UnrarLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (unrar *UnrarLib) decompress(filename string, output string) {
	err := archiver.Rar.Open(filename, output)
	unrar.handleError(err)
}

type UnrarFlag struct {
}

var (
	unrarFlg UnrarFlag
	unrarLib UnrarLib
)

// unrarCmd represents the unrar command
var unrarCmd = &cobra.Command{
	Use:   "unrar",
	Short: i18nUnrarCmdTitle,
	Long:  i18nUnrarCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		output := ""
		if len(args) > 1 {
			output = args[1]
		}
		fmt.Printf("Extracting '%s' at '%s' \n", filename, output)
		unrarLib.decompress(filename, output)
	},
}

func init() {
	rootCmd.AddCommand(unrarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// unrarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// unrarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
