package main

import (
	"fmt"
	"kth/blackfriday"
)

func main() {
	var err error
	var size int

	_, err = fmt.Scanf("%d", &size)
	if err != nil {
		panic(err)
	}

	var n int
	results := make([]int, size)
	for i := 0; i < size; i++ {
		_, err = fmt.Scanf("%d", &n)
		if err != nil {
			panic(err)
		}
		results[i] = n
	}

	none, solution := blackfriday.FindLargestUniqueNumber(results)

	if none {
		fmt.Println("none")
	} else {
		fmt.Printf("%d\n", solution)
	}

}

