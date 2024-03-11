package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, _ := debug.ReadBuildInfo()
	fmt.Printf("Version: %s", info.Main.Version)
}
