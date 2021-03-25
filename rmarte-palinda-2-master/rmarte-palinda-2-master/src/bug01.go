package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	// The code doesn't work because ch is an unbuffered channel. Therefore it requires a reciever to exist when sending to the channel.
	// OLD
	// ch := make(chan string)
	// By replacing it with a buffered channel we can send to the channel without having a reciever ready
	// NEW
	ch := make(chan string, 1)
	//

	ch <- "Hello world!"
	fmt.Println(<-ch)
}
