package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"example.com/stringutil"
)

func main() {
	m := "Hello world!"
	// Show a simple hello world message.
	fmt.Printf("Forwards:\t%v\n", m)
	// Now reverse the hello world message.
	fmt.Printf("Backwards:\t%v\n", stringutil.Reverse(m))
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Dimitri", "Alejandra", "Anya"}

	// Request a greeting messages for the names.
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	for n, g := range message {
		fmt.Printf("Message for %v:\n\t%v\n", n, g)
	}
}
