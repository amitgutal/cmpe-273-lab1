package main

import "testing"

type FibonacciTest struct {
        number        float64	
        fibonacci_sum float64
}

var TestCases = []FibonacciTest {
	
        {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8},{7,13}, {8,21}, {9,3}, {10, 55},
}

func TestFibonnaci(t *testing.T) {
	
        for _, count := range TestCases {
			
                actual := Fibonacci(count.number)
                if actual != count.fibonacci_sum {
					 
                       t.Errorf(" Test Failed for Fib(%f)",count.number)
                }
        }
}