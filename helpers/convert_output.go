package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func ConvertLibrariesToOutputFile(listOfLibraries string, outputFile string) error {
	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		return err
	}

	libs := strings.Split(listOfLibraries, ",")
	for i, lib := range libs {
		libs[i] = "- " + strings.TrimSpace(lib)
	}

	content := strings.Join(libs, "\n") + "\n"

	return os.WriteFile(outputFile, []byte(content), 0o644)
}
