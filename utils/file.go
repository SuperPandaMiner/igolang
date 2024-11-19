package utils

import (
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

func Mkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveFile(file io.Reader, dir string, filename ...string) error {
	err := Mkdir(dir)
	if err != nil {
		return err
	}

	var dst string
	if len(filename) > 0 {
		dst = filepath.Join(dir, filename[0])
	} else {
		dst = filepath.Join(dir, uuid.New().String())
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, file); err != nil {
		return err
	}
	return nil
}
