package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func specifyUnit(size int64) string {
	var newSize float64 = float64(size)
	units := []string{"Byte", "KB", "MB", "GB"}
	index := 0
	divider := 1.0
	for {
		size = size / 1024
		if size > 0 {
			index++
			divider *= 1024
		} else {
			break
		}
	}
	sizeReturn := fmt.Sprintf("%.1f", newSize/divider)
	return sizeReturn + " " + units[index]
}
func scanDir(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDir(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(file.Mode(), "\t", specifyUnit(file.Size()), "\t", filePath)
		}

	}
	return nil
}

func main() {

	scanDir("d:\\freegate")
	fmt.Println(specifyUnit(12))

}
