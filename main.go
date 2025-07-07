package main

import (
	"ascii-art/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("invalid number of aruments ...")
		return
	}
	server.StartServer()
}
