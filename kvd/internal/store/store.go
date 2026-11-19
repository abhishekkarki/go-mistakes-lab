// Package store is kvd's in-memory key-value store with per-key TTLs.
//
// Roadmap (workbook days and the mistakes each one applies):
//
//	Day 6      #24 #25 #27 #28  map + TTL, copy on Set/Get, size hint, Compact
//	Day 12     #60 #62          janitor goroutine with Close(), RWMutex, goleak
//	Day 17     #73              N shards, MGet via errgroup, Watch(key)
//	Day 24     #87              injectable clock for instant TTL tests
//	Day 28-29  #94-97           field alignment, escape analysis, sync.Pool
//	Day 34     #11              New(opts ...Option)
package store

import (
	"errors"
	"time"
)

// ErrNotFound is returned when a key does not exist or has expired.
var ErrNotFound = errors.New("store: key not found")

// errNotImplemented marks the skeleton. Delete it once Day 6 is done;
// the tests that skip on it will then start running.
var errNotImplemented = errors.New("store: not implemented")

// Store is an in-memory key-value store.
type Store struct {
	// TODO(day 6): items map[string]item, where item holds the value and expiresAt.
}

// New returns an empty Store.
//
// TODO(day 6, #27): take a size hint and preallocate the map.
// TODO(day 34, #11): switch to functional options.
func New() *Store {
	return &Store{}
}

// Get returns a copy of the value stored under key, or ErrNotFound.
//
// TODO(day 6, #24): return a copy so callers can't mutate the stored bytes.
func (s *Store) Get(key string) ([]byte, error) {
	_ = key
	return nil, errNotImplemented
}

// Set stores a copy of value under key. A ttl <= 0 means no expiry.
//
// TODO(day 6, #25): copy value; the caller may reuse its slice.
func (s *Store) Set(key string, value []byte, ttl time.Duration) error {
	_, _, _ = key, value, ttl
	return errNotImplemented
}

// Delete removes key. Deleting a missing key is not an error.
//
// TODO(day 6, #28): add Compact() to rebuild the map after mass deletes.
func (s *Store) Delete(key string) error {
	_ = key
	return errNotImplemented
}
