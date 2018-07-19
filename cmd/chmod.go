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

func (chmod *ChmodLib) isNumericMode(mode string) bool {
	_, err := strconv.Atoi(mode)
	if err == nil {
		return true
	}
	return false
}

func (chmod *ChmodLib) toNumericMode(mode string) int {
	num := 0700
	return num
}

func (chmod *ChmodLib) isValidMode(mode string) bool {
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
		// filePath := args[1]
		newMode := args[0]
		if chmodLib.isValidMode(newMode) {
			var numericMode int
			if chmodLib.isNumericMode(newMode) {
				numericMode, _ := strconv.ParseInt(newMode, 8, 64)
				fmt.Printf("Num: %o\n", numericMode)
			} else {
				numericMode = chmodLib.toNumericMode(newMode)
				fmt.Printf("Sym: %d\n", numericMode)
			}
			// chmodLib.handleError(os.Chmod(filePath, os.FileMode(numericMode)))
		} else {
			fmt.Printf(i18nChmodTplInvalidMode, newMode)
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
