package main

import (
	"github.com/tofu511/learn-go-with-tests/http-server"
	"log"
	"net/http"
)

func main()  {
	server := &http_server.PlayerServer{}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
