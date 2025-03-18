package main

import (
	"fmt"
	"myproject/mypack"
)

/*
PS C:\Users\bagew\Desktop\WebDev\GoLang1\23_ExportingRules> go mod init myproject
go: creating new go.mod: module myproject
go: to add module requirements and sums:
        go mod tidy
PS C:\Users\bagew\Desktop\WebDev\GoLang1\23_ExportingRules> go run main.go
hello from mypack
PS C:\Users\bagew\Desktop\WebDev\GoLang1\23_ExportingRules>

*/
func main(){
	mypack.SayHello()
	p1:=mypack.Person{Name:"alice"}
	fmt.Println(p1)
	mypack.RunPrivate()
}