package main

import (
	"fmt"

	_ "blankimport/modA"
)

func main() {
	fmt.Print("Hello, World!")
}

/*
go run main.go
Import> go run main.go                           BlankImport package initialized
Hello, World!

*/