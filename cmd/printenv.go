/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	i18nPrintenvCmdTitle  = "Prints all or part of the environment"
	i18nPrintenvCmdDetail = `
Prints all or part of the environment

Print all or part of environment variable. If no environment VARIABLE 
specified, print them all.
`
)

// printenvCmd represents the printenv command
var printenvCmd = &cobra.Command{
	Use:   "printenv",
	Short: i18nPrintenvCmdTitle,
	Long:  i18nPrintenvCmdDetail,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			key := args[0]
			val, ok := os.LookupEnv(key)
			if !ok {
				fmt.Printf("%s not set\n", key)
			} else {
				fmt.Printf("%s\n", val)
			}
		} else {
			for _, pair := range os.Environ() {
				fmt.Println(pair)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printenvCmd)
}
