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
	i18nTailCmdTitle  = "Output the last part of file"
	i18nTailCmdDetail = `
Output the last part of file

Output the last part of files, print the last part (10 lines by default) of each FILE.
Priority of option 'quiet' is more than 'verbose'.
`
)

type TailLib struct {
}

func (tail *TailLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (tail *TailLib) readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	tail.handleError(err)
	return fmt.Sprintf("%s", content)
}

func (tail *TailLib) printLastKLines(count int, content string) string {
	lines := strings.Split(content, "\n")
	output := ""
	tailCnt := 0
	if count < len(lines) {
		tailCnt = len(lines) - count
	}
	for i := tailCnt; i < len(lines); i += 1 {
		output += fmt.Sprintf("%s\n", lines[i])
	}
	return output
}

type TailFlag struct {
	lines   int
	quiet   bool
	verbose bool
}

var (
	tailFlg TailFlag
	tailLib TailLib
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: i18nTailCmdTitle,
	Long:  i18nTailCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			output := ""
			for i := 0; i < len(args); i += 1 {
				content := tailLib.readFile(args[i])
				content = tailLib.printLastKLines(tailFlg.lines, content)
				if (len(args) > 1 && !tailFlg.quiet) || (tailFlg.verbose && !tailFlg.quiet) {
					content = fmt.Sprintf("==> %s <==\n%s", args[i], content)
				}
				output += content
			}
			fmt.Printf(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// tailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	tailCmd.Flags().IntVarP(&tailFlg.lines, "lines", "n", 10, "output the last K lines")
	tailCmd.Flags().BoolVarP(&tailFlg.quiet, "quiet", "q", false, "never output headers giving file names")
	tailCmd.Flags().BoolVarP(&tailFlg.verbose, "verbose", "v", false, "always output headers giving file names")
}
