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
	i18nTarCmdTitle  = "Store, list or extract files in an archive"
	i18nTarCmdDetail = `
Store, list or extract files in an archive

It could compress or decompress directories and files together using tar, tgz, 
tbz2, txz, tlz4, tsz formats.

Syntax:
  ush tar -c -f <archive-filename> <input-directory>
  ush tar -c -f <archive-filename> <input-files...>
  ush tar -c -f <archive-filename> <output-directory>

Example:  
  # Archives the content of 'test' folder into tar file 
  ush tar -c -f test.tar test

  # Archives the content of current directory into tar file 
  ush tar -c -f test.tar ./
  ush tar -c -f test.tar 

  # Archives the listed files into tar file 
  ush tar -c -f test.tar test/file1.exe test/file2.exe

  # Extracts in 'test' directory
  ush tar -x -f test.tar test

  # Extracts in current working directory
  ush tar -x -f test.tar ""
  ush tar -x -f test.tar
`
	i18nTarTplActionMissing = `
Action flag for compress/decompress is missing
`
)

type TarLib struct {
}

func (tar *TarLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (tar *TarLib) compress(filename string, items []string) {
	err := archiver.Tar.Make(filename, items)
	tar.handleError(err)
}

func (tar *TarLib) decompress(filename string, output string) {
	err := archiver.Tar.Open(filename, output)
	tar.handleError(err)
}

type TarFlag struct {
	file    string
	create  bool
	extract bool
	gzip    bool
	bzip2   bool
	lzip4   bool
	xz      bool
	sz      bool
}

var (
	tarFlg TarFlag
	tarLib TarLib
)

// tarCmd represents the tar command
var tarCmd = &cobra.Command{
	Use:   "tar",
	Short: i18nTarCmdTitle,
	Long:  i18nTarCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := tarFlg.file
		if tarFlg.create {
			items := args[0:]
			fmt.Printf("Archiving %s into '%s' \n", items, filename)
			tarLib.compress(filename, items)
		} else if tarFlg.extract {
			output := ""
			if len(args) > 0 {
				output = args[0]
			}
			fmt.Printf("Extracting '%s' at '%s' \n", filename, output)
			tarLib.decompress(filename, output)
		} else {
			fmt.Errorf(i18nTarTplActionMissing)
		}
	},
}

func init() {
	rootCmd.AddCommand(tarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// tarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	tarCmd.Flags().BoolVarP(&tarFlg.create, "create", "c", false, "create a new archive")
	tarCmd.Flags().BoolVarP(&tarFlg.extract, "extract", "d", false, "extract files from an archive")
	tarCmd.Flags().StringVarP(&tarFlg.file, "file", "f", "tar", "use archive file or device")
	tarCmd.Flags().BoolVarP(&tarFlg.gzip, "gzip", "z", false, "filter the archive through gzip")
	tarCmd.Flags().BoolVarP(&tarFlg.bzip2, "bzip2", "j", false, "filter the archive through bzip2")
	tarCmd.Flags().BoolVarP(&tarFlg.lz4, "lz4", "l", false, "filter the archive through lz4")
	tarCmd.Flags().BoolVarP(&tarFlg.xz, "xz", "J", false, "filter the archive through xz")
	tarCmd.Flags().BoolVarP(&tarFlg.sz, "sz", "s", false, "filter the archive through sz")
}
