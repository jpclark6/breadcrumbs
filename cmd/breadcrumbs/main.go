// Start geochat services
package main

import (
	"github.com/jpclark6/breadcrumbs/internal"
)

func main() {
	geo.SetupDatabase()
	geo.SetupRouter()
}
