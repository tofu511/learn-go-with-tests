package http_server

import (
	"log"
	"net/http"
)

type InMemoryPlayerScore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerScore {
	return &InMemoryPlayerScore{map[string]int{}}
}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerScore) RecordWin(name string)  {
	i.store[name]++
}

func (i *InMemoryPlayerScore) GetLeague() []Player {
	return nil
}

func RunServer()  {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
