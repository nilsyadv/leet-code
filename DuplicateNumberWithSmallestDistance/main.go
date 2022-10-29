package main

import "fmt"

type Distance struct {
	Dist  int64
	Start *int64
	End   *int64
}

var mp map[int64]Distance

func main() {
	test := []int64{0, -1, 3, 5, -1, 8}
	mp := make(map[int64]Distance)

	for index, number := range test {
		tst, ok := mp[number]
		if ok && tst.End != nil {
			diff := int64(index) - *tst.End
			if tst.Dist > diff {
				tst.Dist = diff
				tst.Start = tst.End
				temp := int64(index)
				tst.End = &temp
			}
		} else if tst.End == nil && tst.Start != nil && index != 0 {
			temp := int64(index)
			tst.End = &temp
			tst.Dist = *tst.End - *tst.Start
		} else {
			temp := int64(index)
			tst.Start = &temp
		}

		mp[number] = tst
		fmt.Printf("%+v\n", tst)
	}

	fmt.Printf("%v", mp)

	var smallDuplicatenum int64
	var smallDuplicatedist int64
	var found *bool
	for num, numDis := range mp {
		fmt.Println(num)
		if numDis.End != nil && (smallDuplicatedist > numDis.Dist || found == nil) {
			smallDuplicatenum = num
			smallDuplicatedist = numDis.Dist
			ts := true
			found = &ts
		}
	}

	if found != nil && *found {
		fmt.Println("Found Duplicate Number: ", smallDuplicatenum, " & Smallest distance: ",
			smallDuplicatedist)
	} else {
		fmt.Println("There is no Duplicate number")
	}
}
