package problem2_test

import (
	"github.com/therealriteshkudalkar/go_practice/src/problem2"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				s: "babad",
			},
			"bab",
		},
		{
			"test2",
			args{
				s: "cbbd",
			},
			"bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2.LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
