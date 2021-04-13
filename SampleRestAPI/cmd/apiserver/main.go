package main

import (
	"log"
	"sample-rest-api/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	// teststuff.mainBookRest()
	// teststuff.mainHttp()
}
