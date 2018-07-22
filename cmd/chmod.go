/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	// "bytes"
	// "encoding/binary"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	i18nChmodCmdTitle  = "Change access permissions, change mode."
	i18nChmodCmdDetail = `
Change access permissions, change mode.

It changes the permissions of each given file according to mode, where mode
describes the permissions to modify. Mode can be specified with octal numbers
or with letters. Using letters is easier to understand for most people.
`
	i18nChmodTplInvalidMode = "Invalid mode: %s\n"
)

type ChmodLib struct {
}

func (chmod *ChmodLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (chmod *ChmodLib) isNumber(mode string) bool {
	_, err := strconv.Atoi(mode)
	if err == nil {
		return true
	}
	return false
}

func (chmod *ChmodLib) printFileInfo(path string) {
	fi, err := os.Lstat(path)
	if err == nil {
		fmt.Printf("Info: %s %s\n", fi.Mode().String(), path)
	}
}

func (chmod *ChmodLib) replacePermission(mode string) int {
	perm := map[string]map[string]int{
		"u": {
			"r": 00400,
			"w": 00200,
			"x": 00100,
			"s": 04000,
			"t": 01000,
		},
		"g": {
			"r": 00040,
			"w": 00020,
			"x": 00010,
			"s": 02000,
		},
		"o": {
			"r": 00004,
			"w": 00002,
			"x": 00001,
		},
	}
	sum := 0
	tokens := strings.Split(mode, ",")
	for i := 0; i < len(tokens); i += 1 {
		var separator string
		if strings.Contains(tokens[i], "=") {
			separator = "="
		}
		if strings.Contains(tokens[i], "+") {
			separator = "+"
		}
		if strings.Contains(tokens[i], "-") {
			separator = "-"
		}
		c := strings.Split(tokens[i], separator)
		l, r := c[0], c[1]
		if len(l) == 0 {
			l = "ugo"
		}
		for _, lv := range l {
			for _, rv := range r {
				sum += perm[string(lv)][string(rv)]
			}
		}
	}
	return sum
}

func (chmod *ChmodLib) modifyPermissionChunk(oldMode int, newMode string) int {
	perm := map[string]map[string]int{
		"u": {
			"r": 8,
			"w": 7,
			"x": 6,
			"s": 11,
			"t": 9,
		},
		"g": {
			"r": 5,
			"w": 4,
			"x": 3,
			"s": 10,
		},
		"o": {
			"r": 2,
			"w": 1,
			"x": 0,
		},
	}
	m := []rune(fmt.Sprintf("%#b", oldMode))
	var separator string
	var bit rune
	if strings.Contains(newMode, "+") {
		separator = "+"
		bit = 49
	} else if strings.Contains(newMode, "-") {
		separator = "-"
		bit = 48
	}
	c := strings.Split(newMode, separator)
	l, r := c[0], c[1]
	if len(l) == 0 {
		l = "ugo"
	}
	for _, lv := range l {
		for _, rv := range r {
			n := len(m) - 1
			p := perm[string(lv)][string(rv)]
			if n >= p {
				m[n-p] = bit
			}
		}
	}
	s, _ := strconv.ParseInt(string(m), 2, 64)
	return int(s)
}

// todo
func (chmod *ChmodLib) replacePermissionChunk(oldMode int, newMode string) int {
	perm := map[string]map[string]int{
		"u": {
			"r": 8,
			"w": 7,
			"x": 6,
			"s": 11,
			"t": 9,
		},
		"g": {
			"r": 5,
			"w": 4,
			"x": 3,
			"s": 10,
		},
		"o": {
			"r": 2,
			"w": 1,
			"x": 0,
		},
	}
	m := []rune(fmt.Sprintf("%#b", oldMode))
	c := strings.Split(newMode, "=")
	l, r := c[0], c[1]
	if len(l) == 0 {
		l = "ugo"
	}
	for _, lv := range l {
		for k := range perm[string(lv)] {
			n := len(m) - 1
			p := perm[string(lv)][string(k)]
			if n >= p {
				if strings.Contains(r, string(k)) {
					m[n-p] = 49
				} else {
					m[n-p] = 48
				}
			}
		}
	}
	s, _ := strconv.ParseInt(string(m), 2, 64)
	return int(s)
}

func (chmod *ChmodLib) toAbsoluteMode(mode string, path string) int {
	fi, err := os.Lstat(path)
	chmod.handleError(err)
	absMode := int(fi.Mode())
	if strings.Contains(mode, "a") {
		mode = strings.Replace(mode, "a", "ugo", -1)
	}
	if !chmodFlg.force {
		tokens := strings.Split(mode, ",")
		for i := 0; i < len(tokens); i += 1 {
			if strings.Contains(tokens[i], "+") || strings.Contains(tokens[i], "-") {
				absMode = chmod.modifyPermissionChunk(absMode, mode)
			} else if strings.Contains(mode, "=") {
				absMode = chmod.replacePermissionChunk(absMode, mode)
			}
		}
	} else {
		absMode = chmod.replacePermission(mode)
	}
	return absMode
}

func (chmod *ChmodLib) isSymbolicMode(mode string) bool {
	isValid, err := regexp.MatchString("[ugoa]*([-+=]([rwxXst]*|[ugo]))+|[-+=][0-7]+", mode)
	chmod.handleError(err)
	return isValid
}

type ChmodFlag struct {
	silent    bool
	force     bool
	verbose   bool
	changes   bool
	recursive bool
}

var (
	chmodFlg ChmodFlag
	chmodLib ChmodLib
)

// chmodCmd represents the chmod command
var chmodCmd = &cobra.Command{
	Use:   "chmod",
	Short: i18nChmodCmdTitle,
	Long:  i18nChmodCmdDetail,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[1]
		newMode := args[0]
		if chmodLib.isNumber(newMode) {
			absoluteMode, _ := strconv.ParseInt(newMode, 8, 64)
			if absoluteMode > 0 && absoluteMode <= 07777 {
				chmodLib.handleError(os.Chmod(filePath, os.FileMode(absoluteMode)))
				if chmodFlg.verbose {
					fmt.Printf("Mode: %#o \n", absoluteMode, absoluteMode)
					chmodLib.printFileInfo(filePath)
				}
			} else {
				fmt.Printf(i18nChmodTplInvalidMode, newMode)
			}
		} else {
			if chmodLib.isSymbolicMode(newMode) {
				absoluteMode := chmodLib.toAbsoluteMode(newMode, filePath)
				chmodLib.handleError(os.Chmod(filePath, os.FileMode(absoluteMode)))
				if chmodFlg.verbose {
					fmt.Printf("Mode: %s -> %#o (%d)\n", newMode, absoluteMode, absoluteMode)
					chmodLib.printFileInfo(filePath)
				}
			} else {
				fmt.Printf(i18nChmodTplInvalidMode, newMode)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(chmodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// chmodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	chmodCmd.Flags().BoolVarP(&chmodFlg.silent, "silent", "f", false, "suppress most error messages")
	chmodCmd.Flags().BoolVarP(&chmodFlg.force, "force", "F", false, "overwrite other permissions")
	chmodCmd.Flags().BoolVarP(&chmodFlg.verbose, "verbose", "v", false, "output a diagnostic for every file processed")
	chmodCmd.Flags().BoolVarP(&chmodFlg.changes, "changes", "c", false, "like verbose but report only when a change is made")
	chmodCmd.Flags().BoolVarP(&chmodFlg.recursive, "recursive", "R", false, "change files and directories recursively")

}
