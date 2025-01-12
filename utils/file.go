package utils

import (
	"errors"
	"strings"
	"os"
	"bufio"
	"path/filepath"
	"github.com/otiai10/copy"
)
//chatgpt
func CreateDirectoryIfNotExists(directoryPath string) error {
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
//chatgpt
func DirectoryExists(directoryPath string) bool {
	_, err := os.Stat(directoryPath)
	return !os.IsNotExist(err) 
}
func ReplaceFolder(srcPath, destPath string) error {
	err := os.RemoveAll(destPath)
	if err != nil {
		return err
	}
	err = copy.Copy(srcPath, destPath)
	if err != nil {
		return err
	}
	return nil
}
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	return err
}

func ReplaceFile(srcPath, destPath string) error {
	// 移动或替换文件
	err := os.Rename(srcPath, destPath)
	if err != nil {
		return err
	}
	return nil
}
func CopyFile(srcPath, destPath string) error {
	err := copy.Copy(srcPath, destPath)
	if err != nil {
		return err
	}
	return nil
}
func DeleteDirectory(directoryPath string) error {
	err := os.RemoveAll(directoryPath)
	return err
}
func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
func ReplaceStringInFile(filePath, findStr, replaceStr string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if !strings.Contains(string(content), findStr) {
		return errors.New("Not Found string!!"+findStr+" findpath: "+filePath)
	}
	newContent := strings.ReplaceAll(string(content), findStr, replaceStr)

	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
func FindIMGFiles(directory string) ([]string, error) {
	var imgFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".img") {
			imgFiles = append(imgFiles, strings.TrimSuffix(info.Name(), ".img"))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return imgFiles, nil
}

func Finddiff(a, b []string) []string {
	bMap := make(map[string]bool)
	for _, v := range b {
		bMap[v] = true
	}
	var diff []string
	for _, v := range a {
		if _, exists := bMap[v]; !exists {
			diff = append(diff, v)
		}
	}

	return diff
}
