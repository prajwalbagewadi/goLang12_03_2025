package main

import (
	"fmt"
	"time"
)

func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)
	ch3:=make(chan string)

	go func(){
		fmt.Println("data send ch1.")
		time.Sleep(1*time.Second)
		ch1<-"data1"
	}()

	go func(){
		fmt.Println("data send ch2.")
		time.Sleep(2*time.Second)
		ch2<-"data2"
	}()

	go func(){
		fmt.Println("data send ch3.")
		time.Sleep(2*time.Second)
		ch3<-"data3"
	}()

	select{
	case msg1 := <-ch1:
		fmt.Println("ch1 received=",msg1)
	case msg2 := <-ch2:
		fmt.Println("ch2 received=",msg2)
	case msg3 := <-ch3:
		fmt.Println("ch3 received=",msg3)
	// default:
	// 	fmt.Println("no values in channel.")
	}


}