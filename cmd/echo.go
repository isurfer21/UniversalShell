/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	i18nEchoCmdTitle  = "Write arguments to the standard output"
	i18nEchoCmdDetail = `
Write arguments to the standard output

The echo utility writes any specified operands, separated by single blank 
characters and followed by a newline character, to the standard output.
`
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: i18nEchoCmdTitle,
	Long:  i18nEchoCmdDetail,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
