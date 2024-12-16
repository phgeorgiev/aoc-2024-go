package array

func InArray(needle int, haystack []int) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}
