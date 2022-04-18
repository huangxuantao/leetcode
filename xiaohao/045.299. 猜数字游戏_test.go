package xiaohao

import (
	"fmt"
	"testing"
)

func getHint(secret string, guess string) string {
	var a, b int
	aIndexMap := make(map[int]bool)
	guessMap := make(map[int32]int, 10)
	for i, s := range secret {
		if secret[i] == guess[i] {
			a++
			aIndexMap[i] = true
		} else {
			guessMap[s]++
		}
	}

	for i, s := range guess {
		if aIndexMap[i] {
			continue
		}

		if guessMap[s] > 0 {
			b++
			guessMap[s]--
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}

func Test_getHint(t *testing.T) {
	t.Log(getHint("1807", "7810"))
	t.Log(getHint("1123", "0111"))
	t.Log(getHint("11", "10"))
}
