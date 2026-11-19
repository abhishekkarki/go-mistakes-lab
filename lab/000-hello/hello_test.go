package hello

import "testing"

// TestToolchain proves the setup works: `make lab N=000` should pass.
// Template for every lab: show the bug in one test, the fix in another.
func TestToolchain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"zero", 0, 0, 0},
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.a + tt.b; got != tt.want {
				t.Fatalf("%d + %d = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
