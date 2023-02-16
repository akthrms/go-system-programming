package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var imageSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tiff": true,
	".eps":  true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`Find images

Usage:
	%s [path to find]
`, os.Args[0])
		return
	}

	root := os.Args[1]

	err := filepath.Walk(root, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			if info.Name() == "_build" {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		ext := strings.ToLower(filepath.Ext(info.Name()))
		if imageSuffix[ext] {
			if rel, err := filepath.Rel(root, path); err != nil {
				return nil
			} else {
				fmt.Printf("%s\n", rel)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(1, err)
	}
}
