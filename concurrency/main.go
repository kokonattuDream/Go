package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! ....", phrase)
	doneChan <- true
}

func main() {

	dones := make([]chan bool, 4)
	for i := 0; i < len(dones); i++ {
		dones[i] = make(chan bool)
	}
	go greet("Nice to meet you?", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ....are.... you....?", dones[2])
	go greet("i hope you like it", dones[3])

	for _, done := range dones {
		<-done
	}
}
