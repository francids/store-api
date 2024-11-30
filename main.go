package main

import (
	"store-api/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}
