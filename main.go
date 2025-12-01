package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walk(sourcePath string) {
	FileTypes := [9]string{"jpg", "jpeg", "png", "hvec", "raw", "dcs", "mp4", "avi", "mkv"}
	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") {
				return nil
			}
			return nil
		}

		for _, file := range FileTypes {
			if strings.HasSuffix(info.Name(), file) {
				fmt.Fprintln(os.Stdout, "Found filename", info.Name())
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
	}
}

func main() {
	sourceDirectory := flag.String("source", "", "Where are the pics at")
	targetDirectory := flag.String("target", "", "Where are the pics at")

	flag.Parse()

	if *sourceDirectory == "" || *targetDirectory == "" {
		fmt.Println("Missing required flags | --source, --target")
		fmt.Fprintln(os.Stderr, "Source Directory |", *sourceDirectory)
		fmt.Fprintln(os.Stderr, "Target Directory |", *targetDirectory)
		os.Exit(1)
	}
	walk(*sourceDirectory)

}
