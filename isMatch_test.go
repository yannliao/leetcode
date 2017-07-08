package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "aa", true},
	}
	for _, c := range cases {
		got := IsMatch(c.s, c.p)
		if got != c.want {
			t.Errorf("IsMatch(%q, %q) == %t, want %t", c.s, c.p, got, c.want)
		}
	}
}
