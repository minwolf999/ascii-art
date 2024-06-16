package align

import (
	"ascii-art/color"
	"fmt"
)

func Center(s string) {
	w := GetWidth()
	w2 := len(color.RemoveAnsii(s))
	var res string
	for i := 0; i < (w-w2)/2; i++ {
		res += " "
	}
	res += s
	fmt.Print(res)
}
