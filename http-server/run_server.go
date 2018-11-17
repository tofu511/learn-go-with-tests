package http_server

import (
	"log"
	"net/http"
)

type InMemoryPlayerScore struct {
}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return 123
}

func RunServer()  {
	server := &PlayerServer{&InMemoryPlayerScore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
