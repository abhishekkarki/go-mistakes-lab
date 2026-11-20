//go:build !solution

package m017

import (
	"fmt"
	"os"
	"strconv"
)

// Task 1: Predict: Replace each 0 with the value Go computes.
// Write all your guesses first, then run `make lab N=017`.
var Predictions = map[string]int{
	"010":       8,
	"0o10":      8,
	"0x10":      16,
	"0b10":      2,
	"1_000_000": 1000000,
	"100 + 010": 108,
	"0o644":     420,
}

// Task 2: FIX. ParseUserID parses a decimal, non-negative user ID.
// Leading zeros are allowed and must not change the value: "010" is ten.
// Prefixes like "0x" and digits separators like "1_000" are invalid input.
func ParseUserID(s string) (int64, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse user id %q: %w", s, err)
	}

	if id < 0 {
		return 0, fmt.Errorf("user ID cannot be negative")
	}
	return id, nil
}

// Task 3: Fix. WriteConfig writes data to path as rw-r--r--
// (the owner reads and writes; group and other only read).
func WriteConfig(path string, data []byte) error {
	return os.WriteFile(path, data, 0o644)
}
