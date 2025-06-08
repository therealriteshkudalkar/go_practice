package problem3_test

import (
	"github.com/therealriteshkudalkar/go_practice/src/problem3"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Normal Case",
			args{10},
			4,
		},
		{
			"Edge Case 1",
			args{0},
			0,
		},
		{
			"Edge Case 2",
			args{1},
			0,
		},
		{
			"Edge Case 3",
			args{2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem3.CountPrimes(tt.args.n); got != tt.want {
				t.Errorf("CountPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
