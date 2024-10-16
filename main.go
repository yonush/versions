package main

import (
	"fmt"
)

var (
	Version string
	Build   string
	Entry   string

	funcs = map[string]func(){
		"f1": functionOne, "f2": functionTwo,
	}
)

func functionOne() {
	fmt.Println("This is function one")
}

func functionTwo() {
	fmt.Println("This is function two")
}

func main() {

	fmt.Println("Version: ", Version)
	fmt.Println("Build Time: ", Build)

	funcs[Entry]()

}
