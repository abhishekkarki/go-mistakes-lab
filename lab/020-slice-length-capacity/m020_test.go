package m020

import (
	"bytes"
	"slices"
	"strconv"
	"testing"
)

func TestPredictions(t *testing.T) {
	actual := map[string]int{}

	s := make([]int, 3, 6)
	actual["1. len(s)"] = len(s)
	actual["2. cap(s)"] = cap(s)

	s2 := s[1:3]
	actual["3. s2 := s[1:3]; len(s2)"] = len(s2)
	actual["4. cap(s2)"] = cap(s2)
	actual["5. cap(s[1:3:4])"] = cap(s[1:3:4])

	s2 = append(s2, 99)
	actual["6. s2 = append(s2, 99); len(s)"] = len(s)
	actual["7. s[:4][3]"] = s[:4][3]

	s2 = append(s2, 1, 2, 3)
	s2[0] = 42
	actual["8. s2 = append(s2, 1, 2, 3); s2[0] = 42; s[1]"] = s[1]
	actual["9. len(s2)"] = len(s2)

	actual["10. cap([]int(nil)) + cap([]int{})"] = cap([]int(nil)) + cap([]int{})

	for expr, want := range actual {
		got, ok := Predictions[expr]
		if !ok {
			t.Errorf("%-48s missing from Predictions", expr)
			continue
		}
		if got != want {
			t.Errorf("%-48s you predicted %d, Go says %d", expr, got, want)
		}
	}
}

func TestHead(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"first two", 2, []int{1, 2}},
		{"zero", 0, []int{}},
		{"all", 5, []int{1, 2, 3, 4, 5}},
		{"more than len", 10, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			src := make([]int, 5, 16) // spare capacity is where the bug hides
			copy(src, []int{1, 2, 3, 4, 5})
			orig := slices.Clone(src)

			got := Head(src, tt.n)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("Head(src, %d) = %v, want %v", tt.n, got, tt.want)
			}
			if len(got) > 0 && &got[0] != &src[0] {
				t.Fatalf("Head copied the data; it should share src's backing array")
			}

			_ = append(got, -1, -2) // the caller owns the result
			if !slices.Equal(src, orig) {
				t.Fatalf("appending to Head's result changed src: %v, was %v", src, orig)
			}
		})
	}
}

func TestIDs(t *testing.T) {
	// Not parallel: AllocsPerRun counts allocations across the whole process.
	users := []User{{"u1", "Ana"}, {"u2", "Ben"}, {"u3", "Chen"}}

	got := IDs(users)
	want := []string{"u1", "u2", "u3"}
	if !slices.Equal(got, want) {
		t.Fatalf("IDs = %q (len %d), want %q", got, len(got), want)
	}

	allocs := testing.AllocsPerRun(100, func() { _ = IDs(users) })
	if allocs != 1 {
		t.Fatalf("IDs made %.0f allocations per call, want exactly 1", allocs)
	}
	if got := IDs(nil); len(got) != 0 {
		t.Fatalf("IDs(nil) = %q, want empty", got)
	}
}

func TestBufferReuse(t *testing.T) {
	// Not parallel: AllocsPerRun counts allocations across the whole process.
	req := bytes.Repeat([]byte("x"), 4096)

	var b Buffer
	b.Write(req[:10])
	b.Reset()
	if len(b.Bytes()) != 0 {
		t.Fatalf("after Reset, len = %d, want 0", len(b.Bytes()))
	}

	b.Write(req) // warm up: grow once to fit a full request
	b.Reset()

	allocs := testing.AllocsPerRun(100, func() {
		b.Write(req)
		b.Reset()
	})
	if allocs != 0 {
		t.Fatalf("refilling a warm Buffer made %.0f allocations, want 0", allocs)
	}

	b.Write([]byte("abc"))
	if got := string(b.Bytes()); got != "abc" {
		t.Fatalf("Bytes() = %q, want %q", got, "abc")
	}
}

func TestGrowthSteps(t *testing.T) {
	t.Parallel()
	const n = 2000
	steps := GrowthSteps(n)
	if len(steps) == 0 {
		t.Fatal("GrowthSteps returned nothing: implement Task 5")
	}
	if steps[0] != 1 {
		t.Errorf("first capacity = %d, want 1", steps[0])
	}
	if !slices.IsSorted(steps) || len(slices.Compact(slices.Clone(steps))) != len(steps) {
		t.Errorf("capacities must strictly increase: %v", steps)
	}
	if last := steps[len(steps)-1]; last < n {
		t.Errorf("final capacity %d < %d elements", last, n)
	}
	// Small slices double; past 256 elements growth slows toward 1.25x.
	for i := 1; i < len(steps); i++ {
		ratio := float64(steps[i]) / float64(steps[i-1])
		if steps[i-1] < 256 && ratio != 2 {
			t.Errorf("growth %d -> %d is %.2fx; below 256 it should double", steps[i-1], steps[i], ratio)
		}
		if steps[i-1] >= 512 && ratio >= 2 {
			t.Errorf("growth %d -> %d is %.2fx; large slices grow by less than 2x", steps[i-1], steps[i], ratio)
		}
	}

	var line []byte
	for i, c := range steps {
		if i > 0 {
			line = append(line, " -> "...)
		}
		line = strconv.AppendInt(line, int64(c), 10)
	}
	t.Logf("capacities while appending %d ints: %s", n, line)
}

// Run with: make bench PKG=lab/020-slice-length-capacity
func BenchmarkAppendNoHint(b *testing.B) {
	for b.Loop() {
		var s []int
		for i := range 10_000 {
			s = append(s, i)
		}
		_ = s
	}
}

func BenchmarkAppendWithCap(b *testing.B) {
	for b.Loop() {
		s := make([]int, 0, 10_000)
		for i := range 10_000 {
			s = append(s, i)
		}
		_ = s
	}
}
