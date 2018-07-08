/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/labstack/gommon/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	i18nRootAppName    = "Ush - Universal Shell"
	i18nRootAppVersion = "0.0.2"
	i18nRootAppLicense = "MIT License"

	i18nRootAppCopyrightYear = 2018
	i18nRootAppObjective     = `
Shell commands that runs similar everywhere
`
	i18nRootAppLicenseDetail = `
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
	i18nRootTplCommandMissing = `
Error: Command is missing!

Tip: Try any of these, 
     $ ush -h
     $ ush --help
`
	i18nRootTplVersionCopyright = `version %s
Copyright (c) %s Abhishek Kumar
`
	i18nRootTplLicenseStatement = `It is licensed under the '%s'
%s
`
)

type RootLib struct {
}

func (root *RootLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type rootFlag struct {
	version bool
	license bool
}

var (
	dye        color.Color
	configFile string
	rootFlg    rootFlag
	rootLib    RootLib
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ush",
	Short: i18nRootAppName,
	Long:  proviso() + dye.Bold(i18nRootAppName) + i18nRootAppObjective,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(dye.Bold(i18nRootAppName))
		if rootFlg.version || rootFlg.license {
			currentYear := time.Now().UTC().Year()
			copyrightDuration := strconv.Itoa(i18nRootAppCopyrightYear)
			if currentYear > i18nRootAppCopyrightYear {
				copyrightDuration += "-" + strconv.Itoa(currentYear)
			}
			fmt.Printf(i18nRootTplVersionCopyright, i18nRootAppVersion, copyrightDuration)
			if rootFlg.license {
				fmt.Printf(i18nRootTplLicenseStatement, i18nRootAppLicense, i18nRootAppLicenseDetail)
			}
		} else {
			fmt.Printf(i18nRootTplCommandMissing)
		}
	},
}

func proviso() (str string) {
	if runtime.GOOS == "windows" {
		dye.Disable()
	}
	return ""
}

// Execute adds all child commands to the root command and sets flags appropriately. This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings. Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.ush.toml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().BoolVarP(&rootFlg.version, "version", "v", false, "show version number and exit")
	rootCmd.Flags().BoolVarP(&rootFlg.license, "license", "l", false, "show applied license and exit")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		rootLib.handleError(err)

		// Search config in home directory with name ".ush" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ush")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	rootLib.handleError(viper.ReadInConfig())
}
