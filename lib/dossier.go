package lib

import (
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
