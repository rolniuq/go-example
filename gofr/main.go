package main

import "gofr.dev/pkg/gofr"

func main() {
	// Create a new application
	a := gofr.New()

	// Add the routes
	a.GET("/hello", func(c *gofr.Context) (interface{}, error) {
		return "Hello World", nil
	})

	// Run the application
	a.Run()
}
