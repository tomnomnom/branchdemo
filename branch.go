package branch

// LessMore returns the number of ints less than, and more than the pivot
func LessMore(ints []int, pivot int) (int, int) {
	less, more := 0, 0

	for _, num := range ints {
		if num < pivot {
			less++
		} else if num > pivot {
			more++
		}
	}

	return less, more
}
