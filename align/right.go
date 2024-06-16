package align

import (
	"ascii-art/color"
	"fmt"
)

// Center the string to the width of the terminal.
// When the width is unknown, the string is left-aligned.
func Right(s string) {

	w := GetWidth()
	w2 := len(color.RemoveAnsii(s))
	var res string
	for i := 0; i <= (w - w2); i++ {
		res += " "
	}
	res += s
	fmt.Print(res)
}
