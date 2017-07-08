package leetcode

import "testing"

func TestThreeSum(t *testing.T) {
	cases := []struct {
		input []int
		want  [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			},
		},
	}
}
