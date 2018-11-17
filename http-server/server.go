package http_server

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}

	return 0
}