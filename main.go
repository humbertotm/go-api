package main

import (
	"github.com/humbertotm/go-api/controllers"
)

func main() {
	router := controllers.SetupRouter()
	// Listen and Server in 0.0.0.0:3000
	router.Run(":3000")
}
