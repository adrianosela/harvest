package main

import "fmt"

var (
	// injected at build-time
	version string
)

func main() {
	fmt.Println(version)
}
