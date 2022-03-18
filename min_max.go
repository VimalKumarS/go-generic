package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"sort"
	"strconv"
	"strings"
)

func max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}

	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func main() {
	fmt.Println(min([]int{10, 2, 4, 1, 6, 8, 2}))
	fmt.Println(max([]float64{3.2, 5.1, 6.2, 7.6, 8.2, 1.5, 4.8}))

	vegetableSet := map[string]bool{
		"potato":  true,
		"cabbage": true,
		"carrot":  true,
	}

	fruitRank := map[int]string{
		1: "strawberry",
		2: "raspberry",
		3: "blueberry",
	}

	fmt.Printf("vegetableSet keys: %+v\n", keys(vegetableSet))
	fmt.Printf("fruitRank keys: %+v\n", keys(fruitRank))

	floatSlice := []float64{2.3, 1.2, 0.2, 51.2}
	sortSlice(floatSlice)
	fmt.Println(floatSlice)

	stringSlice := []string{"z", "a", "b"}
	sortSlice(stringSlice)
	fmt.Println(stringSlice)

	intSlice := []int{0, 3, 2, 1, 6}
	sortSlice(intSlice)
	fmt.Println(intSlice)

	websites := []string{"http://foo.com", "https://bar.com", "https://gosamples.dev"}
	httpsWebsites := filter(websites, func(v string) bool {
		return strings.HasPrefix(v, "https://")
	})
	fmt.Println(httpsWebsites)

	numbers1 := []int{1, 2, 3, 4, 5, 6}
	divisibleBy2 := filter(numbers1, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(divisibleBy2)

	numbers := []float64{4, 9, 16, 25}

	newNumbers := mapSlice(numbers, math.Sqrt)
	fmt.Println(newNumbers)

	words := []string{"a", "b", "c", "d"}
	quoted := mapSlice(words, func(s string) string {
		return "\"" + s + "\""
	})
	fmt.Println(quoted)

	stringPowNumbers := mapSlice(numbers, func(n float64) string {
		return strconv.FormatFloat(math.Pow(n, 2), 'f', -1, 64)
	})
	fmt.Println(stringPowNumbers)
}
