package common

import (
	"io/ioutil"
	"os"
)

// GetPath lookups the full path for the current working directory
func GetPath() (string, error) {

	currPath, err := os.Getwd()

	return currPath, err

}

// GetFileList gets a list of files within the target directory
func GetFileList(dir string) ([]os.FileInfo, error) {
	fileList, err := ioutil.ReadDir(dir)
	return fileList, err
}
