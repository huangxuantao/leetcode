package xiaohao

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	min := 0
	max := n

	var errL = -1
	var err int
	for {
		errL = err
		err = (min + max) / 2
		if errL == err {
			return err + 1
		}

		if isBadVersion(err) {
			max = err
		} else {
			min = err
		}
	}
}
