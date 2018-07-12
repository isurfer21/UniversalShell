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
	i18nZipCmdTitle  = "Package and compress (archive) files"
	i18nZipCmdDetail = `
Package and compress (archive) files

It could package and compress directories and files together in a zipped file 
having .zip extension 

Syntax:
  ush zip <archive-filename> <input-directory>
  ush zip <archive-filename> <input-files...>

Example:
  # Archives the content of 'test' folder into zip file 
  ush zip test.zip test

  # Archives the content of current directory into zip file 
  ush zip test.zip ./
  ush zip test.zip 

  # Archives the listed files into zip file 
  ush zip test.zip test/file1.exe test/file2.exe
`
)

type ZipLib struct {
}

func (zip *ZipLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (zip *ZipLib) compress(filename string, items []string) {
	err := archiver.Zip.Make(filename, items)
	zip.handleError(err)
}

type ZipFlag struct {
}

var (
	zipFlg ZipFlag
	zipLib ZipLib
)

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: i18nZipCmdTitle,
	Long:  i18nZipCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		items := args[1:]
		fmt.Printf("Archiving %s into '%s' \n", items, filename)
		zipLib.compress(filename, items)
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// zipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// zipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
