package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Setup")
}

func shutdown() {
	fmt.Println("Shut Down")
}
