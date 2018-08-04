/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"../lib"
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
	dossier lib.Dossier
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
	inplace    bool
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
		r, _ := regexp.Compile("(s/.*/.*/[gmisU]*)")
		if r.MatchString(args[0]) {
			var substitute string
			var filename string
			if sedFlg.expression != "" {
				substitute = sedFlg.expression
				if sedFlg.input != "" {
					filename = sedFlg.input
				} else {
					if len(args) > 0 {
						filename = args[0]
					} else if sedFlg.output != "" {
						filename = sedFlg.output
					} else if sedFlg.append != "" {
						filename = sedFlg.append
					}
				}
			} else {
				if len(args) > 0 {
					substitute = args[0]
					if sedFlg.input != "" {
						filename = sedFlg.input
					} else {
						if len(args) > 1 {
							filename = args[1]
						} else if sedFlg.output != "" {
							filename = sedFlg.output
						} else if sedFlg.append != "" {
							filename = sedFlg.append
						}
					}
				}
			}
			iContent := sedLib.readFile(filename)
			s := strings.Split(substitute, "/")
			pattern, replacement, rexflag := s[1], s[2], s[3]
			rex := regexp.MustCompile(pattern)
			oContent := ""
			if strings.Contains(rexflag, "g") {
				oContent = rex.ReplaceAllString(iContent, replacement)
			} else {
				if strings.Contains(rexflag, "i") {
					iContent = strings.ToLower(iContent)
				}
				motif := rex.FindAllString(iContent, 1)
				if len(motif) > 0 {
					oContent = strings.Replace(iContent, motif[0], replacement, 1)
				}
			}
			if sedFlg.output != "" {
				sedLib.handleError(sedLib.dossier.CopyFile(filename, sedFlg.output))
				sedLib.writeFile(sedFlg.output, oContent)
			} else if sedFlg.append != "" {
				sedLib.handleError(sedLib.dossier.CopyFile(filename, sedFlg.append))
				sedLib.writeFile(sedFlg.append, oContent)
			} else if sedFlg.inplace {
				sedLib.writeFile(filename, oContent)
			} else {
				fmt.Println(oContent)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// sedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	sedCmd.Flags().StringVarP(&sedFlg.input, "input", "I", "", "input all output lines")
	sedCmd.Flags().StringVarP(&sedFlg.output, "output", "O", "", "output all output lines")
	sedCmd.Flags().StringVarP(&sedFlg.append, "append", "A", "", "append all output lines")
	// ‹›«»
	sedCmd.Flags().StringVarP(&sedFlg.expression, "expression", "e", "", "expression all output lines")
	sedCmd.Flags().BoolVarP(&sedFlg.inplace, "in-place", "i", false, "files are to be edited in-place")
}
