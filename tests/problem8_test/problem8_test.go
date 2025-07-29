package problem8_test

import (
	"testing"

	"github.com/therealriteshkudalkar/go_practice/src/problem8"
)

func TestEquationsPossible(t *testing.T) {
	tests := []struct {
		name      string
		equations []string
		want      bool
	}{
		{
			"Test Case 1",
			[]string{"a==b", "a!=b"},
			false,
		},
		{
			"Test Case 2",
			[]string{"a==b", "b==a"},
			true,
		},
		{
			"Test Case 3",
			[]string{"a==b", "b==c", "c==a"},
			true,
		},
		{
			"Test Case 4",
			[]string{"a==b", "b!=c", "a==c"},
			false,
		},
		{
			"Test Case 5",
			[]string{"c==c", "b==d", "x!=z"},
			true,
		},
		{
			"Test Case 6",
			[]string{"a==b", "b==c", "d==e", "e==f", "d==a", "f!=a"},
			false,
		},
		{
			"Test Case 7",
			[]string{"b!=f", "c!=e", "f==f", "d==f", "b==f", "a==f"},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := problem8.EquationsPossible(test.equations); got != test.want {
				t.Errorf("EquationsPossible(): %v, but want: %v", got, test.want)
			}
		})
	}
}
