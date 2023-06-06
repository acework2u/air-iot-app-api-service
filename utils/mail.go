package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	fmt.Println("Am parsing template...")
	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
