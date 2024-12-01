package main

import (
	"advent-of-code/extractor"
	"log"
	"math"
	"sort"
	"strconv"
)

func orderAscending(strings []string) []int {
	ints := make([]int, len(strings))
	for i, str := range strings {
		num, _ := strconv.Atoi(str)
		ints[i] = num
	}

	sort.Ints(ints)
	return ints
}

func calculateColumnDifference(firstColumn []int, secondColumn []int) int {
	difference := 0
	for i := 0; i < len(firstColumn); i++ {
		difference += int(math.Abs(float64(firstColumn[i] - secondColumn[i])))
	}
	return difference
}

func main() {
	fileName := "input.txt"

	columns, _ := extractor.ExtractColumns(fileName)

	firstColumnAsc := orderAscending(columns[0])
	secondColumnAsc := orderAscending(columns[1])

	log.Println(calculateColumnDifference(firstColumnAsc, secondColumnAsc))
}
