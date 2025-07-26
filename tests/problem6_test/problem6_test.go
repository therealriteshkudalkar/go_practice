package problem6test

import (
	"slices"
	"testing"

	"github.com/therealriteshkudalkar/go_practice/src/problem6"
)

func TestLoudAndRich(t *testing.T) {
	type args struct {
		richer [][]int
		quiet  []int
	}

	tests := []struct {
		name string
		arg  args
		want []int
	}{
		{
			"Test Case 1",
			args{[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}},
			[]int{5, 5, 2, 5, 4, 5, 6, 7},
		},
		{
			"Test Case 2",
			args{[][]int{}, []int{0}},
			[]int{0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := problem6.LoudAndRich(test.arg.richer, test.arg.quiet); !slices.Equal(got, test.want) {
				t.Errorf("LoudAndRich(): %v, want: %v", got, test.want)
			}
		})
	}
}
