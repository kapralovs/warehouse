package main

import (
	"log"

	"github.com/kapralovs/warehouse/internal/server"
)

func main() {
	s := server.New()

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
