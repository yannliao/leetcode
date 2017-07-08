package leetcode

import "testing"

func TestIntToRome(t *testing.T) {
	cases := []struct {
		num  int
		want string
	}{
		{34, "XXXIV"},
		{1666, "MDCLXVI"},
		{1984, "MCMLXXXIV"},
		{1976, "MCMLXXVI"},
		{90, "XC"},
	}
	for _, c := range cases {
		got := intToRoman(c.num)
		if got != c.want {
			t.Errorf("intToRoman(%d) == %q, want %q", c.num, got, c.want)
		}
	}
}
