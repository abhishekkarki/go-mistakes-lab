// Package api exposes the kvd store over HTTP.
package api

import (
	"net/http"

	"github.com/abhishekkarki/go-mistakes-lab/kvd/internal/store"
)

// NewHandler returns kvd's HTTP handler.
//
// TODO(day 22): add GET, PUT and DELETE /keys/{key} using Go 1.22 mux patterns.
// Map errors.Is(err, store.ErrNotFound) to 404. #80: return after http.Error.
// TODO(day 33, #6): depend on a small consumer-side Store interface defined in
// this package instead of *store.Store.
func NewHandler(s *store.Store) http.Handler {
	_ = s // used once the /keys routes exist

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	})
	return mux
}
