package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "buffered"
	messages <- "channel"
	messages <- "gucci"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
