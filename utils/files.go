package utils

import (
	"os"
	"os/user"
)

func relativeToFullDir(dir string) string {
	user, _ := user.Current()
	return user.HomeDir + dir
}

func CreateDir(relativeDir string) {
	fullDir := relativeToFullDir(relativeDir)
	err := os.MkdirAll(fullDir, os.ModePerm)
	Check(err)
}

func WriteFile(relativePath string, content []byte) {
	fullPath := relativeToFullDir(relativePath)
	err := os.WriteFile(fullPath, content, os.ModePerm)
	Check(err)
}
