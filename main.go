package main

import "tugasdua/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}