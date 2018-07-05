/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/rhysd/abspath"
	"github.com/spf13/cobra"
)

const (
	i18nHomedirCmdTitle  = "Path of home directory"
	i18nHomedirCmdDetail = `
Returns the path of home directory

If home directory cannot be obtained or is not an absolute path, it will return an error.
`
)

// homedirCmd represents the homedir command
var homedirCmd = &cobra.Command{
	Use:   "homedir",
	Short: i18nHomedirCmdTitle,
	Long:  i18nHomedirCmdDetail,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := abspath.HomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(home.String())
	},
}

func init() {
	rootCmd.AddCommand(homedirCmd)
}
