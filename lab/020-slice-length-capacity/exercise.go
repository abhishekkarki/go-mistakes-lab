//go:build solution

// Reference solution. Open it only after your tests pass:
//   go test -tags solution ./lab/020-slice-length-capacity/

package m020

var Predictions = map[string]int{
	"1. len(s)":                      3,
	"2. cap(s)":                      6,
	"3. s2 := s[1:3]; len(s2)":       2,
	"4. cap(s2)":                     5,  // capacity runs to the end of the backing array: 6 - 1
	"5. cap(s[1:3:4])":               3,  // full slice expression s[low:high:max] sets cap = max - low
	"6. s2 = append(s2, 99); len(s)": 3,  // s's header didn't change...
	"7. s[:4][3]":                    99, // ...but its backing array did: s2's append wrote into s's spare capacity
	"8. s2 = append(s2, 1, 2, 3); s2[0] = 42; s[1]": 0, // len 6 > cap 5, so s2 moved to a new array; s is untouched
	"9. len(s2)":                         6,
	"10. cap([]int(nil)) + cap([]int{})": 0,
}

// Head caps the result's capacity at n, so the caller's first append has to
// reallocate instead of writing over s[n].
func Head(s []int, n int) []int {
	n = max(0, min(n, len(s)))
	return s[:n:n]
}

type User struct {
	ID   string
	Name string
}

// IDs sets the capacity, not the length. make([]string, len(users)) gave
// len(users) empty strings, and append added after them.
func IDs(users []User) []string {
	ids := make([]string, 0, len(users))
	for _, u := range users {
		ids = append(ids, u.ID)
	}
	return ids
}

type Buffer struct {
	data []byte
}

func (b *Buffer) Write(p []byte) {
	b.data = append(b.data, p...)
}

func (b *Buffer) Bytes() []byte {
	return b.data
}

// Reset keeps the backing array: length 0, capacity unchanged. Setting it to
// nil threw away the capacity, so the next request grew it again from zero.
// (Callers must not keep Bytes() across a Reset; the next Write overwrites it.)
func (b *Buffer) Reset() {
	b.data = b.data[:0]
}

// GrowthSteps records cap(s) each time append moves the slice to a new array.
func GrowthSteps(n int) []int {
	var s []int
	var steps []int
	for i := range n {
		before := cap(s)
		s = append(s, i)
		if cap(s) != before {
			steps = append(steps, cap(s))
		}
	}
	return steps
}
