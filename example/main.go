package main

import (
	"github.com/albrow/gopherjs-router"
	"github.com/gopherjs/jquery"
)

var jq = jquery.NewJQuery

func main() {
	print("Starting...")

	// Create a new router
	r := router.New()

	// Add some routes. For now, they
	// will just write messages to the console
	r.HandleFunc("/", func() {
		print("At home page!")
		jq("#current-page").SetHtml("Home Page")
	})
	r.HandleFunc("/about", func() {
		print("At about page!")
		jq("#current-page").SetHtml("About Page")
	})
	r.HandleFunc("/faq", func() {
		print("At faq page!")
		jq("#current-page").SetHtml("FAQ Page")
	})

	r.Start()
}
