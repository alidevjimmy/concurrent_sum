package main

import "testing"

func TestNaiveCuncurrentSum(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "empty array",
			arr:      []int{},
			expected: 0,
		},
		{
			name:     "non empty array",
			arr:      []int{1, 2, 3, 4},
			expected: 10,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := naiveCuncurrentSum(test.arr)
			if sum != test.expected {
				t.Errorf("expexted %d got %d", test.expected, sum)
			}
		})
	}
}
