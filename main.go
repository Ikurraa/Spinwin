package main

import (
	"awesomespinner/Config"
	"awesomespinner/Routes"
)

func main() {
	db := Config.ConnectDatabase()
	r := Routes.Routes(db)
	r.Run(":5000")
}
