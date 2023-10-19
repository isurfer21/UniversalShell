/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/isurfer21/UniversalShell/lib"
)

const (
	i18nLnCmdTitle  = "Create a symbolic link to a file"
	i18nLnCmdDetail = `
Create a symbolic link to a file

Make links between files, by default, it makes hard links; with the -s option,
it makes symbolic (or "soft") links.

Syntax:
  ln [OPTION]... [-T] OriginalSourceFile NewLinkFile (1st form)
  ln [OPTION]... OriginalSourceFile                  (2nd form)
  ln [OPTION]... OriginalSourceFile... DIRECTORY     (3rd form)
  ln [OPTION]... -t DIRECTORY OriginalSourceFile...  (4th form)

In the 1st form, create a link to OriginalSourceFile with the name NewLinkFile. 
In the 2nd form, create a link to OriginalSourceFile in the current directory.
In the 3rd and 4th forms, create links to each OriginalSourceFile in DIRECTORY.

Create hard links by default, symbolic links with --symbolic. When creating 
hard links, each OriginalSourceFile must exist. Symbolic links can hold 
arbitrary text; if later resolved, a relative link is interpreted in relation
to its parent directory.
`
	i18nLnTplConfirmationMsg = "Delete %s \nAre you sure? (yes/no) "
)

type LnLib struct {
	evoke lib.Confirm
}

func (ln *LnLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (ln *LnLib) getBackup(linkFile string) {
	if _, err := os.Stat(linkFile); err == nil {
		oldLinkFile := linkFile
		bakLinkFile := linkFile
		if lnFlg.suffix != "" {
			bakLinkFile += lnFlg.suffix
		} else {
			bakLinkFile += "~"
		}
		err := os.Rename(oldLinkFile, bakLinkFile)
		lnLib.handleError(err)
	}
}

func (ln *LnLib) delExistingLink(linkFile string) {
	if _, err := os.Stat(linkFile); err == nil {
		if lnFlg.interactive {
			fmt.Printf(i18nLnTplConfirmationMsg, linkFile)
			if lnLib.evoke.AskForConfirmation() {
				lnLib.handleError(os.Remove(linkFile))
			}
		} else {
			lnLib.handleError(os.Remove(linkFile))
		}
	}
}

type LnFlag struct {
	backup            bool
	directory         bool
	force             bool
	interactive       bool
	logical           bool
	nodereference     bool
	physical          bool
	symbolic          bool
	suffix            string
	targetdirectory   string
	notargetdirectory bool
	verbose           bool
}

var (
	lnFlg LnFlag
	lnLib LnLib
)

// lnCmd represents the ln command
var lnCmd = &cobra.Command{
	Use:   "ln",
	Short: i18nLnCmdTitle,
	Long:  i18nLnCmdDetail,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		origSrcFile := args[0]
		newLinkFile := filepath.Base(origSrcFile)
		if len(args) > 1 {
			newLinkFile = args[1]
		}
		if lnFlg.targetdirectory != "" {
			newLinkFile = path.Join(lnFlg.targetdirectory, newLinkFile)
		} else if lnFlg.notargetdirectory {
			trgFilepath, err := os.Getwd()
			lnLib.handleError(err)
			newLinkFile = path.Join(trgFilepath, newLinkFile)
		}
		if lnFlg.backup {
			lnLib.getBackup(newLinkFile)
		}
		if lnFlg.force {
			lnLib.delExistingLink(newLinkFile)
		}
		if lnFlg.physical {
			lnLib.handleError(os.Link(origSrcFile, newLinkFile))
		}
		if lnFlg.symbolic {
			lnLib.handleError(os.Symlink(origSrcFile, newLinkFile))
		}
	},
}

func init() {
	rootCmd.AddCommand(lnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// lnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	lnCmd.Flags().BoolVarP(&lnFlg.backup, "backup", "b", false, "make a backup of each existing destination file")
	lnCmd.Flags().BoolVarP(&lnFlg.directory, "directory", "d", false, "allow the superuser to attempt to hard link directories")
	lnCmd.Flags().BoolVarP(&lnFlg.force, "force", "f", false, "remove existing destination files")
	lnCmd.Flags().BoolVarP(&lnFlg.interactive, "interactive", "i", false, "prompt whether to remove destinations")
	lnCmd.Flags().BoolVarP(&lnFlg.logical, "logical", "L", false, "make hard links to symbolic link references")
	// lnCmd.Flags().BoolVarP(&lnFlg.nodereference, "no-dereference", "n", false, "treat destination that is a symlink to a directory as if it were a normal file")
	lnCmd.Flags().BoolVarP(&lnFlg.physical, "physical", "P", false, "make hard links directly to symbolic links")
	lnCmd.Flags().BoolVarP(&lnFlg.symbolic, "symbolic", "s", false, "make symbolic links instead of hard links")
	lnCmd.Flags().StringVarP(&lnFlg.suffix, "suffix", "S", "", "override the usual backup suffix")
	lnCmd.Flags().StringVarP(&lnFlg.targetdirectory, "target-directory", "t", "", "specify the DIRECTORY in which to create the links")
	lnCmd.Flags().BoolVarP(&lnFlg.notargetdirectory, "no-target-directory", "T", false, "treat NewLinkFile as a normal file")
	// lnCmd.Flags().BoolVarP(&lnFlg.verbose, "verbose", "v", false, "print name of each linked file")
}
