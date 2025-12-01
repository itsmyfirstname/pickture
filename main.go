package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func navigateToSource(sourcePath string) string {
	error := os.Chdir(sourcePath)

	if error != nil {
		fmt.Fprintln(os.Stderr, "Unable to navigate to source path |", sourcePath)
		os.Exit(1)
	}

	dir, error := os.Getwd()

	if error != nil {
		fmt.Println("Error getting current working directory")
		fmt.Fprintln(os.Stderr, "Source directory |", sourcePath)
		os.Exit(1)
	}

	return dir
}

func walk(sourcePath string) {
	FileTypes := [9]string{"jpg", "jpeg", "png", "hvec", "raw", "dcs", "mp4", "avi", "mkv"}
	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		//fmt.Printf("Visited: %s (IsDir: %t)\n", path, info.IsDir())

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

func fileSearch(sourcePath string) {
	dir := navigateToSource(sourcePath)
	files, error := os.ReadDir(dir)

	if error != nil {
		fmt.Println("Unable to read directory")
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			dirInfo, error := file.Info()
			if error != nil {
				fmt.Println("Unexpected error getting dir info")
			}
			dirName := dirInfo.Name()

			if strings.HasPrefix(dirName, ".") {
				continue
			}
			fmt.Println(dirName)
			fileSearch(dirName)

		}
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
