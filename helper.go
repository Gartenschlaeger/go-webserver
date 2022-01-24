package main

import (
	"log"
)

func must(err error) {
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
}