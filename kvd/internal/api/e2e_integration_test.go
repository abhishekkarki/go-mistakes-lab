//go:build integration

package api_test

import "testing"

// Runs only with `make integration` (go test -tags integration). See #82.
func TestEndToEnd(t *testing.T) {
	t.Skip("TODO(day 22): start a real server on a random port, then set, get, 404, shutdown")
}
