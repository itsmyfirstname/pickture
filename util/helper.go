package util

import (
	"fmt"
	"os"
)

func NavigateToSource(sourcePath string) string {
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
