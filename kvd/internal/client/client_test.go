package client_test

import (
	"net/http/httptest"
	"testing"

	"github.com/abhishekkarki/go-mistakes-lab/kvd/internal/api"
	"github.com/abhishekkarki/go-mistakes-lab/kvd/internal/client"
	"github.com/abhishekkarki/go-mistakes-lab/kvd/internal/store"
)

func TestHealth(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(api.NewHandler(store.New()))
	t.Cleanup(srv.Close)

	if err := client.New(srv.URL).Health(t.Context()); err != nil {
		t.Fatalf("Health: %v", err)
	}
}
