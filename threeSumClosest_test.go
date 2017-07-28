package leetcode

import "testing"

func TestThreeSumClosest(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
	}
	for _, c := range cases {
		got := threeSumClosest(c.input, c.target)
		if got != c.want {
			t.Errorf("threeSumClosest(%q) == %d, want %d", c.input, got, c.want)
		}
	}
}
