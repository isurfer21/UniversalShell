/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	i18nUnameCmdTitle  = "Print operating system name"
	i18nUnameCmdDetail = `
Print operating system name

It writes symbols representing one or more system characteristics to the 
standard output.
`
)

type unameFlag struct {
	all       bool
	machine   bool
	node      bool
	processor bool
	release   bool
	system    bool
	version   bool
}

var unameFlg unameFlag

// unameCmd represents the uname command
var unameCmd = &cobra.Command{
	Use:   "uname",
	Short: i18nUnameCmdTitle,
	Long:  i18nUnameCmdDetail,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var macArch string
		if runtime.GOARCH == "amd64" {
			macArch = "x86_64"
		} else {
			macArch = "x86"
		}

		platform := runtime.GOOS

		hostname, err := os.Hostname()
		checkError(err)

		if unameFlg.all {
			fmt.Println(platform, hostname, macArch)
		}
		if unameFlg.machine {
			fmt.Println(macArch)
		}
		if unameFlg.node {
			fmt.Println(hostname)
		}
		if unameFlg.processor {
			fmt.Println("N/A")
		}
		if unameFlg.release {
			fmt.Println("N/A")
		}
		if unameFlg.version {
			fmt.Println("N/A")
		}

		if unameFlg.system || (!unameFlg.all && !unameFlg.machine && !unameFlg.node && !unameFlg.processor && !unameFlg.release && !unameFlg.system && !unameFlg.version) {
			fmt.Println(platform)
		}
	},
}

func init() {
	rootCmd.AddCommand(unameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// unameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	unameCmd.Flags().BoolVarP(&unameFlg.all, "all", "a", false, "Behave as though all of the options were specified.")
	unameCmd.Flags().BoolVarP(&unameFlg.machine, "machine", "m", false, "print the machine hardware name.")
	unameCmd.Flags().BoolVarP(&unameFlg.node, "node", "n", false, "print the nodename (the system is known by to a communications network).")
	unameCmd.Flags().BoolVarP(&unameFlg.processor, "processor", "p", false, "print the machine processor architecture name.")
	unameCmd.Flags().BoolVarP(&unameFlg.release, "release", "r", false, "print the operating system release.")
	unameCmd.Flags().BoolVarP(&unameFlg.system, "system", "s", false, "print the operating system name.")
	unameCmd.Flags().BoolVarP(&unameFlg.version, "version", "v", false, "print the operating system version.")
}
