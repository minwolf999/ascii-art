package color

import "math"

func HSLtoRGB(hsl HSL) RGB {
	//the value s and l being % we divide it by 100 and store the result
	s := float64(hsl.S) / 100
	l := float64(hsl.L) / 100

	//we multiply by s the absyss -1 of 2*hsl.L-1
	c := (1 - math.Abs(2*l-1)) * s
	//we multiply c by the absyss -1 of the modulo of hsl.h/60 by 2 -1
	x := c * (1 - math.Abs(math.Mod(float64(hsl.H)/60, 2)-1))
	//we remove half of c at hsl.L/100
	m := l - c/2

	var r, g, b float64

	switch {
	//if the hue is less than 60: r=c, g=x, b=0
	case hsl.H < 60:
		r = c
		g = x
		b = 0
	//if the hue is between 60 and 120: r=x,g=c, b=0
	case hsl.H < 120 && hsl.H >= 60:
		r = x
		g = c
		b = 0
	//if the hue is between 120 and 180: r=0, g=c, b=x
	case hsl.H < 180 && hsl.H >= 120:
		r = 0
		g = c
		b = x
	//if the hue is between 180 and 240: r=0, g=c, b=c
	case hsl.H < 240 && hsl.H >= 180:
		r = 0
		g = x
		b = c
	//if the hue is between 240 and 300: r=x, g=0, b=c
	case hsl.H < 300 && hsl.H >= 240:
		r = x
		g = 0
		b = c
	//if the hue is between 300 and 360: r=c, g=0, b=x
	case hsl.H < 360 && hsl.H >= 300:
		r = c
		g = 0
		b = x
	}

	//we round to int after adding it to m and multiplying by 255
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))

	return RGB{R, G, B}
}
