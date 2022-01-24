package helper

import "log"

func Must(err error) {
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
}
