package problem7_test

import (
	"testing"

	"github.com/therealriteshkudalkar/go_practice/src/problem7"
)

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		{
			"Test Case 1",
			[][]int{{1}, {2}, {3}, {}},
			true,
		},
		{
			"Test Case 2",
			[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := problem7.CanVisitAllRooms(test.rooms); got != test.want {
				t.Errorf("CanVisitAllRooms: %v, but want: %v", test.want, got)
			}
		})
	}
}
