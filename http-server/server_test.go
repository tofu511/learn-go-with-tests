package http_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T)  {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}