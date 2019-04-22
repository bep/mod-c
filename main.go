package main

import (
	"fmt"
)

var (
	name    = "c"
	version = "v0.54"
)

func main() {
	fmt.Println(Version())
}

func Version() string {
	return name + ": " + version
}
