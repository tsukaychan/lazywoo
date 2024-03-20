package main

import (
	"log"

	"github.com/tsukaychan/lazywoo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("[lazywoo] execute failed: %v", err)
	}
}
