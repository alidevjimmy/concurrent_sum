package array_sum

import "testing"

func TestNaiveConcurrentSum(t *testing.T) {
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
		{
			name:     "non empty large array",
			arr:      make([]int, 1*1024*1024),
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := naiveConcurrentSum(test.arr)
			if sum != test.expected {
				t.Errorf("expexted %d got %d", test.expected, sum)
			}
		})
	}
}

func TestConcurrentSum(t *testing.T) {
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
		{
			name:     "non empty large array",
			arr:      make([]int, 1*1024*1024),
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := concurrentSum(test.arr)
			if sum != test.expected {
				t.Errorf("expexted %d got %d", test.expected, sum)
			}
		})
	}
}

func BenchmarkNaiveConcurrentSum(b *testing.B) {
	largeArray := make([]int, 10*1024*1024)
	for n := 0; n < b.N; n++ {
		naiveConcurrentSum(largeArray)
	}
}

func BenchmarkConcurrentSum(b *testing.B) {
	largeArray := make([]int, 10*1024*1024)
	for n := 0; n < b.N; n++ {
		concurrentSum(largeArray)
	}
}
