package main

import (
	"fmt"
	"os"
)

func main(){
	//create file
	file,err := os.Create("example.txt")
	if err!=nil {
		fmt.Println("error in file:",err)
		return
	}
	file.Close()

	//writing file
	data:=[]byte("welcome to golang.")
	err = os.WriteFile("example.txt",data,0644)
	/*
	os.WriteFile("example.txt", data, 0644)
	filename → Name of the file to write.
	data → Data to be written (must be a []byte slice).
	perm → File permission mode (e.g., 0644 for read/write permissions).
	*/
	if err!=nil{
		fmt.Println("error writing file:",err)
		return
	}
	fmt.Println("data written successfully.")

	//reading file
	data,err = os.ReadFile("example.txt")
	if err!=nil{
		fmt.Println("error reading file:",err)
		return
	}
	fmt.Println("file content:",string(data))

	//file close
	file.Close()
}