package problem4_test

import (
	"github.com/therealriteshkudalkar/go_practice/src/problem4"
	"testing"
)

func TestMyPowRec(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Test Case 1",
			args{2.0, 10},
			1024.0,
		},
		{
			"Test Case 2",
			args{2.1, 3},
			9.261000000000001,
		},
		{
			"Test Case 3",
			args{2.0, -2},
			0.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem4.MyPowRec(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("MyPowRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyPowRecMemoized(t *testing.T) {
	type args struct {
		memo map[string]float64
		x    float64
		n    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Test Case 1",
			args{map[string]float64{}, 2.0, 10},
			1024.0,
		},
		{
			"Test Case 2",
			args{map[string]float64{}, 2.1, 3},
			9.261000000000001,
		},
		{
			"Test Case 3",
			args{map[string]float64{}, 2.0, -2},
			0.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem4.MyPowRecMemoized(tt.args.memo, tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("MyPowRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
