package main

import (
	"log"

	"github.com/kapralovs/warehouse/internal/server"
)

func main() {
	// Создаём новый сервер
	s := server.New()

	// Запускаем сервер
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
