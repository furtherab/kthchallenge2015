package main

import (
	"fmt"
	"kth/proteins"
)

func main() {
	var n int
	var s string

	_, err := fmt.Scanf("%d\n%s", &n, &s)
	if err != nil {
		panic(err)
	}

	solution := proteins.FindLeastInsertions(n, s);
	fmt.Printf("%d\n", solution)
}
