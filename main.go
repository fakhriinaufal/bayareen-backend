package main

import (
	"bayareen-backend/driver"
	"bayareen-backend/routes"
)

func main() {
	driver.InitDB()

	e := routes.New()

	e.Start(":8080")
}
