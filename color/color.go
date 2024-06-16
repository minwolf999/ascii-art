package color

import (
	"fmt"
	"strconv"
	"strings"
)

func Color(s string) string {
	//if the color indicated is RGB
	if s[0:3] == "rgb" {
		//we put the given RGB values in a []string
		t := strings.Split(s[3:], ", ")

		//convert strings to int
		r, err := strconv.Atoi(t[0][1:])
		if err != nil || r > 255 {
			panic("La valeur donné en R est invalide : " + t[0][1:])
		}

		g, err := strconv.Atoi(t[1])
		if err != nil || g > 255 {
			panic("La valeur donné en G est invalide : " + t[1])
		}

		b, err := strconv.Atoi(t[2][:len(t[2])-1])
		if err != nil || b > 255 {
			panic("La valeur donné en B est invalide : " + t[2][:len(t[2])-1])
		}

		//we return the RGB value which we translate into Ansi
		return RGBtoAnsi(RGB{r, g, b})

		//if the color indicated is HSL
	} else if s[0:3] == "hsl" {
		//we put the given HSL values in a []string
		t := strings.Split(s[3:], ", ")

		var h, s, l int
		var err error
		//we convert the string to int by removing the %

		h, err = strconv.Atoi(t[0][1:])
		if err != nil || h > 360 {
			panic("La valeur donné en H est invalide : " + t[0][1:])
		}

		if len(t[1]) == 1 {
			s, err = strconv.Atoi(t[1])
			if err != nil || s > 100 {
				panic("La valeur donné en S est invalide : " + t[1])
			}
		} else {
			s, err = strconv.Atoi(t[1][:len(t[1])-1])
			if err != nil || s > 100 {
				panic("La valeur donné en S est invalide : " + t[1][:len(t[1])-1])
			}
		}

		if (len(t[2])) == 2 {
			l, err = strconv.Atoi(t[2][:len(t[2])-1])
			if err != nil || l > 100 {
				panic("La valeur donné en L est invalide : " + t[2][:len(t)-1])
			}
		} else {
			l, err = strconv.Atoi(t[2][:len(t[2])-2])
			if err != nil || l > 100 {
				panic("tLa valeur donné en L est invalide : " + t[2][:len(t[2])-2])
			}
		}

		return RGBtoAnsi(HSLtoRGB(HSL{h, s, l}))

		//if the color indicated is an Hexadecimal
	} else if s[0] == '#' {
		//we convert the values 2 by 2 and store them in RGB
		r, err := strconv.ParseUint(s[1:3], 16, 32)
		if err != nil {
			panic("Le code donné " + s + " n'est pas un code Hexadécimal")
		}

		g, err := strconv.ParseUint(s[3:5], 16, 32)
		if err != nil {
			panic("Le code donné " + s + " n'est pas un code Hexadécimal")
		}

		b, err := strconv.ParseUint(s[5:7], 16, 32)
		if err != nil {
			panic("Le code donné " + s + " n'est pas un code Hexadécimal")
		}

		return RGBtoAnsi(RGB{R: int(r), G: int(g), B: int(b)})

	} else {
		//we make a map with string colors and their associated Ansi values
		tab := make(map[string]string)
		tab["reset"] = "\033[0m"
		tab["black"] = "\033[30m"
		tab["red"] = "\033[31m"
		tab["green"] = "\033[32m"
		tab["yellow"] = "\033[33m"
		tab["blue"] = "\033[34m"
		tab["magenta"] = "\033[35m"
		tab["cyan"] = "\033[36m"
		tab["light gray"] = "\033[37m"
		tab["dark gray"] = "\033[90m"
		tab["light red"] = "\033[91m"
		tab["light green"] = "\033[92m"
		tab["light yellow"] = "\033[93m"
		tab["light blue"] = "\033[94m"
		tab["light magenta"] = "\033[95m"
		tab["light cyan"] = "\033[96m"
		tab["gray"] = "\033[38;5;243m"
		tab["white"] = "\033[97m"
		tab["orange"] = "\033[38;2;255;165;0m"

		//if the color is not indicated, an error is returned
		_, test := tab[s]
		if !test {
			fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return ""
		}

		//we return the Ansi value
		return tab[s]
	}
}
