package utils

import "os"

func IsDirExist(dir string) bool {
	fh, err := os.Open(dir)
	if err != nil {
		return false
	}
	defer fh.Close()
	return true
}
