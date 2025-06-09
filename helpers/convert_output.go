package helpers

import (
	"os"
	"strings"
)

func ConvertLibrariesToOutputFile(listOfLibraries string) error {
	libs := strings.Split(listOfLibraries, ",")
	for i, lib := range libs {
		libs[i] = "- " + strings.TrimSpace(lib)
	}

	content := strings.Join(libs, "\n") + "\n"

	return os.WriteFile("output.txt", []byte(content), 0o200)
}
