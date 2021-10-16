package main

import (
	"log"

	"bitbucket.org/adhikag24/golang-api/app"
)

func main() {
	log.Print("Prepare db")
	if err := app.PrepareApp(); err != nil {
		log.Fatal(err)
	}

	app.StartApp()
}
