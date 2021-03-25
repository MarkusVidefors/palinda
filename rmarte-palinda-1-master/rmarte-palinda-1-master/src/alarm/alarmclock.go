package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		h, m, s := time.Now().Clock()
		fmt.Printf("The time is %v.%v.%v: %v\n", h, m, s, text)
		time.Sleep(delay)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", 60*time.Second)
}
