package main

import (
	"log"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/driver"
	_ "github.com/lib/pq"
)

func main() {
	connection, err := driver.connectDb()
	if err != nil {
		log.Fatal("Error at connection : ", err)
	}
}
