package problem5test

import (
	"testing"

	"github.com/therealriteshkudalkar/go_practice/src/problem5"
)

func TestNetworkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n int
		k int
	}

	tests := []struct {
		name string
		arg args
		want int
	} {
		{
			"Test Case 1",
			args{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2},
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := problem5.NetworkDelayTime(test.arg.times, test.arg.n, test.args.k); got != test.want {
				t.Errorf("NetworkDelayTime() = %v, want %v", got, test.want)
			}
		})
	}
}