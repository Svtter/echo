package main

import "os"

func main() {
	hello("world")

	arg := os.Args[1]

	if arg == "s" {
		server()
	} else {
		client()
	}

}
