// Example of how to embed versioning into the project during compile/build time
package main

import (
	"fmt"
)

//go:generate go-winres make --product-version=git-tag

// variables that are are going to be injected with version/build information
var (
	Version string
	Build   string
)

// go build -ldflags "-w -s -X main.Version=1.2.3.4 -X main.Build=17/11/2024" .
func main() {

	fmt.Println("Version: ", Version)
	fmt.Println("Build Time: ", Build)

}
