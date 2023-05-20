package main

import (
	"fmt"
	"grpc/pkg/postgres"
	"log"
)

func main() {
	db, err := postgres.OpenDB("localhost", 5432, "postgres", "postgres", "example")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("DB successfully started")
	defer db.Close()
}
