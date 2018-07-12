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
	i18nUnzipCmdTitle  = "Unpacks zip archives"
	i18nUnzipCmdDetail = `
Unpacks zip archives

It could unpack zip archives and extract files from archives of the .zip format. 
The destination path is optional; default is current directory.

Syntax:
  ush unzip <archive-filename> <output-directory>

Example:
  # Extracts in 'test' directory
  ush unzip test.zip test

  # Extracts in current working directory
  ush unzip test.zip ""
  ush unzip test.zip
`
)

type UnzipLib struct {
}

func (unzip *UnzipLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (unzip *UnzipLib) decompress(filename string, output string) {
	err := archiver.Zip.Open(filename, output)
	unzip.handleError(err)
}

type UnzipFlag struct {
}

var (
	unzipFlg UnzipFlag
	unzipLib UnzipLib
)

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: i18nUnzipCmdTitle,
	Long:  i18nUnzipCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		output := ""
		if len(args) > 1 {
			output = args[1]
		}
		fmt.Printf("Extracting '%s' at '%s' \n", filename, output)
		unzipLib.decompress(filename, output)
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// unzipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// unzipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
