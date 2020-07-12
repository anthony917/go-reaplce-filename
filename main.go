package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string

	searchPath := "C:\\ExamplePath"
	keywordSearch := "keyword"
	keywordToReplace := "willReplace"
	keywordReplaceTo := "ReplaceTo"
	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.Contains(file, keywordSearch){
			newName := strings.ReplaceAll(file, keywordToReplace, keywordReplaceTo)
			err := os.Rename(file, newName)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
