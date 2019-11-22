package main

import "fmt"

func main() {
	fmt.Println(isShuffle("abc", "def", "dabecf"))
	fmt.Println(isShuffle("bac", "def", "dabecf"))
	fmt.Println(isShuffle("otto", "anna", ""))
	fmt.Println(isShuffle("otto", "anna", "oatntnoa"))
}

func isShuffle(str1, str2, str3 string) bool {
	var ptr1 int
	var ptr2 int

	if len(str1)+len(str2) != len(str3) {
		return false
	}

	for _, char := range str3 {
		if ptr1 < len(str1) && char == rune(str1[ptr1]) {
			ptr1++
			continue
		}
		if ptr2 < len(str2) && char == rune(str2[ptr2]) {
			ptr2++
			continue
		}
		return false
	}

	return true
}
