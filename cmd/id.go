/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
)

const (
	i18nIdCmdTitle  = "Print user and group id's"
	i18nIdCmdDetail = `
Print user and group id's

Print real and effective user id (uid) and group id (gid), prints identity 
information about the given user, or if no user is specified the current 
process.

Syntax:
  id [options]... [username]

By default, it prints the real user id, real group id, effective user id if 
different from the real user id, effective group id if different from the real
group id, and supplemental group ids.

Each of these numeric values is preceded by an identifying string and followed
by the corresponding user or group name in parentheses.
`
)

type IdLib struct {
}

func (id *IdLib) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (id *IdLib) userInfo(usr *user.User) string {
	var output string = ""
	if idFlg.name || idFlg.real {
		if idFlg.name {
			output = usr.Username
		}
		if idFlg.real {
			output = usr.Name
		}
	} else {
		output = usr.Uid
	}
	return output
}

func (id *IdLib) groupInfo(usr *user.User) string {
	output := ""
	if idFlg.name {
		grp, err := user.LookupGroupId(usr.Gid)
		idLib.handleError(err)
		output = grp.Name
	} else {
		output = usr.Gid
	}
	return output
}

func (id *IdLib) groupsInfo(usr *user.User) string {
	output := ""
	grps, err := usr.GroupIds()
	idLib.handleError(err)
	if idFlg.name {
		list := []string{}
		for i := 0; i < len(grps); i += 1 {
			grp, err := user.LookupGroupId(grps[i])
			idLib.handleError(err)
			list = append(list, grp.Name)
		}
		output = strings.Join(list[:], " ")
	} else {
		output = strings.Join(grps[:], " ")
	}
	return output
}

func (id *IdLib) allInfo(usr *user.User) string {
	grp, err := user.LookupGroupId(usr.Gid)
	idLib.handleError(err)
	grpName := grp.Name
	grps, err := usr.GroupIds()
	idLib.handleError(err)
	list := []string{}
	for i := 0; i < len(grps); i += 1 {
		grp, err := user.LookupGroupId(grps[i])
		idLib.handleError(err)
		grpPair := fmt.Sprintf("%s(%s)", grp.Gid, grp.Name)
		list = append(list, grpPair)
	}
	grpsPair := strings.Join(list[:], ",")
	output := fmt.Sprintf("uid=%s(%s) gid=%s(%s) groups=%s", usr.Uid, usr.Username, usr.Gid, grpName, grpsPair)
	return output
}

type IdFlag struct {
	group  bool
	groups bool
	name   bool
	real   bool
	user   bool
}

var (
	idFlg IdFlag
	idLib IdLib
)

// idCmd represents the id command
var idCmd = &cobra.Command{
	Use:   "id",
	Short: i18nIdCmdTitle,
	Long:  i18nIdCmdDetail,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var usr *user.User
		var err error
		if len(args) > 0 {
			usr, err = user.Lookup(args[0])
		} else {
			usr, err = user.Current()
		}
		idLib.handleError(err)
		if idFlg.user {
			fmt.Println(idLib.userInfo(usr))
		} else if idFlg.group {
			fmt.Println(idLib.groupInfo(usr))
		} else if idFlg.groups {
			fmt.Println(idLib.groupsInfo(usr))
		} else {
			fmt.Println(idLib.allInfo(usr))
		}
	},
}

func init() {
	rootCmd.AddCommand(idCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// idCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	idCmd.Flags().BoolVarP(&idFlg.group, "group", "g", false, "print only the group id")
	idCmd.Flags().BoolVarP(&idFlg.groups, "groups", "G", false, "print only the supplementary groups")
	idCmd.Flags().BoolVarP(&idFlg.name, "name", "n", false, "print the user or group name instead of the ID number")
	idCmd.Flags().BoolVarP(&idFlg.real, "real", "r", false, "print the real, instead of effective, user")
	idCmd.Flags().BoolVarP(&idFlg.user, "user", "u", false, "print only the user id")
}
