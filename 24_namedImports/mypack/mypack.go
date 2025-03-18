package mypack

import (
	"fmt"
)

func Sayhello(){
	fmt.Println("hello from mypack")
}

func init(){
    fmt.Println("init running blank import.")
}