package main

import (
	f "fmt" // Renaming the fmt package to "f"
	prajwal "myproject/mypack"
)

func main(){
	f.Println("named import")
	prajwal.Sayhello()
}
