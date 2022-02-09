package leetcode

import (
	"strings"
	"testing"
)

func convert(s string, numRows int) string {
	type P struct {
		X int
		Y int
	}

	if numRows == 1 {
		return s
	}
	n := numRows + numRows - 2
	var x, y []int

	for i := 0; i < numRows; i++ {
		x = append(x, 0)
	}
	for i := 0; i < numRows-2; i++ {
		x = append(x, i+1)
	}

	for i := 0; i < numRows; i++ {
		y = append(y, i)
	}
	for i := numRows - 2; i > 0; i-- {
		y = append(y, i)
	}

	//fmt.Println(x)
	//fmt.Println(y)

	var maxX, maxY int
	xx := make(map[P]string)

	for i, i2 := range s {
		index := i % n
		//fmt.Println(i, (numRows-1)*(i/n)+x[index], y[index], string(i2))

		p := P{
			X: (numRows-1)*(i/n) + x[index],
			Y: y[index],
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}

		xx[p] = string(i2)
	}

	var build strings.Builder

	for i := 0; i < maxY+1; i++ {
		for j := 0; j < maxX+1; j++ {
			p := P{X: j, Y: i}
			if v, has := xx[p]; has {
				build.WriteString(v)
			}
		}
	}

	return build.String()
}

func Test_convert(t *testing.T) {
	t.Log(convert("enrxihcsanbtgxdcttnujvfscrwqtyuynmxwvbqxorquowzhpmdzjlrlcncyoywbmvzhxpenhvivthfivkhfxbqaquyetwifthnsxrggoqbhxiqsvzzscqutmszfqjnmtaeqtmykcbrzkjuhltznluiyokfhvstouzgqmeaogrqsdmzohydtuotjyysttlknmqdyvdpbxftatuqastvphoahpsdifnxrfbqaghdfoyhhsxhntdcctcmiupqzeqsjrkmzrphfoooioyvjdxnwbzhvqzwuprgibsitjpazfritpfesfsqgrxekxcgmtmvvgfqqwspdjxzaddukvlqpkuzjzhwsutpcafsyaibjhammegwbtpqelrnkbldxguzgcseccinlizyogwqzlifxcthdgmanjztltanviajschhkdxlrfrohmxmsmmhvodihdvfnxofbxmlclxvrojacjpwxbeurhcbmzgzwwgyvtlzeivxygaappzosdikkmlwltxirdioytnfeieepehgvgvqjfavsntfiqnbvxputczznfdcmkkhshxdnzyhormwjcxfbobwrcvehbitnsdgacjpeiysbmrnoqssfvoyxkeglmaygfgihqznazgdmzqcpiuudjucvyjimlivqpdzhfnhevksudvjlrgrcavxzehlrqgjhmjqtyzztjsnopyagetjfqiexqroiayrojhjhgiarcpgrniysdhztpfqhwhpyfioopxxvgxniovabdatgjszazsiwzzweiluxirvqqkzefbhiddqmxrmxjwtiwrogckdldadtkczpfhzikpujhjgqfbbbt", 373))
}

func Benchmark_convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert("enrxihcsanbtgxdcttnujvfscrwqtyuynmxwvbqxorquowzhpmdzjlrlcncyoywbmvzhxpenhvivthfivkhfxbqaquyetwifthnsxrggoqbhxiqsvzzscqutmszfqjnmtaeqtmykcbrzkjuhltznluiyokfhvstouzgqmeaogrqsdmzohydtuotjyysttlknmqdyvdpbxftatuqastvphoahpsdifnxrfbqaghdfoyhhsxhntdcctcmiupqzeqsjrkmzrphfoooioyvjdxnwbzhvqzwuprgibsitjpazfritpfesfsqgrxekxcgmtmvvgfqqwspdjxzaddukvlqpkuzjzhwsutpcafsyaibjhammegwbtpqelrnkbldxguzgcseccinlizyogwqzlifxcthdgmanjztltanviajschhkdxlrfrohmxmsmmhvodihdvfnxofbxmlclxvrojacjpwxbeurhcbmzgzwwgyvtlzeivxygaappzosdikkmlwltxirdioytnfeieepehgvgvqjfavsntfiqnbvxputczznfdcmkkhshxdnzyhormwjcxfbobwrcvehbitnsdgacjpeiysbmrnoqssfvoyxkeglmaygfgihqznazgdmzqcpiuudjucvyjimlivqpdzhfnhevksudvjlrgrcavxzehlrqgjhmjqtyzztjsnopyagetjfqiexqroiayrojhjhgiarcpgrniysdhztpfqhwhpyfioopxxvgxniovabdatgjszazsiwzzweiluxirvqqkzefbhiddqmxrmxjwtiwrogckdldadtkczpfhzikpujhjgqfbbbt", 373)
	}
}
