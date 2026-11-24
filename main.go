package main

import "fmt"

// ContainsDuplicate reports whether any value appears more than once.
// A set of seen values answers it in one pass, exiting on the first repeat.
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 1}), ContainsDuplicate([]int{1, 2, 3, 4}))
}
