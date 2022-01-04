package arrayutil

func IsContains(first, second []int, minimumApperance int) bool {
	exists := make(map[int]struct{})
	for _, value := range second {
		exists[value] = struct{}{}
	}
	var counter int
	for _, value := range first {
		if _, ok := exists[value]; ok {
			counter++
			if counter >= minimumApperance {
				return true
			}
		}
	}
	return false
}

func IsAny(first, second []int) []int {
	var any []int
	exists := make(map[int]struct{})
	for _, value := range second {
		exists[value] = struct{}{}
	}
	for _, value := range first {
		if _, ok := exists[value]; ok {
			any = append(any, value)
		}
	}
	return any
}
