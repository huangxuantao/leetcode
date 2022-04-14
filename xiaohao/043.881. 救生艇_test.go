package xiaohao

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	left := 0
	right := len(people) - 1

	var s int
	for {
		if people[left]+people[right] > limit {
			right -= 1
			s++
		} else {
			left++
			right--
			s++
		}

		if right-left == 0 {
			s++
			break
		}

		if right < left {
			break
		}
	}

	return s
}
