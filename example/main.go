package main

import (
	"fmt"
	"github.com/albrow/gopherjs-router"
	"time"
)

func main() {
	fmt.Println("Starting...")

	// Create a new router
	r := router.New()

	// Add some routes. For now, they
	// will just write messages to the console
	r.HandleFunc("/", func() {
		fmt.Println("At home page!")
	})
	r.HandleFunc("/about", func() {
		fmt.Println("At about page!")
	})
	r.HandleFunc("/faq", func() {
		fmt.Println("At faq page!")
	})

	time.Sleep(1 * time.Second)

	// Start listening for changes
	r.Start()
}
