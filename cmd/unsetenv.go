/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/isurfer21/UniversalShell/lib"
)

const (
	i18nUnsetenvCmdTitle  = "Unset environment variables"
	i18nUnsetenvCmdDetail = `
Unset environment variables

It is used to undo the effect of setenv or export commands.

It would unset environment variable into 'ush' config file.
`
	i18nUnsetenvTplConfirmationMsg = "%s %s \nAre you sure? (yes/no) "
	i18nUnsetenvTplPermDenied      = "Permission denied while %s %s\n"
)

type UnsetenvLib struct {
	evoke lib.Confirm
}

func (unsetenv *UnsetenvLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (unsetenv *UnsetenvLib) getWay(key string) string {
	if viper.IsSet(key) {
		return "Overwritting"
	} else {
		return "Unsetting"
	}
}

func (unsetenv *UnsetenvLib) isConfirmed(way string, key string) bool {
	fmt.Printf(i18nUnsetenvTplConfirmationMsg, way, key)
	if unsetenv.evoke.AskForConfirmation() {
		return true
	} else {
		return false
	}
}

type UnsetenvFlag struct {
	interactive bool
}

var (
	unsetenvFlg UnsetenvFlag
	unsetenvLib UnsetenvLib
)

// unsetenvCmd represents the unsetenv command
var unsetenvCmd = &cobra.Command{
	Use:   "unsetenv",
	Short: i18nUnsetenvCmdTitle,
	Long:  i18nUnsetenvCmdDetail,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			key := args[0]
			if unsetenvFlg.interactive {
				way := unsetenvLib.getWay(key)
				if unsetenvLib.isConfirmed(way, key) {
					viper.Set(key, "")
					viper.WriteConfig()
				} else {
					fmt.Printf(i18nUnsetenvTplPermDenied, strings.ToLower(way), key)
				}
			} else {
				viper.Set(key, "")
				viper.WriteConfig()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(unsetenvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// unsetenvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	unsetenvCmd.Flags().BoolVarP(&unsetenvFlg.interactive, "interactive", "i", false, "prompt before setting environment variable")
}
