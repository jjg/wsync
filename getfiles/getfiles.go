package getfiles

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func Flat(sourcePath *string) []string {

	sourceFilePaths := []string{}
	d, _ := os.ReadDir(*sourcePath)

	for _, f := range d {
		if !f.IsDir() {
			filename := fmt.Sprintf("%v/%v", *sourcePath, f.Name())
			sourceFilePaths = append(sourceFilePaths, filename)
		}
	}
	return sourceFilePaths
}

func Recursive(sourcePath *string) []string {

	sourceFilePaths := []string{}
	fileSystem := os.DirFS(*sourcePath)

	fs.WalkDir(fileSystem, ".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			if !d.IsDir() {
				sourceFilePaths = append(sourceFilePaths, path)
			}
			return nil
		})
	return sourceFilePaths
}
