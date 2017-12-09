package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	n := 37
	fmt.Println(fibSeriesRecursive(n))
	end := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	fmt.Println(end-start, "milliseconds")

	start = time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	fmt.Println(fibSeriesMemoization(n))
	end = time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	fmt.Println(end-start, "milliseconds")

	start = time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	fmt.Println(fibDynamic(n))
	end = time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	fmt.Println(end-start, "milliseconds")
}

func fibSeriesRecursive(n int) []int {
	a := make([]int, n)
	for i := 1; i <= n; i++ {
		a[i-1] = fib(i)
	}
	return a
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fibDynamic(n int) []int {
	a := []int{1}
	if n == 1 {
		return a
	}
	a = append(a, 1)
	if n == 2 {
		return a
	}
	for i := 2; i < n; i++ {
		a = append(a, a[i-2]+a[i-1])
	}
	return a
}

func fibSeriesMemoization(n int) []int {
	a := make([]int, n)
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		a[i-1] = fibMemo(i, m)
	}
	return a
}

func fibMemo(x int, m map[int]int) int {
	if x < 2 {
		m[x] = x
		return x
	}
	_, ok := m[x-1]
	if !ok {
		m[x-1] = fibMemo(x-1, m)
	}
	_, ok = m[x-2]
	if !ok {
		m[x-2] = fibMemo(x-2, m)
	}
	return m[x-1] + m[x-2]
}
