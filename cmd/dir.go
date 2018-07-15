/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"

	"../lib"
)

const (
	i18nDirCmdTitle  = "Briefly list directory contents"
	i18nDirCmdDetail = `
Briefly list directory contents

Equivalent to 'ls -C'; that is, by default files are listed in columns, sorted 
vertically.
`
)

type DirLib struct {
	folder lib.Folder
}

func (dir *DirLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (dir *DirLib) padWithSpace(text string, limit int) string {
	pad := ""
	absLmt := limit
	if limit < 0 {
		absLmt = -limit
	}
	remain := absLmt - len(text)
	for i := 0; i < remain; i += 1 {
		pad += " "
	}
	output := ""
	if limit < 0 {
		output = pad + text
	} else {
		output = text + pad
	}
	return output
}

func (dir *DirLib) listFileInGrid(list []os.FileInfo) string {
	output := ""
	colSize := dir.maxTextLen(list) + 2
	colCount := tm.Width() / colSize
	index := 0
	for _, file := range list {
		if dirFlg.all || !dir.isDotEntry(file) {
			fileName := file.Name()
			output += dir.padWithSpace(fileName, colSize)
			if (index % colCount) == colCount-1 {
				output += "\n"
			}
			index += 1
		}
	}
	return output
}

func (dir *DirLib) maxTextLen(list []os.FileInfo) int {
	max := 0
	span := 0
	for _, file := range list {
		span = len(file.Name())
		if max < span {
			max = span
		}
	}
	return max
}

func (dir *DirLib) isDotEntry(file os.FileInfo) bool {
	fileName := file.Name()
	if string(fileName[0]) == "." {
		return true
	}
	return false
}

func (dir *DirLib) list(location string) string {
	items, err := dir.folder.ReadDir(location)
	dir.handleError(err)
	return dir.listFileInGrid(items)
}

type DirFlag struct {
	all bool
}

var (
	dirFlg DirFlag
	dirLib DirLib
)

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: i18nDirCmdTitle,
	Long:  i18nDirCmdDetail,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			location := pwd
			if len(args) > 0 {
				location = args[0]
			}
			fmt.Println(dirLib.list(location))
		}
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)

	// Here you will define your flags and configuration settings.
	dirCmd.Flags().BoolVarP(&dirFlg.all, "all", "a", false, "list all entries except for . and ..")
}
