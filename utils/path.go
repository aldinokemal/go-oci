package utils

import (
	"os"
)

func CreateTmpFolder() (string, error) {
	return os.MkdirTemp("", "registry")
}

func RemoveFile(path string) error {
	return os.Remove(path)
}
