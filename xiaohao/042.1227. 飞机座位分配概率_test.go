package xiaohao

import "testing"

func nthPersonGetsNthSeat(n int) float64 {
	var s float64
	s = 1
	if n == 1 {
		return s
	}

	for i := 0; i < n; i++ {
		s = s * float64(1) / float64(n-i)
	}
	return s
}

func Test_nthPersonGetsNthSeat(t *testing.T) {
	t.Log(nthPersonGetsNthSeat(1))
	t.Log(nthPersonGetsNthSeat(2))
	t.Log(nthPersonGetsNthSeat(3))
}
