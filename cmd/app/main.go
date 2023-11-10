package main

import (
	"log"

	"github.com/damirsharip/contentquo-test/internal/service"
)

func main() {
	s := service.NewService()
	err := s.SortBreeds()
	if err != nil {
		log.Fatalln(err)
	}
}
