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

	"github.com/inhies/go-bytesize"
	"github.com/spf13/cobra"
)

const (
	i18nLsCmdTitle  = "List directory contents"
	i18nLsCmdDetail = `
List directory contents

It displays a list of files and sub-directories in a directory which could be 
rendered in various ways based on passed flags.
`
)

type lsFlag struct {
	horizontal bool
	vertical   bool
	tabulated  bool
	raw        bool
}

var lsFlg lsFlag

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
			items, err := readDir(location)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				separator := "  "
				list := []string{}
				for _, file := range items {
					if lsFlg.tabulated || lsFlg.raw {
						separator = "\n"

						fileMod := file.Mode().String()
						fileSize := "-"
						fileModTime := ""
						isDir := ""
						fileName := file.Name()
						var row []string

						if lsFlg.raw {
							if file.IsDir() {
								isDir = "Dir"
							}
							fileSize = strconv.FormatInt(file.Size(), 10)
							fileModTime = file.ModTime().String()
							row = []string{
								fileMod,
								fileSize,
								fileModTime,
								isDir,
								fileName,
							}
						}
						if lsFlg.tabulated {
							if file.IsDir() {
								var pathSeparator string
								if runtime.GOOS == "windows" {
									pathSeparator = "\\"
								} else {
									pathSeparator = "/"
								}
								fileName += pathSeparator
							} else {
								bytesize.Format = "%.1f"
								fileSize = bytesize.New(float64(file.Size())).String()
							}
							fileModTime = file.ModTime().Format("Jan 02, 2006 15:04")
							row = []string{
								fileMod,
								fileSize,
								fileModTime,
								fileName,
							}
						}
						list = append(list, strings.Join(row[:], "\t"))
					} else {
						if lsFlg.vertical {
							separator = "\n"
						}
						if lsFlg.horizontal {
							separator = " \t"
						}
						list = append(list, file.Name())
					}
				}
				fmt.Printf("%s\n", strings.Join(list[:], separator))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.
	lsCmd.Flags().BoolVarP(&lsFlg.horizontal, "horz", "x", false, "display the list horizontally")
	lsCmd.Flags().BoolVarP(&lsFlg.vertical, "vert", "y", false, "display the list vertically")
	lsCmd.Flags().BoolVarP(&lsFlg.tabulated, "long", "l", false, "display extended file metadata as a table")
	lsCmd.Flags().BoolVarP(&lsFlg.raw, "raw", "r", false, "display raw extended file metadata as a table")
}

func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}
