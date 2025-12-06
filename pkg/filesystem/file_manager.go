package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Read(path string) ([]string, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileContent), "\n"), nil
}

func ResolvePath(basePath string, file string) (string, error) {
	absPath, err := filepath.Abs(fmt.Sprintf("%s/%s", basePath, file))

	if err != nil {
		return "", err
	}
	return absPath, nil
}
