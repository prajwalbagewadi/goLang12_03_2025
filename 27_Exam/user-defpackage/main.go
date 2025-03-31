package main

import (
	"user-defpackage/geo"
)

func main(){
	var w,l int = 5,5
	geo.Area(w,l)
}

/*
PS C:\Users\bagew\Desktop\WebDev\GoLang1\27_Exam\user-defpackage>go mod init user-defpackage

*/
// Output:
//go run main.go
//Area of rectangle is 25
