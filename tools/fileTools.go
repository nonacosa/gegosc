package tools

import (
	"io/ioutil"
	"os"
	"path/filepath"
)


type Files struct {
	isDir   bool
	name 	string
}


func FilePathWalkDir(root string) ([] Files, error) {
	var files [] Files
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//if !info.IsDir() {
		//	files = append(files, path)
		//}
		files = append(files, Files{
			isDir: info.IsDir(),
			name:  info.Name(),
		})
		return nil
	})
	return files, err
}

func IOReadDir(root string) ([] Files, error) {
	var files [] Files
	fileInfo, err := ioutil.ReadDir(root)
	_ = fileInfo
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, Files{
			isDir: file.IsDir(),
			name:  file.Name(),
		})

	}
	return files, nil
}