package color

import (
	"strconv"
)

// \x1n[38;2;R;G;Bm] we therefore replace R;G;B with the RGB values
func RGBtoAnsi(rgb RGB) string {
	str := "\x1b[38;2;" + strconv.Itoa(rgb.R) + ";" + strconv.Itoa(rgb.G) + ";" + strconv.Itoa(rgb.B) + "m"

	return str
}
