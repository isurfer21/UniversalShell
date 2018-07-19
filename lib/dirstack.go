/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package lib

import (
	"errors"
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type DirStack struct {
	Stack []string
}

func (ds *DirStack) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (ds *DirStack) Load() {
	str := viper.GetString("DIRSTACK")
	if len(str) > 0 {
		ds.Stack = strings.Split(str, ",")
	} else {
		ds.Stack = []string{}
	}
}

func (ds *DirStack) Short() []string {
	home, err := homedir.Dir()
	ds.handleError(err)
	Stack := []string{}
	for i := 0; i < len(ds.Stack); i += 1 {
		path := strings.Replace(ds.Stack[i], home, "~", -1)
		Stack = append(Stack, path)
	}
	return Stack
}

func (ds *DirStack) Reverse(list []string) []string {
	for i := len(list)/2 - 1; i >= 0; i-- {
		opp := len(list) - 1 - i
		list[i], list[opp] = list[opp], list[i]
	}
	return list
}

func (ds *DirStack) Clear() error {
	if len(ds.Stack) > 0 {
		ds.Stack = []string{}
		return nil
	}
	return errors.New("Directory Stack is empty!")
}

func (ds *DirStack) Push(path string) {
	ds.Stack = append(ds.Stack, path)
}

func (ds *DirStack) Pop() (string, error) {
	if len(ds.Stack) > 0 {
		lastItem := ds.Stack[len(ds.Stack)-1]
		ds.Stack = ds.Stack[:len(ds.Stack)-1]
		return lastItem, nil
	}
	return "", errors.New("Directory stack is empty!")
}

func (ds *DirStack) Save() {
	input := strings.Join(ds.Stack, ",")
	viper.Set("DIRSTACK", input)
	viper.WriteConfig()
}
