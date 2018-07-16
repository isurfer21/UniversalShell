/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	i18nUniqCmdTitle  = "Uniquify files"
	i18nUniqCmdDetail = `
Uniquify files

Report or filter out repeated lines in a file. 

Reads standard input comparing adjacent lines, and writes a copy of each unique
input line to the standard output. The second and succeeding copies of 
identical adjacent input lines are not written. 
`
)

type UniqLib struct {
}

func (uniq *UniqLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (uniq *UniqLib) readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	uniq.handleError(err)
	return fmt.Sprintf("%s", content)
}

func (uniq *UniqLib) processLines(content string) string {
	lines := strings.Split(content, "\n")
	output := ""
	counter := 1
	for i := 0; i < len(lines)-1; i += 1 {
		isRepeated := false
		if uniqFlg.ignorecase {
			isRepeated = strings.ToLower(lines[i]) == strings.ToLower(lines[i+1])
		} else {
			isRepeated = (lines[i] == lines[i+1])
		}
		if isRepeated {
			counter += 1
		} else {
			if uniqFlg.count {
				output += fmt.Sprintf("   %d %s\n", counter, lines[i])
			} else if uniqFlg.repeated {
				if counter > 1 {
					output += fmt.Sprintf("%s\n", lines[i])
				}
			} else if uniqFlg.unique {
				if counter == 1 {
					output += fmt.Sprintf("%s\n", lines[i])
				}
			} else {
				output += fmt.Sprintf("%s\n", lines[i])
			}
			counter = 1
		}
	}
	return output
}

type UniqFlag struct {
	count      bool
	ignorecase bool
	repeated   bool
	unique     bool
	skipchars  int
	checkchars int
}

var (
	uniqFlg UniqFlag
	uniqLib UniqLib
)

// uniqCmd represents the uniq command
var uniqCmd = &cobra.Command{
	Use:   "uniq",
	Short: i18nUniqCmdTitle,
	Long:  i18nUniqCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			output := ""
			content := uniqLib.readFile(args[0])
			output = uniqLib.processLines(content)
			fmt.Printf(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(uniqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// uniqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	uniqCmd.Flags().BoolVarP(&uniqFlg.count, "count", "c", false, "print the number of times each line occurred along with the line")
	uniqCmd.Flags().BoolVarP(&uniqFlg.ignorecase, "ignore-case", "i", false, "ignore differences in case when comparing lines")
	uniqCmd.Flags().BoolVarP(&uniqFlg.repeated, "repeated", "d", false, "print only duplicate lines")
	uniqCmd.Flags().BoolVarP(&uniqFlg.unique, "unique", "u", false, "print only lines that are unique in the INPUT")
	// uniqCmd.Flags().IntVarP(&uniqFlg.skipchars, "skip-chars", "s", 0, "skip N characters before checking for uniqueness")
	// uniqCmd.Flags().IntVarP(&uniqFlg.checkchars, "check-chars", "w", 0, "compare N characters on each line")
}
