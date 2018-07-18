/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	// "bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"os"
	// "strings"

	"github.com/spf13/cobra"
)

const (
	i18nSedCmdTitle  = "Stream Editor"
	i18nSedCmdDetail = `
Stream Editor

It is a stream editor. A stream editor is used to perform basic text 
transformations on an input stream (a file or input from a pipeline). 
While in some ways similar to an editor which permits scripted edits, SED works
by making only one pass over the input(s), and is consequently more efficient.
But it is SED's ability to filter text in a pipeline which particularly 
distinguishes it from other types of editors.
`
)

type SedLib struct {
}

func (sed *SedLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (sed *SedLib) readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	sed.handleError(err)
	return fmt.Sprintf("%s", content)
}

func (sed *SedLib) writeFile(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	sed.handleError(err)
}

func (sed *SedLib) adjoinFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	sed.handleError(err)
	defer file.Close()
	_, err = file.WriteString(content)
	sed.handleError(err)
}

type SedFlag struct {
	input      string
	output     string
	append     string
	expression string
}

var (
	sedFlg SedFlag
	sedLib SedLib
)

// sedCmd represents the sed command
var sedCmd = &cobra.Command{
	Use:   "sed",
	Short: i18nSedCmdTitle,
	Long:  i18nSedCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(sedFlg.input, sedFlg.output, sedFlg.append)
	},
}

func init() {
	rootCmd.AddCommand(sedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// sedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	sedCmd.Flags().StringVarP(&sedFlg.input, "input", "i", "", "input all output lines")
	sedCmd.Flags().StringVarP(&sedFlg.output, "output", "o", "", "output all output lines")
	sedCmd.Flags().StringVarP(&sedFlg.append, "append", "a", "", "append all output lines")
	// ‹›«»
	sedCmd.Flags().StringVarP(&sedFlg.expression, "expression", "e", "", "expression all output lines")
}
