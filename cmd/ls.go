/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/inhies/go-bytesize"
	"github.com/spf13/cobra"

	"github.com/isurfer21/UniversalShell/lib"
)

const (
	i18nLsCmdTitle  = "List information about file(s)"
	i18nLsCmdDetail = `
List information about file(s)

It displays a list of files and sub-directories in a directory which could be 
rendered in various ways based on passed flags.
`
)

type LsLib struct {
	folder lib.Folder
}

func (ls *LsLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (ls *LsLib) isSymlink(file os.FileInfo) bool {
	return (file.Mode()&os.ModeSymlink != 0)
}

func (ls *LsLib) classify(file os.FileInfo) string {
	fileName := file.Name()
	if file.IsDir() {
		var pathSeparator string
		if runtime.GOOS == "windows" {
			pathSeparator = "\\"
		} else {
			pathSeparator = "/"
		}
		fileName += pathSeparator
	}
	return fileName
}

func (ls *LsLib) fmtFileSize(file os.FileInfo, fmt string) string {
	bytesize.Format = fmt
	return bytesize.New(float64(file.Size())).String()
}

func (ls *LsLib) padWithSpace(text string, limit int) string {
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

func (ls *LsLib) listFileInGrid(list []os.FileInfo) string {
	output := ""
	colSize := ls.maxTextLen(list, "Name") + 2
	colCount := tm.Width() / colSize
	index := 0
	for _, file := range list {
		if lsFlg.all || !ls.isDotEntry(file) {
			fileName := file.Name()
			output += ls.padWithSpace(fileName, colSize)
			if (index % colCount) == colCount-1 {
				output += "\n"
			}
			index += 1
		}
	}
	return output
}

func (ls *LsLib) listFileWithSizeInGrid(list []os.FileInfo) string {
	output := ""
	nameLen := ls.maxTextLen(list, "Name")
	sizeLen := ls.maxTextLen(list, "Size")
	colSize := nameLen + sizeLen + 3
	colCount := tm.Width() / colSize
	index := 0
	for _, file := range list {
		if lsFlg.all || !ls.isDotEntry(file) {
			fileName := ls.padWithSpace(file.Name(), nameLen)
			fileSize := ls.padWithSpace(ls.fmtFileSize(file, "%.0f"), -sizeLen)
			output += ls.padWithSpace(fmt.Sprintf("%s %s", fileSize, fileName), colSize)
			if (index % colCount) == colCount-1 {
				output += "\n"
			}
			index += 1
		}
	}
	return output
}

func (ls *LsLib) maxTextLen(list []os.FileInfo, prop string) int {
	max := 0
	span := 0
	for _, file := range list {
		if prop == "Name" {
			span = len(file.Name())
		} else if prop == "Size" {
			bytesize.Format = "%.0f"
			fileSize := bytesize.New(float64(file.Size())).String()
			span = len(fileSize)
		}
		if max < span {
			max = span
		}
	}
	return max
}

func (ls *LsLib) isDotEntry(file os.FileInfo) bool {
	fileName := file.Name()
	if string(fileName[0]) == "." {
		return true
	}
	return false
}

func (ls *LsLib) listFileInTable(file os.FileInfo) []string {
	fileName := ls.classify(file)
	fileMod := file.Mode().String()
	if ls.isSymlink(file) {
		link, err := os.Readlink(file.Name())
		ls.handleError(err)
		fileName += " -> " + link
	}
	fileSize := "-"
	isDir := ""
	if file.IsDir() {
		isDir = "Dir"
	} else {
		if lsFlg.raw {
			fileSize = strconv.FormatInt(file.Size(), 10)
		} else {
			fileSize = ls.fmtFileSize(file, "%.1f")
		}
	}
	fileModTime := ""
	if lsFlg.time {
		fileModTime = file.ModTime().String()
	} else {
		fileModTime = file.ModTime().Format("Jan 02, 2006 15:04")
	}
	row := []string{}
	if lsFlg.raw {
		row = []string{
			fileMod,
			fileSize,
			fileModTime,
			isDir,
			fileName,
		}
	} else {
		row = []string{
			fileMod,
			fileSize,
			fileModTime,
			fileName,
		}
	}
	return row
}

func (ls *LsLib) dir(location string) string {
	items, err := ls.folder.ReadDir(location)
	ls.handleError(err)

	separator := "  "
	list := []string{}
	if lsFlg.column {
		return ls.listFileInGrid(items)
	} else if lsFlg.size {
		return ls.listFileWithSizeInGrid(items)
	} else {
		for _, file := range items {
			if lsFlg.all || !ls.isDotEntry(file) {
				if lsFlg.tabulated {
					separator = "\n"
					row := ls.listFileInTable(file)
					list = append(list, strings.Join(row[:], "\t"))
				} else {
					fileName := file.Name()
					if lsFlg.vertical {
						separator = "\n"
					}
					if lsFlg.horizontal {
						separator = " \t"
					}
					if lsFlg.csv {
						separator = ", "
					}
					list = append(list, fileName)
				}
			}
		}
		return strings.Join(list[:], separator)
	}
}

func (ls *LsLib) exist(location string) bool {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return false
	}
	return true
}

type LsFlag struct {
	all        bool
	almostall  bool
	classify   bool
	column     bool
	csv        bool
	exist      bool
	horizontal bool
	tabulated  bool
	raw        bool
	size       bool
	time       bool
	vertical   bool
}

var (
	lsFlg LsFlag
	lsLib LsLib
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: i18nLsCmdTitle,
	Long:  i18nLsCmdDetail,
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
			if lsFlg.exist {
				fmt.Println(strconv.FormatBool(lsLib.exist(location)))
			} else {
				fmt.Println(lsLib.dir(location))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.
	lsCmd.Flags().BoolVarP(&lsFlg.all, "all", "a", false, "list all entries including those starting with a dot .")
	// lsCmd.Flags().BoolVarP(&lsFlg.almostall, "almost-all", "A", false, "list all entries except for . and ..")
	lsCmd.Flags().BoolVarP(&lsFlg.classify, "classify", "F", false, "append indicator (one of */=@|) to entries")
	lsCmd.Flags().BoolVarP(&lsFlg.column, "column", "C", false, "list entries by columns (vertical)")
	lsCmd.Flags().BoolVarP(&lsFlg.csv, "csv", "m", false, "fill width with a comma separated list of entries")
	lsCmd.Flags().BoolVarP(&lsFlg.exist, "exist", "e", false, "returns true/false based on path existence")
	lsCmd.Flags().BoolVarP(&lsFlg.horizontal, "line", "x", false, "list entries by lines (horizontal)")
	lsCmd.Flags().BoolVarP(&lsFlg.tabulated, "long", "l", false, "use a long listing format")
	lsCmd.Flags().BoolVarP(&lsFlg.raw, "raw", "r", false, "display raw extended file metadata in a table")
	lsCmd.Flags().BoolVarP(&lsFlg.size, "size", "s", false, "print size of each file, in blocks")
	lsCmd.Flags().BoolVarP(&lsFlg.time, "time", "T", false, "display complete time information for the file")
	lsCmd.Flags().BoolVarP(&lsFlg.vertical, "vertical", "1", false, "list one file per line")
}
