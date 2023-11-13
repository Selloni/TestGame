package main

import (
	"WB/pkg/handler"
	"log"
)

func main() {
	err := handler.Route()
	if err != nil {
		log.Fatal(err)
	}
}
