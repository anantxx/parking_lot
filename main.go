package main

import (
	"flag"

	"github.com/parking_lot/handler"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		// ToDo: File Implement later
		handler.ExecuteFile(flag.Args()[0])
		return
	}
	handler.ExecuteCommand()
}
