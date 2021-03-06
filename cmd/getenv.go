/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	i18nGetenvCmdTitle  = "Gets the environment variable"
	i18nGetenvCmdDetail = `
Gets the environment variable

It retrieves the value of the environment variable named by the key. 
It returns the value, which will be empty if the variable is not present.

It would get environment variable from 'ush' config file as well.
`
)

// getenvCmd represents the getenv command
var getenvCmd = &cobra.Command{
	Use:   "getenv",
	Short: i18nGetenvCmdTitle,
	Long:  i18nGetenvCmdDetail,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(getenvCmd)
}
