package m017

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPredictions(t *testing.T) {
	actual := map[string]int{
		"010":       0o10,
		"0o10":      0o10,
		"0x10":      0x10,
		"0b10":      0b10,
		"1_000_000": 1_000_000,
		"100 + 010": 100 + 0o10,
		"0o644":     0o644,
	}
	for expr, want := range actual {
		got, ok := Predictions[expr]
		if !ok {
			t.Errorf("%-10s missing from Predictions", expr)
			continue
		}
		if got != want {
			t.Errorf("%-10s you predicted %d, Go says %d", expr, got, want)
		}
	}
}

func TestParseUserID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in      string
		want    int64
		wantErr bool
	}{
		{in: "42", want: 42},
		{in: "007", want: 7},
		{in: "010", want: 10},
		{in: "0", want: 0},
		{in: "0x1F", wantErr: true},
		{in: "0b1", wantErr: true},
		{in: "0o17", wantErr: true},
		{in: "1_000", wantErr: true},
		{in: "-5", wantErr: true},
		{in: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			got, err := ParseUserID(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseUserID(%q) = %d, want an error", tt.in, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseUserID(%q) unexpected error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("ParseUserID(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestWriteConfigPermissions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unix permissions only")
	}
	t.Parallel()

	path := filepath.Join(t.TempDir(), "app.yaml")
	if err := WriteConfig(path, []byte("replicas: 3\n")); err != nil {
		t.Fatalf("WriteConfig: %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	perm := info.Mode().Perm()

	// Only the owner bits are checked, because your umask can clear group/other bits.
	if perm&0o700 != 0o600 {
		t.Errorf("owner bits = %v (%#o), want rw- (0o600 for the owner)", perm, perm)
	}
	if perm&0o022 != 0 {
		t.Errorf("group/other can write: %v (%#o)", perm, perm)
	}
}
