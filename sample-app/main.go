package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
	a, b := 0, 1
	count := 0
	return func() int {
		if count == 0 || count == 1 {
			count++
			return count - 1
		}
		result := a + b
		a = b
		b = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
