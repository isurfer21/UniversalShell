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

	"../lib"
)

const (
	i18nSetenvCmdTitle  = "Sets the environment variable"
	i18nSetenvCmdDetail = `
Sets the environment variable

It sets the value of the environment variable named by the key. It returns an 
error, if any.

It would set environment variable into 'ush' config file.
`
	i18nSetenvTplConfirmationMsg = "%s %s \nAre you sure? (yes/no) "
	i18nSetenvTplInvalidKVPair   = "%s is an invalid key-value pair\n"
	i18nSetenvTplPermDenied      = "Permission denied while %s %s\n"
)

type SetenvLib struct {
	evoke lib.Confirm
}

func (setenv *SetenvLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (setenv *SetenvLib) getWay(key string) string {
	_, ok := os.LookupEnv(key)
	if !ok {
		return "Setting"
	} else {
		return "Overwritting"
	}
}

func (setenv *SetenvLib) isConfirmed(way string, key string) bool {
	fmt.Printf(i18nSetenvTplConfirmationMsg, way, key)
	if setenv.evoke.AskForConfirmation() {
		return true
	} else {
		return false
	}
}

type SetenvFlag struct {
	interactive bool
}

var (
	setenvFlg SetenvFlag
	setenvLib SetenvLib
)

// setenvCmd represents the setenv command
var setenvCmd = &cobra.Command{
	Use:   "setenv",
	Short: i18nSetenvCmdTitle,
	Long:  i18nSetenvCmdDetail,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			key, val := args[0], args[1]
			if setenvFlg.interactive {
				way := setenvLib.getWay(key)
				if setenvLib.isConfirmed(way, key) {
					viper.Set(key, val)
					viper.WriteConfig()
				} else {
					fmt.Printf(i18nSetenvTplPermDenied, strings.ToLower(way), key)
				}
			} else {
				viper.Set(key, val)
				viper.WriteConfig()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setenvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// setenvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	setenvCmd.Flags().BoolVarP(&setenvFlg.interactive, "interactive", "i", false, "prompt before setting environment variable")
}
