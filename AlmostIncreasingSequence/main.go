package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

var mp map[int]int

func main() {
	mp := make(map[int]int)
	array := []int{0, 1, 3}
	var sortedarray []float64
	var ispossible *bool
	for _, element := range array {
		_, ok := mp[element]
		if ok {
			tmp := false
			ispossible = &tmp
			break
		}
		mp[element] = 1
	}

	if ispossible != nil {
		fmt.Println("AlmostIncreasingSequence:", false)
		os.Exit(0)
	}

	for key := range mp {
		sortedarray = append(sortedarray, math.Abs(float64(key)))
	}

	sort.Float64s(sortedarray)

	for idx := 0; idx < len(sortedarray)-1; idx++ {
		fmt.Println(sortedarray[idx+1] - sortedarray[idx])
		if math.Abs(sortedarray[idx+1]-sortedarray[idx]) > 1 {
			fmt.Println("AlmostIncreasingSequence:", false)
			os.Exit(0)
		}
	}

	fmt.Println("AlmostIncreasingSequence:", true)
}
