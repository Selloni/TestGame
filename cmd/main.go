package main

import (
	"WB/interal"
	"WB/pkg/handler"
	"WB/pkg/posgresql"
	"context"
	"log"
)

func main() {
	conn, err := posgresql.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tmp := interal.Model{}
	newHandler := handler.NewHandler(context.Background(), conn, tmp)
	err = newHandler.Route()
	if err != nil {
		log.Fatal(err)
	}
}
