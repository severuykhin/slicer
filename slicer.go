package slicer

func MergeUnique[T any](slice1 []T, slice2 []T) []T {
	seen := make(map[any]bool)
	result := []T{}

	for _, num := range slice1 {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	for _, num := range slice2 {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
