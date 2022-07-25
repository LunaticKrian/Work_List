package main

import (
	config "awesomeProject/config/local"
	"awesomeProject/routes"
)

func main() {
	config.Init()

	r := routes.NewRouter()
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
