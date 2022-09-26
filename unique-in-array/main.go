package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 500, 600, 400, 300, 500}
	fmt.Println(isUniqueThird(a))
}

// func isUniqueFirst(a []int) bool {
// 	for i := 0; i < len(a); i++ {
// 		for j := 0; j < len(a); j++ {
// 			if i == j {
// 				continue
// 			} else {
// 				if a[i] == a[j] {
// 					return true
// 				}
// 			}

// 		}
// 	}
// 	return false
// }

// func isUniqueSecond(a []int) bool {
// 	var b []int

// 	for i := 0; i < len(a); i++ {
// 		if b[a[i]] == 1 {
// 			return true
// 		}
// 		b[a[i]] = 1
// 	}
// 	return false
// }

func isUniqueThird(a []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		if m[a[i]] > 0 {
			return true
		} else {
			m[a[i]]++
		}
	}
	return false
}
