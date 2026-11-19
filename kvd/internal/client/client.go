// Package client is a Go client for the kvd HTTP API.
package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Client talks to one kvd server.
type Client struct {
	baseURL string
	hc      *http.Client
}

// New returns a Client for the kvd server at baseURL.
//
// TODO(day 21, #81): http.DefaultClient has no timeouts. Build an http.Client
// with Timeout and a Transport (dial, TLS handshake, response header, idle conns).
func New(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		hc:      http.DefaultClient,
	}
}

// Health returns nil if the server answers GET /healthz with 200.
func (c *Client) Health(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/healthz", http.NoBody)
	if err != nil {
		return fmt.Errorf("build health request: %w", err)
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return fmt.Errorf("health request: %w", err)
	}
	// TODO(day 19, #54): surface the Close error instead of dropping it.
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health: unexpected status %s", resp.Status)
	}
	return nil
}

// TODO(day 22): Get, Put, Delete. Map 404 back to a sentinel error (#51).
