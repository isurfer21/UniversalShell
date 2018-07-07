/*
Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
*/

package lib

import (
	"io"
	"os"
)

type Dossier struct {
}

func (d *Dossier) IsWritable(path string) (bool, error) {
	w := false
	f, err := os.Lstat(path)
	if err != nil {
		return false, err
	}
	if f.Mode().Perm()&(1<<(uint(7))) == 0 {
		w = true
	}
	return w, err
}

/*
@author Roland Singer [roland.singer@desertbit.com]

CopyFile copies the contents of the file named src to the file named
by dst. The file will be created if it does not already exist. If the
destination file exists, all it's contents will be replaced by the contents
of the source file. The file mode will be copied from the source and
the copied data is synced/flushed to stable storage.
*/
func (d *Dossier) CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}
