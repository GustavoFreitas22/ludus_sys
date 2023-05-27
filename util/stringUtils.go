package util

import (
	"os"
	"strings"
)

func getPathFromString(values []string) string {
	pathSeparator := string(os.PathSeparator)
	return strings.Join(values, pathSeparator)
}
