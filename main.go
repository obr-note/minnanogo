package main

import (
	"fmt"
	"sort"
)

// type IntSlice []int

// func (p IntSlice) Len() int {
// 	return len(p)
// }
// func (p IntSlice) Less(i, j int) bool {
// 	return p[i] < p[j]
// }
// func (p IntSlice) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

func main() {
	numbers := []int{10, 8, 2, 5, 1, 3, 4, 9, 7, 6}
	sort.Sort(sort.IntSlice(numbers))
	fmt.Println(numbers)
}
