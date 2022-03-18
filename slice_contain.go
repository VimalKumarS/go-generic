package main

import "fmt"

func contains[T comparable](elms []T, v T) bool {
	for _, s := range elms {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(contains([]string{"a", "b", "c"}, "b"))
	fmt.Println(contains([]int{1, 2, 3}, 2))
	fmt.Println(contains([]int{1, 2, 3}, 10))
}
