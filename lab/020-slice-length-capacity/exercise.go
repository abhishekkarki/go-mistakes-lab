package m020

// Task 1: PREDICT. Replace each 0 with the value Go computes.
// Start from:  s := make([]int, 3, 6), so s = [0, 0, 0] with room for 6.
// Lines run in order: each one sees the effects of the lines above it.
var Predictions = map[string]int{
	"1. len(s)":                      3,
	"2. cap(s)":                      6,
	"3. s2 := s[1:3]; len(s2)":       2,
	"4. cap(s2)":                     5,
	"5. cap(s[1:3:4])":               4,
	"6. s2 = append(s2, 99); len(s)": 3,
	"7. s[:4][3]":                    0,
	"8. s2 = append(s2, 1, 2, 3); s2[0] = 42; s[1]": 1,
	"9. len(s2)":                         4,
	"10. cap([]int(nil)) + cap([]int{})": 0,
}
