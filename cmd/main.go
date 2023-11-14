package main

import (
	"WB/interal"
	"WB/interal/posgresql"
	"WB/pkg/handler"
	"context"
	"log"
)

func main() {
	conn, err := posgresql.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	tmp := interal.Model{}
	newHandler := handler.NewHandler(context.Background(), conn, tmp)
	err = newHandler.Route()
	if err != nil {
		log.Fatal(err)
	}
}
