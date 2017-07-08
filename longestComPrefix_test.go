package leetcode

import "testing"

func TestLongestComPrefix(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"a", "b"}, ""},
		{[]string{"a"}, "a"},
	}
	for _, c := range cases {
		got := longestCommonPrefix(c.input)
		if got != c.want {
			t.Errorf("longestCommonPrefix(%v) == %q, want %q", c.input, got, c.want)
		}
	}

}
