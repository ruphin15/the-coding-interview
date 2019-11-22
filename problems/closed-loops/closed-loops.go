package main

import "fmt"

func main() {
	fmt.Println(getTotalLoops("2876"))
}

func getTotalLoops(number string) int {
	loops := map[rune]int{
		'0': 1,
		'6': 1,
		'8': 2,
		'9': 1,
	}

	var totalLoops int
	for _, char := range number {
		if l, ok := loops[char]; ok {
			totalLoops += l
		}
	}
	return totalLoops
}
