/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Folder struct {
	dossier Dossier
}

func (f *Folder) IsDirEmpty(name string) (bool, error) {
	file, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.Readdirnames(1) // Or file.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func (f *Folder) ReadDir(dirname string) ([]os.FileInfo, error) {
	file, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := file.Readdir(-1)
	file.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
@author Roland Singer [roland.singer@desertbit.com]

CopyDir recursively copies a directory tree, attempting to preserve permissions.
Source directory must exist, destination directory must *not* exist.
Symlinks are ignored and skipped.
*/
func (f *Folder) CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("%s already exists", dst)
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = f.CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = f.dossier.CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}
