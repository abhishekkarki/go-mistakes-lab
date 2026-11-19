package store

import (
	"bytes"
	"errors"
	"testing"
)

// skipIfSkeleton keeps CI green until Day 6. Once errNotImplemented is gone,
// delete this helper and these tests become your Day 6 acceptance tests.
func skipIfSkeleton(t testing.TB, err error) {
	t.Helper()
	if errors.Is(err, errNotImplemented) {
		t.Skip("TODO(day 6): store not implemented yet")
	}
}

func TestSetGet(t *testing.T) {
	t.Parallel()
	s := New()

	err := s.Set("k", []byte("v"), 0)
	skipIfSkeleton(t, err)
	if err != nil {
		t.Fatalf("Set: %v", err)
	}

	got, err := s.Get("k")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !bytes.Equal(got, []byte("v")) {
		t.Fatalf("Get = %q, want %q", got, "v")
	}
}

func TestGetMissing(t *testing.T) {
	t.Parallel()
	s := New()

	_, err := s.Get("missing")
	skipIfSkeleton(t, err)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Get(missing) err = %v, want ErrNotFound", err)
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	s := New()

	err := s.Set("k", []byte("v"), 0)
	skipIfSkeleton(t, err)
	if err := s.Delete("k"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := s.Get("k"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Get after Delete err = %v, want ErrNotFound", err)
	}
}

// TODO(day 6, #25): TestSetCopiesValue: mutate the slice passed to Set, Get must be unchanged.
// TODO(day 6, #24): TestGetReturnsCopy: mutate the slice from Get, the next Get must be unchanged.
// TODO(day 12): TestExpiry, then on Day 24 rewrite it with an injected clock (no sleeps, #86).

func BenchmarkSet(b *testing.B) {
	s := New()
	val := bytes.Repeat([]byte("x"), 256)
	skipIfSkeleton(b, s.Set("warmup", val, 0))

	for b.Loop() { // Go 1.24: b.Loop avoids the classic #89 benchmark traps
		_ = s.Set("k", val, 0)
	}
}
