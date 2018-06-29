package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/mkideal/cli"
)

const (
	appName    = "Ush - Universal Shell"
	appVersion = "0.0.1"
	appLicense = "MIT License"

	appLicenseDetail = "A short and simple permissive license with conditions only requiring \npreservation of copyright and license notices. Licensed works, \nmodifications, and larger works may be distributed under different \nterms and without source code."
	appCopyrightYear = 2018
)

var (
	clr = color.Color{}
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(pwd),
		cli.Tree(ls),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	print("\n")
}

var help = cli.HelpCommand("Provides Help information for built-in commands")

type rootT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"show version number and quit"`
	License bool `cli:"l,license" usage:"show license and quit"`
}

var root = &cli.Command{
	Desc: clr.Bold(appName),
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("%s \n", clr.Bold(appName))
		if argv.Version || argv.License {
			if argv.Version {
				currentYear := time.Now().UTC().Year()
				copyrightDuration := strconv.Itoa(appCopyrightYear)
				if currentYear > appCopyrightYear {
					copyrightDuration += "-" + strconv.Itoa(currentYear)
				}
				ctx.String("version %s\nCopyright (c) %s Abhishek Kumar\n", appVersion, copyrightDuration)
			}
			if argv.License {
				ctx.String("It is licensed under the '%s'\n\n%s\n", appLicense, appLicenseDetail)
			}
		} else {
			ctx.String("\nError: Command is missing! \n\nTip: Try any of these, \n     $ ush -h\n     $ ush --help\n     $ ush help\n")
		}
		return nil
	},
}

type pwdT struct {
	cli.Helper
}

var pwd = &cli.Command{
	Name: "pwd",
	Desc: "Displays the path of present working directory",
	Argv: func() interface{} { return new(pwdT) },
	Fn: func(ctx *cli.Context) error {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		ctx.String("%s\n", dir)
		return nil
	},
}

type lsT struct {
	cli.Helper
	Tab  bool `cli:"t,tab" usage:"show tab separated list"`
	Long bool `cli:"l,long" usage:"display extended file metadata as a table"`
}

var ls = &cli.Command{
	Name: "ls",
	Desc: "Displays a list of files and sub-directories in a directory",
	Argv: func() interface{} { return new(lsT) },
	Fn: func(ctx *cli.Context) error {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		} else {
			items, err := ReadDir(pwd)
			if err != nil {
				log.Fatal(err)
			} else {
				argv := ctx.Argv().(*lsT)
				separator := "\n"
				if argv.Tab {
					separator = " \t"
				}
				list := []string{}
				for _, file := range items {
					if argv.Long {
						row := []string{string(file.Mode()), string(file.Size()), file.Name() /*, string(file.IsDir()), string(file.ModTime())*/}
						list = append(list, strings.Join(row[:], "\t"))
					} else {
						list = append(list, file.Name())
					}
				}
				ctx.String("%s\n", strings.Join(list[:], separator))
			}
		}
		return nil
	},
}

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}
