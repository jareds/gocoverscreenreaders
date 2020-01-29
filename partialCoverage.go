package gocoverscreenreaders

func partialCoverage(num int) int {
	if num == 1 {
		return 1
	} else if num == 2 {
		return 2
	} else if num == 3 {
		return 3
	} else if num == 4 {
		return 4
	} else if num == 5 {
		return 5
	} else if num == 6 {
		return 6
	} else if num == 33 {
		returnValue := 0
		for i := 0; i < 5; i++ {
			returnValue += i
		}
		return returnValue
	} else {
		return -1
	}
}
