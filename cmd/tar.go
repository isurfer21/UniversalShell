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
  ush tar -x -f <archive-filename> <output-directory>

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

  # Archives the source directory into various compression formats 
  ush tar -c -z -f test.gz test
  ush tar -c -j -f test.bz2 test
  ush tar -c -J -f test.xz test
  ush tar -c -s -f test.sz test
  ush tar -c -l -f test.lz4 test

  # Extracts files of various compression formats in 'test' directory 
  ush tar -x -z -f test.gz test
  ush tar -x -j -f test.bz2 test
  ush tar -x -J -f test.xz test
  ush tar -x -s -f test.sz test
  ush tar -x -l -f test.lz4 test
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
	var err error
	if tarFlg.gzip {
		err = archiver.TarGz.Make(filename, items)
	} else if tarFlg.bzip2 {
		err = archiver.TarBz2.Make(filename, items)
	} else if tarFlg.lz4 {
		err = archiver.TarLz4.Make(filename, items)
	} else if tarFlg.xz {
		err = archiver.TarXZ.Make(filename, items)
	} else if tarFlg.sz {
		err = archiver.TarSz.Make(filename, items)
	} else {
		err = archiver.Tar.Make(filename, items)
	}
	tar.handleError(err)
}

func (tar *TarLib) decompress(filename string, output string) {
	var err error
	if tarFlg.gzip {
		err = archiver.TarGz.Open(filename, output)
	} else if tarFlg.bzip2 {
		err = archiver.TarBz2.Open(filename, output)
	} else if tarFlg.lz4 {
		err = archiver.TarLz4.Open(filename, output)
	} else if tarFlg.xz {
		err = archiver.TarXZ.Open(filename, output)
	} else if tarFlg.sz {
		err = archiver.TarSz.Open(filename, output)
	} else {
		err = archiver.Tar.Open(filename, output)
	}
	tar.handleError(err)
}

type TarFlag struct {
	file    string
	create  bool
	extract bool
	gzip    bool
	bzip2   bool
	lz4     bool
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
	tarCmd.Flags().StringVarP(&tarFlg.file, "file", "f", "tar", "use archive file or device")
	tarCmd.Flags().BoolVarP(&tarFlg.create, "create", "c", false, "create a new archive")
	tarCmd.Flags().BoolVarP(&tarFlg.extract, "extract", "d", false, "extract files from an archive")
	tarCmd.Flags().BoolVarP(&tarFlg.gzip, "gzip", "z", false, "filter the archive through gzip (.gz)")
	tarCmd.Flags().BoolVarP(&tarFlg.bzip2, "bzip2", "j", false, "filter the archive through bzip2 (.bz2)")
	tarCmd.Flags().BoolVarP(&tarFlg.lz4, "lz4", "l", false, "filter the archive through lz4")
	tarCmd.Flags().BoolVarP(&tarFlg.xz, "xz", "J", false, "filter the archive through xz")
	tarCmd.Flags().BoolVarP(&tarFlg.sz, "sz", "s", false, "filter the archive through sz")
}
