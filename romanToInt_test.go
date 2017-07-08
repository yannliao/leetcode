package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{
			"DCXXI", 621,
		}, {
			"XCV", 95,
		}, {
			"XCVIII", 98,
		}, {
			"MDCLXVI", 1666,
		}, {
			"MDCCCLXXXVIII", 1888,
		},
	}
	for _, c := range cases {
		got := romanToInt(c.input)
		if got != c.want {
			t.Errorf("romanToInt(%q) == %d, want %d", c.input, got, c.want)
		}
	}
}
