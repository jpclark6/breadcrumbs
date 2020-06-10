// Start geochat services
package main

import (
	"fmt"
	
	"github.com/jpclark6/breadcrumbs/internal"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	geo.SetupDatabase()
	geo.SetupRouter()
}
