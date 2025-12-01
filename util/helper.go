package util

import (
	"fmt"
	"os"
	"strings"
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

func FileSearch(sourcePath string) {
	dir := NavigateToSource(sourcePath)
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
			FileSearch(dirName)

		}
	}

}
