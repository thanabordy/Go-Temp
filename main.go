package main

import (
	"log"

	"app/app/console"
	_ "app/config"
)

func main() {
	if err := console.Execute(); err != nil {
		log.Fatalln(err)
	}
}
