package main

import (
	"fmt"

	"github.com/jeynesrya/adalpha-solutions/api"
)

func main() {
	fmt.Println("Starting application...")

	// Initialise application
	a := api.Api{}
	a.Initialise()

	// Run application
	a.Run(":8080")
}
