package files

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"
)

type File struct {
	Name    string      `json:"name"`
	Size    int64       `json:"size"`
	Mode    os.FileMode `json:"mode"`
	ModTime time.Time   `json:"mod_time"`
	IsDir   bool        `json:"is_dir"`
}

//Returns all files and dirs on current level
func GetFilesInDirectory(path string) ([]File, error) {
	result := []File{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return result, err
	}

	for _, f := range files {
		file := File{}
		file.Name = f.Name()
		file.Size = f.Size()
		file.Mode = f.Mode()
		file.ModTime = f.ModTime()
		file.IsDir = f.IsDir()
		result = append(result, file)
	}

	return result, nil
}

//Returns file and lines count
func ViewFile(path string) (string, int, error) {
	var result string
	var lineCount int

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return result, lineCount, err
	}

	file, _ := os.Open(path)
	fileScanner := bufio.NewScanner(file)
	lineCount = 0
	for fileScanner.Scan() {
		lineCount++
	}

	return string(f), lineCount, nil
}

//Deletes a file
func DeleteFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() == true {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	} else {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

//Write new content to file
func WriteFile(path string, content string) error {
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
