/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	i18nCatCmdTitle  = "Concatenate and print (display) the content of files"
	i18nCatCmdDetail = `
Concatenate and print (display) the content of files

Concatenate FILE(s), or standard input, to standard output. With no FILE, or 
when FILE is -, read standard input.
`
)

type CatLib struct {
}

func (cat *CatLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (cat *CatLib) readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	cat.handleError(err)
	return fmt.Sprintf("%s", content)
}

func (cat *CatLib) writeFile(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	cat.handleError(err)
}

func (cat *CatLib) createFile(filename string) {
	file, err := os.Create(filename)
	cat.handleError(err)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		_, err := file.WriteString(scanner.Text() + "\n")
		cat.handleError(err)
		file.Sync()
	}
	defer file.Close()
}

func (cat *CatLib) prefixLineNumber(content string) string {
	lines := strings.Split(content, "\n")
	output := ""
	for i := 0; i < len(lines)-1; i += 1 {
		output += fmt.Sprintf("     %d  %s\n", i+1, lines[i])
	}
	return output
}

func (cat *CatLib) prefixLineNumberExcludingBlankLine(content string) string {
	lines := strings.Split(content, "\n")
	output := ""
	for i := 0; i < len(lines)-1; i += 1 {
		if lines[i] == "" {
			output += fmt.Sprintf("%s\n", lines[i])
		} else {
			output += fmt.Sprintf("     %d  %s\n", i+1, lines[i])
		}
	}
	return output
}

func (cat *CatLib) suffixDollarAtLineEnd(content string) string {
	lines := strings.Split(content, "\n")
	output := ""
	for i := 0; i < len(lines)-1; i += 1 {
		output += fmt.Sprintf("%s$\n", lines[i])
	}
	return output
}

func (cat *CatLib) replaceTabWithTag(content string) string {
	return strings.Replace(content, "\t", "^I", -1)
}

type CatFlag struct {
	input           string
	output          string
	append          string
	number          bool
	showeof         bool
	showends        bool
	numbernonblank  bool
	showtabs        bool
	shownonprinting bool
}

var (
	catFlg CatFlag
	catLib CatLib
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: i18nCatCmdTitle,
	Long:  i18nCatCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			output := ""
			for i := 0; i < len(args); i += 1 {
				content := catLib.readFile(args[i])
				if catFlg.number {
					content = catLib.prefixLineNumber(content)
				}
				if catFlg.numbernonblank {
					content = catLib.prefixLineNumberExcludingBlankLine(content)
				}
				if catFlg.showends {
					content = catLib.suffixDollarAtLineEnd(content)
				}
				if catFlg.showtabs {
					content = catLib.replaceTabWithTag(content)
				}
				if catFlg.showeof {
					content += "--\n"
				}
				output += content
			}
			if catFlg.output != "" {
				catLib.writeFile(catFlg.output, output)
			} else {
				fmt.Print(output)
			}
		} else {
			if catFlg.output != "" {
				catLib.createFile(catFlg.output)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	catCmd.Flags().StringVarP(&catFlg.input, "input", "i", "", "input all output lines")
	catCmd.Flags().StringVarP(&catFlg.output, "output", "o", "", "output all output lines")
	catCmd.Flags().StringVarP(&catFlg.append, "append", "a", "", "append all output lines")

	catCmd.Flags().BoolVarP(&catFlg.number, "number", "n", false, "number all output lines")
	catCmd.Flags().BoolVarP(&catFlg.showeof, "show-eof", "e", false, "display -- at end of each file")
	catCmd.Flags().BoolVarP(&catFlg.showends, "show-ends", "E", false, "display $ at end of each line")
	catCmd.Flags().BoolVarP(&catFlg.numbernonblank, "number-nonblank", "b", false, "number nonblank output lines")
	catCmd.Flags().BoolVarP(&catFlg.showtabs, "show-tabs", "T", false, "display TAB characters as ^I")
	// catCmd.Flags().BoolVarP(&catFlg.shownonprinting, "show-nonprinting", "v", false, "use ^ and M- notation, except for LFD and TAB")

}
