package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		height []int
		want   int
	}{
		{[]int{1}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 1, 2}, 4},
	}
	for _, c := range cases {
		got := maxArea(c.height)
		if got != c.want {
			t.Errorf("MaxArea(%v) == %d, want %d", c.height, got, c.want)
		}
	}
}
