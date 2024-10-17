package main

import (
	"fmt"
)

//go:generate go-winres make --product-version=git-tag
var (
	Version string
	Build   string
	//Entry   string

	/*funcs = map[string]func(){
		"f1": functionOne, "f2": functionTwo,
	}*/
)

/*
	func functionOne() {
		fmt.Println("This is function one")
	}

	func functionTwo() {
		fmt.Println("This is function two")
	}
*/

// go build -ldflags "-w -s -X main.Version=1.2.3.4 -X main.Build=17/11/2024" .
func main() {

	fmt.Println("Version: ", Version)
	fmt.Println("Build Time: ", Build)

	//funcs[Entry]()

}
