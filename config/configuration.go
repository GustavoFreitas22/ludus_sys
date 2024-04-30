package config

import (
	"fmt"
	"runtime"
)

func DetectOsRuntime() string {
	if runtime.GOOS == "windows" {
		fmt.Println("say hey")
		return "win"
	} else {
		return runtime.GOOS
	}
}
