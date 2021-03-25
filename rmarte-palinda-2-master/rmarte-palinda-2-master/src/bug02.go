package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
// After completing the loop the main function closes the channel and is then complete. When the main goroutine ends all other running goroutines are cancelled. To fix this we need to make sure that main does not close until all other goroutines have completed. Since we only have one goroutine this can easily be achieved by giving Print another channel to which it can send a message when done. Since recieving from a channel blocks further execution until data is recieved main will not terminate until it has recieved a message.
func main() {
	ch, done := make(chan int), make(chan bool)
	go Print(ch, done) // fix
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)

    <-done
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, done chan bool) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
    done <- true
}
