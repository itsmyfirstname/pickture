package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func move(fileObjectSource string, fileObjectDestination string) (string, error) {
	err := os.Rename(fileObjectSource, fileObjectDestination)

	if err != nil {
		fmt.Errorf("Failed to move | %w", err)
		return "", err
	}
	return "success", nil
}

func copy(fileObjectSource string, fileObjectDestination string) (string, error) {
	// create temp dir
	// add source file to temp object
	file, err := os.Open(fileObjectSource)

	if err != nil {
		return "", fmt.Errorf("Failed to open file | %w", err)
	}
	defer file.Close()

	newFile, err := os.Create(fileObjectDestination)
	if err != nil {
		return "", fmt.Errorf("Failed to create dest file | %w", err)
	}
	defer newFile.Close()

	_, err = io.Copy(file, newFile)
	if err != nil {
		return "", fmt.Errorf("Failed to copy file | %w", err)
	}

	stat, err := file.Stat()

	if err != nil {
		return "", fmt.Errorf("Failed to read src file permissions | %w", err)
	}

	err = os.Chmod(fileObjectDestination, stat.Mode())

	if err != nil {
		return "", fmt.Errorf("Failed to apply src permissions to destination | %w", err)
	}
	return "Successfully copied", nil
}

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
