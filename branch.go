package branchdemo

// CountInt returns the number of instances of search in ints
func CountInt(search int, ints []int) int {
	count := 0

	for _, num := range ints {
		if num == search {
			count++
		}
	}

	return count
}
