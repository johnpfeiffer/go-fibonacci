package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFibonacciRecursive(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.n), func(t *testing.T) {
			result := fib(tc.n)
			if result != tc.expected {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
	m := make(map[int]int)
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.n), func(t *testing.T) {
			result := fibMemo(tc.n, m)
			if result != tc.expected {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

func TestFibSeries(t *testing.T) {
	var testCases = []struct {
		n        int
		expected []int
	}{
		{n: 1, expected: []int{1}},
		{n: 2, expected: []int{1, 1}},
		{n: 3, expected: []int{1, 1, 2}},
		{n: 4, expected: []int{1, 1, 2, 3}},
		{n: 5, expected: []int{1, 1, 2, 3, 5}},
		{n: 6, expected: []int{1, 1, 2, 3, 5, 8}},
		{n: 7, expected: []int{1, 1, 2, 3, 5, 8, 13}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.n), func(t *testing.T) {
			result := fibSeriesRecursive(tc.n)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.n), func(t *testing.T) {
			result := fibSeriesMemoization(tc.n)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.n), func(t *testing.T) {
			result := fibDynamic(tc.n)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

// go test -v -run=NOMATCH -bench=.
// go test -v -run=NOMATCH -bench=BenchmarkFibonacciSeriesRecursive

func BenchmarkFibonacciSeriesRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibSeriesRecursive(20)
	}
}

func BenchmarkFibonacciSeriesMemoization(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibSeriesMemoization(20)
	}
}

func BenchmarkFibonacciSeriesDynamicProgramming(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibDynamic(20)
	}
}
