package main

import (
	"fmt"
	"time"
)

func main(){
	timer := time.NewTimer(3*time.Second) //timer for 3 sec
	fmt.Println("waiting for timer to expire.")
	<-timer.C //blocking wait for timer
	fmt.Println("timer Expired.")

	timer2 := time.NewTimer(5*time.Second) //timer for 5 sec
	go func(){
		<-timer2.C //blocking wait for timer
		fmt.Println("timer2 Expired.")
	}()
	time.Sleep(2*time.Second)
	stopped := false 
	//stopped = timer2.Stop()
	timer2.Reset(3*time.Second) //if you want to reuse a timer, use timer.Reset() instead of creating a new one
	if stopped {
		fmt.Println("timer2 is stopped.")
	}else {
		fmt.Println("Timer2 already expired")
	}
}