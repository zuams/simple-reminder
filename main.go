package main

import (
	"log"

	"github.com/zuams/simple-reminder/db"
	"github.com/zuams/simple-reminder/routes"
)

func main() {
	db, err := db.New()
	logFatal(err)

	db.LogMode(true)
	defer db.Close()

	e := routes.Init()

	err = e.Start(":3000")
	logFatal(err)
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
