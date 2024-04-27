package main

import (
	"log"

	"nemesis-cli/src"
)

func main() {
	if err := src.NemesisCli(); err != nil {
		log.Fatal("Error: ", err)
		
	}
}
