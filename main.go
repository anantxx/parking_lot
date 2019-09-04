package main

import (
	"flag"

	"github.com/parking_lot/handler"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		// ToDo: File Implement later
		handler.HandleFile(flag.Args()[0])
		return
	}

	handler.HandleCommand()
}
