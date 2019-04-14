package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Printf("Hello, I am container %s running on %s/%s\n", hostname, runtime.GOOS, runtime.GOARCH)
}
