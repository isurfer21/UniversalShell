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
	i18nArchiveCmdTitle  = "Create and extract compressed files"
	i18nArchiveCmdDetail = `
Create and extract compressed files

It could create and extract .zip, .tar, .tar.gz, .tar.bz2, .tar.xz, .tar.lz4, 
.tar.sz, and .rar (extract-only) files.

Supported formats are zip, tar, tgz, tbz2, txz, tlz4, tsz, rar.

Syntax:
  ush archive -c -f=zip <archive-filename> <input-directory>
  ush archive -c -f=zip <archive-filename> <input-files...>
  ush archive -d -f=zip <archive-filename> <output-directory>

Example:
  ush archive -d test.zip test
  ush archive -c test.zip test/file1.exe test/file2.exe
  ush archive -d test.zip ""   ... decompresses in current folder
`
	i18nArchiveTplActionMissing = `
Action flag for compress/decompress is missing
`
	i18nArchiveTplMultipleDestination = `
Multiple output destination is not supported
`
	i18nArchiveTplParamMissing = `
Essential parameters are missing
`
)

type ArchiveLib struct {
}

func (archive *ArchiveLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (archive *ArchiveLib) compress(filename string, items []string) {
	err := archiver.Zip.Make(filename, items)
	archive.handleError(err)
}

func (archive *ArchiveLib) decompress(filename string, output string) {
	err := archiver.Zip.Open(filename, output)
	archive.handleError(err)
}

type archiveFlag struct {
	compress   bool
	decompress bool
	format     string
}

var (
	archiveFlg archiveFlag
	archiveLib ArchiveLib
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: i18nArchiveCmdTitle,
	Long:  i18nArchiveCmdDetail,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			filename := args[0]
			if archiveFlg.compress {
				items := args[1:]
				fmt.Printf("Compressing %s into '%s' \n", items, filename)
				archiveLib.compress(filename, items)
			} else if archiveFlg.decompress {
				if len(args) > 2 {
					fmt.Errorf(i18nArchiveTplMultipleDestination)
				} else {
					output := args[1]
					fmt.Printf("Decompressing '%s' at '%s' \n", filename, output)
					archiveLib.decompress(filename, output)
				}
			} else {
				fmt.Errorf(i18nArchiveTplActionMissing)
			}
		} else {
			fmt.Errorf(i18nArchiveTplParamMissing)
		}
	},
}

func init() {
	rootCmd.AddCommand(archiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// archiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	archiveCmd.Flags().BoolVarP(&archiveFlg.compress, "compress", "c", false, "Compress files or directories")
	archiveCmd.Flags().BoolVarP(&archiveFlg.decompress, "decompress", "d", false, "Decompress files or directories")
	archiveCmd.Flags().StringVarP(&archiveFlg.format, "format", "f", "zip", "Formats: zip, tar, tgz, tbz2, txz, tlz4, tsz, rar")
}
