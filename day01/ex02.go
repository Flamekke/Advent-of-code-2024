package main

import (
	"advent-of-code/extractor"
	"fmt"
)

func main() {
	arr0 := []int{}
	dict1 := make(map[int]int)
	extractor.StreamFile("input.txt", func(line string) {
		var value0, value1 int
		fmt.Sscanf(line, "%d   %d", &value0, &value1)
		arr0 = append(arr0, value0)
		dict1[value1] = dict1[value1] + 1
	})
	result := 0
	for _, val0 := range arr0 {
		result += dict1[val0] * val0
	}
	fmt.Println(result)
}
