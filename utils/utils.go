package utils

import (
	"log"
)

type Position struct {
	X, Y int
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
