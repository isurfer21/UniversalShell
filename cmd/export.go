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
	i18nExportCmdTitle  = "Exports variables to the environment"
	i18nExportCmdDetail = `
Exports variables to the environment

It is used to mark a shell variable for export to child processes.

It would set/unset environment variable into 'ush' config file.
`
	i18nExportTplConfirmationMsg = "%s %s \nAre you sure? (yes/no) "
	i18nExportTplInvalidKVPair   = "%s is an invalid key-value pair\n"
	i18nExportTplPermDenied      = "Permission denied while %s %s\n"
)

type ExportLib struct {
	evoke lib.Confirm
}

func (export *ExportLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (export *ExportLib) isKVPair(ev string) bool {
	if strings.Index(ev, "=") >= 0 {
		return true
	} else {
		return false
	}
}

func (export *ExportLib) getWay(key string) string {
	if viper.IsSet(key) {
		return "Overwritting"
	} else {
		return "Setting"
	}
}

func (export *ExportLib) isConfirmed(way string, key string) bool {
	fmt.Printf(i18nExportTplConfirmationMsg, way, key)
	if export.evoke.AskForConfirmation() {
		return true
	} else {
		return false
	}
}

func (export *ExportLib) setKVPair(ev string) error {
	if export.isKVPair(ev) {
		pair := strings.Split(ev, "=")
		key := pair[0]
		val := pair[1]
		if val != "" {
			viper.Set(key, val)
		} else {
			viper.Set(key, "")
		}
		viper.WriteConfig()
		return nil
	}
	return fmt.Errorf(i18nExportTplInvalidKVPair, ev)
}

func (export *ExportLib) getKey(ev string) (string, error) {
	if export.isKVPair(ev) {
		pair := strings.Split(ev, "=")
		return pair[0], nil
	}
	return "", fmt.Errorf(i18nExportTplInvalidKVPair, ev)
}

type ExportFlag struct {
	interactive bool
}

var (
	exportFlg ExportFlag
	exportLib ExportLib
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: i18nExportCmdTitle,
	Long:  i18nExportCmdDetail,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			pair := args[0]
			if exportFlg.interactive {
				key, err := exportLib.getKey(pair)
				exportLib.handleError(err)
				way := exportLib.getWay(key)
				if exportLib.isConfirmed(way, key) {
					exportLib.handleError(exportLib.setKVPair(pair))
				} else {
					fmt.Printf(i18nExportTplPermDenied, strings.ToLower(way), key)
				}
			} else {
				exportLib.handleError(exportLib.setKVPair(pair))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	exportCmd.Flags().BoolVarP(&exportFlg.interactive, "interactive", "i", false, "prompt before setting environment variable")
}
