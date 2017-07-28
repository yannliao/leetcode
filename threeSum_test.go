package leetcode

import (
	"reflect"
	"testing"
)

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
	for _, c := range cases {
		got := threeSum(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("threeSum(%q) == %q, want %q", c.input, got, c.want)
		}
	}
}
