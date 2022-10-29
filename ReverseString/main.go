package main

import (
	"fmt"
)

func main() {
	ModifyString()
}

func ModifyString() {
	str := "12abc123ieuwli"
	out := ""
	for _, val := range str {
		if val >= 48 && val <= 57 {
			continue
		}
		out = string(val) + out
	}
	fmt.Println(out)
}
