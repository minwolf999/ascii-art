package align

import (
	"ascii-art/color"
	"fmt"
)

func Justify(s string, tab []int) {
	w := GetWidth()
	w2 := len(color.RemoveAnsii(s))
	var res string
	y := 0

	if len(tab) == 0 {
		Left(s)
		return
	}

	for i := range s {
		if i == tab[y] {
			for x := 0; x <= (w-w2)/len(tab); x++ {
				if y != 0 && x == (w-w2)/len(tab) {
					continue
				}
				res += " "
			}

			if y < len(tab)-1 {
				y++
			}
		}

		res += string(s[i])
	}
	fmt.Print(res)
}
