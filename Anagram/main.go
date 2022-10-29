package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abced.", "abce.d"))
}

func isAnagram(str1, str2 string) bool {
	isAnagram := false
	if len(str1) != len(str2) {
		return isAnagram
	}
	mp := make(map[string]int)
	for _, key := range str1 {
		_, ok := mp[string(key)]
		if !ok {
			mp[string(key)] = 1
			continue
		}
		mp[string(key)] += 1
	}

	for _, key := range str2 {
		_, ok := mp[string(key)]
		if !ok {
			mp[string(key)] = 1
			continue
		}
		mp[string(key)] -= 1
	}

	for _, value := range mp {
		if value != 0 {
			return false
		}
	}
	return true
}
