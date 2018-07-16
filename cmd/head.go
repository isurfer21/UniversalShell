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
	i18nHeadCmdTitle  = "Output the first part of file(s)"
	i18nHeadCmdDetail = `
Output the first part of file(s)

Output the first part of files, prints the first part (10 lines by default) of each file.
Priority of option 'quiet' is more than 'verbose'.
`
)

type HeadLib struct {
}

func (head *HeadLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (head *HeadLib) readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	head.handleError(err)
	return fmt.Sprintf("%s", content)
}

func (head *HeadLib) printFirstNLines(count int, content string) string {
	lines := strings.Split(content, "\n")
	output := ""
	headCnt := len(lines)
	if count < len(lines) {
		headCnt = count
	}
	for i := 0; i < headCnt; i += 1 {
		output += fmt.Sprintf("%s\n", lines[i])
	}
	return output
}

type HeadFlag struct {
	lines   int
	quiet   bool
	verbose bool
}

var (
	headFlg HeadFlag
	headLib HeadLib
)

// headCmd represents the head command
var headCmd = &cobra.Command{
	Use:   "head",
	Short: i18nHeadCmdTitle,
	Long:  i18nHeadCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			output := ""
			for i := 0; i < len(args); i += 1 {
				content := headLib.readFile(args[i])
				content = headLib.printFirstNLines(headFlg.lines, content)
				if (len(args) > 1 && !headFlg.quiet) || (headFlg.verbose && !headFlg.quiet) {
					content = fmt.Sprintf("==> %s <==\n%s", args[i], content)
				}
				output += content
			}
			fmt.Printf(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(headCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// headCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	headCmd.Flags().IntVarP(&headFlg.lines, "lines", "n", 10, "output the first N lines")
	headCmd.Flags().BoolVarP(&headFlg.quiet, "quiet", "q", false, "never print file name headers")
	headCmd.Flags().BoolVarP(&headFlg.verbose, "verbose", "v", false, "always print file name headers")
}
