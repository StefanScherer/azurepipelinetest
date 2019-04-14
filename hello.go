package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Printf("Hello, I'm %s running on %s/%s\n", hostname, runtime.GOOS, runtime.GOARCH)
}
