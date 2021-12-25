package bool

import (
	"strings"
)

func Sp(pattern string) []string {
	l := len(pattern)
	s := "()&|"
	var ans []string
	var k string = ""
	for i := 0; i < l; i++ {
		tmp := string(pattern[i])
		if strings.Contains(s, tmp) {
			if k != "" {
				ans = append(ans, k)
				k = ""
			}
			ans = append(ans, tmp)
		} else {
			k += tmp
		}

	}
	if k != "" {
		ans = append(ans, k)
	}
	return ans
}

//func To_intarr(pattern []string) [][]int {
//	s := "()&|"
//	for _, v := range pattern {
//		if !strings.Contains(s,v) {
//
//		}
//	}
//}
func Search(pattern []string, lists [][]int) []int {
	if len(lists) != 1 {
		if
		for i, v := range pattern {
			if v == ")" {
				if pattern[i-1] == "&" {

				}
			}
		}

	}

}

func union(slice1, slice2 []int) []int {
	m := make(map[int]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}
func intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}
