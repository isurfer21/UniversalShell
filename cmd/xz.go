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
	i18nXzCmdTitle  = "Compress or decompress .xz files"
	i18nXzCmdDetail = `
Compress or decompress .xz files

It could compress or decompress directories and files together in .xz formats.

Syntax:
  ush xz -z <archive-filename> <input-directory>
  ush xz -z <archive-filename> <input-files...>
  ush xz -d <archive-filename> <output-directory>

Example:  
  # Compresses the content of 'test' folder into xz file 
  ush xz -z test.xz test

  # Compresses the content of current directory into xz file 
  ush xz -z test.xz ./
  ush xz -z test.xz 

  # Compresses the listed files into xz file 
  ush xz -z test.xz test/file1.exe test/file2.exe

  # Decompresses in 'test' directory
  ush xz -d test.xz test

  # Decompresses in current working directory
  ush xz -d test.xz ""
  ush xz -d test.xz
`
	i18nXzTplActionMissing = `
Action flag for compress/decompress is missing
`
)

type XzLib struct {
}

func (xz *XzLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (xz *XzLib) compress(filename string, items []string) {
	var err error
	err = archiver.TarXZ.Make(filename, items)
	xz.handleError(err)
}

func (xz *XzLib) decompress(filename string, output string) {
	var err error
	err = archiver.TarXZ.Open(filename, output)
	xz.handleError(err)
}

type XzFlag struct {
	compress   bool
	decompress bool
}

var (
	xzFlg XzFlag
	xzLib XzLib
)

// xzCmd represents the xz command
var xzCmd = &cobra.Command{
	Use:   "xz",
	Short: i18nXzCmdTitle,
	Long:  i18nXzCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		if xzFlg.compress {
			items := args[1:]
			fmt.Printf("Compressing %s into '%s' \n", items, filename)
			xzLib.compress(filename, items)
		} else if xzFlg.decompress {
			output := ""
			if len(args) > 1 {
				output = args[1]
			}
			fmt.Printf("Decompressing '%s' at '%s' \n", filename, output)
			xzLib.decompress(filename, output)
		} else {
			fmt.Errorf(i18nXzTplActionMissing)
		}
	},
}

func init() {
	rootCmd.AddCommand(xzCmd)

	// Here you will define your flags and configuration settings.
	xzCmd.Flags().BoolVarP(&xzFlg.compress, "compress", "z", false, "force compression")
	xzCmd.Flags().BoolVarP(&xzFlg.decompress, "decompress", "d", false, "force decompression")
}
