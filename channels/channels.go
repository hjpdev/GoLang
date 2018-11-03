package main // Channels are the pipes that connect concurrent goroutines
// You can send values into channels from one goroutine and receive those values into another goroutine
import "fmt"

func main() {

	messages := make(chan string) // Create a new channel with make(chan val-type)
	// Channels are typed by the values they convey
	go func() { messages <- "ping" }() // Send a value into a channel using the channel <- syntax
	// The line above sends "ping" to the messages channel made above
	msg := <-messages
	fmt.Println(msg)
}
