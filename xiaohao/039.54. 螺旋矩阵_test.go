package xiaohao

import "fmt"

func spiralOrder(matrix [][]int) []int {
	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	type P struct {
		X int
		Y int
	}
	p := P{0, 0}
	di := 1 // 方向 1右 2下 3左 4上

	var s []int

	for {
		if up > down || left > right {
			break
		}

		fmt.Println(di, p, up, down, left, right)
		s = append(s, matrix[p.X][p.Y])

		switch di {
		case 1:
			p.Y++
			if p.Y > right {
				up++
				p.Y--
				p.X++
				di = 2
			}
			break

		case 2:
			p.X++
			if p.X > down {
				right--
				p.X--
				p.Y--
				di = 3
			}
			break

		case 3:
			p.Y--
			if p.Y < left {
				down--
				p.Y++
				p.X--
				di = 4
			}
			break

		case 4:
			p.X--
			if p.X < up {
				left++
				p.X++
				p.Y++
				di = 1
			}
			break

		}
	}

	return s
}
