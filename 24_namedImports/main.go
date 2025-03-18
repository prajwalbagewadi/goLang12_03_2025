package main

import (
	f "fmt" // Renaming the fmt package to "f"
	//blank import Sometimes, you might want to import a package just for its side effects (like initialization).
	_ "myproject/my2" //Only executes init() in mypack, but doesn't use it
	prajwal "myproject/mypack"
)

func main(){
	f.Println("named import")
	prajwal.Sayhello()
}
