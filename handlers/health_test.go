package handlers_test

import (
	"canvas/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/matryer/is"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		is := is.New(t)

		mux := chi.NewMux()
		handlers.Health(mux)
		code, _, _ := makeGetRequest(mux, "/health")
		is.Equal(http.StatusOK, code)
	})
}

// makeGetRequest and returns the status code, response headers, and the body.
func makeGetRequest(handler http.Handler, target string) (int, http.Header, string) {

}
