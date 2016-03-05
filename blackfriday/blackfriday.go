package blackfriday

func FindLargestUniqueNumber(numbers []int) (none bool, number int) {

	// make groups for all results
	for i := 6; i > 0; i-- {
		count := 0

		// distribute numbers in the correct group
		for _, n := range numbers {
			if i == n {
				count += 1
			}
		}

		if count == 1 {
			index := FindFirstIndex(numbers, i)
			return false, index + 1
		}
	}

	return true, 0
}

func FindFirstIndex(numbers []int, value int) int {
	for index, n := range numbers {
		if n == value {
			return index
		}
	}
	return -1
}
