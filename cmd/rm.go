/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files and directories",
	Long: ` 
Remove files and directories

It attempts to remove the non-directory type files specified on the 
command line. If the permissions of the file do not permit writing,
and the standard input device is a terminal, the user is prompted
(on the standard error output) for confirmation.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			dir := strings.TrimSuffix(filepath.Base(args[0]), filepath.Ext(args[0]))
			dir = filepath.Join(os.TempDir(), dir)
			dirs := filepath.Join(dir, `tmpdir`)
			err := os.MkdirAll(dirs, 0777)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			file := filepath.Join(dir, `tmpfile`)
			f, err := os.Create(file)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			f.Close()
			file = filepath.Join(dirs, `tmpfile`)
			f, err = os.Create(file)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			f.Close()

			err = removeContents(dir)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command	is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
