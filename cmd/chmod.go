/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
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

func (chmod *ChmodLib) permissions() map[string]map[string]int {
	p := map[string]map[string]int{
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
	return p
}

func (chmod *ChmodLib) printFileInfo(path string) {
	fi, err := os.Lstat(path)
	if err == nil {
		fmt.Printf("Info: %s %s\n", fi.Mode().String(), path)
	}
}

func (chmod *ChmodLib) toNumericMode(l string, r string) int {
	s := 0
	p := chmod.permissions()
	if len(l) > 1 {
		for _, lv := range l {
			for _, rv := range r {
				s += p[string(lv)][string(rv)]
			}
		}
	} else {
		for _, v := range r {
			s += p[string(l)][string(v)]
		}
	}
	return s
}

func (chmod *ChmodLib) toAbsoluteMode(mode string, path string) int {
	absmode := 0
	sum := 0
	if strings.Contains(mode, "a") {
		mode = strings.Replace(mode, "a", "ugo", -1)
	}
	fmt.Printf("mode: %s \n", mode)
	tokens := strings.Split(mode, ",")
	for i := 0; i < len(tokens); i += 1 {
		if strings.Contains(tokens[i], "+") {
			c := strings.Split(tokens[i], "+")
			if len(c[0]) == 0 {
				c[0] = "u"
			}
			sum += chmod.toNumericMode(c[0], c[1])
		} else if strings.Contains(tokens[i], "-") {
			c := strings.Split(tokens[i], "-")
			if len(c[0]) == 0 {
				c[0] = "u"
			}
			sum -= chmod.toNumericMode(c[0], c[1])
		} else if strings.Contains(tokens[i], "=") {
			c := strings.Split(tokens[i], "=")
			if len(c[0]) == 0 {
				c[0] = "ugo"
			}
			sum += chmod.toNumericMode(c[0], c[1])
		}
	}
	if strings.Contains(mode, "+") || strings.Contains(mode, "-") {
		fi, err := os.Lstat(path)
		chmod.handleError(err)
		fmt.Printf("%v -> %#o + %#o \n", fi.Mode(), int(fi.Mode()), sum)
		absmode = int(fi.Mode()) + sum
		fmt.Printf("%#o \n", absmode)
	}
	if strings.Contains(mode, "=") {
		absmode = sum
	}
	return absmode
}

func (chmod *ChmodLib) isSymbolicMode(mode string) bool {
	isValid, err := regexp.MatchString("[ugoa]*([-+=]([rwxXst]*|[ugo]))+|[-+=][0-7]+", mode)
	chmod.handleError(err)
	return isValid
}

type ChmodFlag struct {
	silent    bool
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
					fmt.Printf("Mode: %#o (%d)\n", absoluteMode, absoluteMode)
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
	chmodCmd.Flags().BoolVarP(&chmodFlg.verbose, "verbose", "v", false, "output a diagnostic for every file processed")
	chmodCmd.Flags().BoolVarP(&chmodFlg.changes, "changes", "c", false, "like verbose but report only when a change is made")
	chmodCmd.Flags().BoolVarP(&chmodFlg.recursive, "recursive", "R", false, "change files and directories recursivel")

}
