package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FileOrDirExist(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsExist(err) {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

func CreateDirNotExist(dir string) error {
	exist, _ := FileOrDirExist(dir)
	if exist {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}

func ListFileInDir(dir string) ([]string, error) {
	if exist, e := FileOrDirExist(dir); !exist {
		return nil, e
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fileList []string
	// 获取目录内文件名
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList, nil
}

func WriteToFile(filePath string, content []byte) error {
	err := ioutil.WriteFile(filePath, content, 0644)
	return err
}

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		er := file.Close()
		if er != nil {
			fmt.Println("close file error: " + er.Error())
		}
	}()
	content, err := ioutil.ReadAll(file)
	return content, err
}
