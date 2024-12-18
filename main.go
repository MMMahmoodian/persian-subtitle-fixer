package main

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"os"
	"path/filepath"
)

func main() {
	directory := "example" // The current directory

	subs, err := getFilesByExtension(directory, "srt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sub := range subs {
		content := readfile(sub)
		converted, err := charmap.Windows1256.NewDecoder().Bytes(content)
		if err != nil {
			return
		}
		err = writefile(sub, converted)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func getFilesByExtension(dir string, ext string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && path[len(dir):] != "" && filepath.Ext(path) == "."+ext {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func readfile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return data
}

func writefile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
