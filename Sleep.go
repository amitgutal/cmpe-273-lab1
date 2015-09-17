package main

import	"fmt"
import "time"

func FunctionSleep(seconds int) {
	
    <-time.After(time.Second * time.Duration(seconds))
}

func main(){
	
	var input int
	input = 3
	FunctionSleep(input)
	fmt.Println("Slept for %f, seconds",input)
}