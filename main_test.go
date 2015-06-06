package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	err := exec.Command("go", "build", "-o", "tmp/imdb-tags").Run()
	if err != nil {
		fmt.Println("Failed to build imdb-tags binary:", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
