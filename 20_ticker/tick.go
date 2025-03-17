package main

import (
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTicker(2*time.Second)
	defer ticker.Stop()
	go func(){
		for t:=range ticker.C{
			fmt.Println("tick=",t)
		}
	}()
	time.Sleep(10*time.Second)
	fmt.Println("stopping ticker")
}