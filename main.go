package main

import "github.com/Oktapratama/assignment2.git/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}