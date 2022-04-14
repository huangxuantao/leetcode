package xiaohao

import (
	"math"
	"testing"
)

func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	ledMap := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		ledMap[i] = true
	}

	for i := 1; i < n; i++ {
		k := i
		for {
			if k >= n {
				break
			}

			ledMap[k] = !ledMap[k]
			k += i+1
		}
	}

	var c int
	for _, v := range ledMap {
		if v==true {
			c++
		}
	}
	return c
}

func bulbSwitch2(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}

func Test_bulbSwitch(t *testing.T) {
	t.Log(bulbSwitch(1))
	t.Log(bulbSwitch(2))
	t.Log(bulbSwitch(3))
	t.Log(bulbSwitch(4))
	t.Log(bulbSwitch(5))
	t.Log(bulbSwitch(6))
	t.Log(bulbSwitch(7))
	t.Log(bulbSwitch(8))
	t.Log(bulbSwitch(9))
	t.Log(bulbSwitch(10))
	t.Log(bulbSwitch(11))
	t.Log(bulbSwitch(12))
	//t.Log(bulbSwitch(9999999))
}