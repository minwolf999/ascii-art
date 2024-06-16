package art

import (
	"ascii-art/align"
	"ascii-art/color"
	"ascii-art/structure"
	"fmt"
	"os"
	"strings"
)

func Ascii_Art(flag structure.Flag) {
	//We retrieve the data from the standard.txt document and store it temporarily in a string
	data, err := os.ReadFile("txtfiles/" + flag.Font + ".txt")
	if err != nil {
		return
	}
	texte := string(data)

	//We recreate the .txt document in a tab table
	var tab []string
	var temp string
	for _, element := range texte {
		if element == '\n' {
			if temp != "\n" {
				tab = append(tab, temp)
				temp = ""
			}
		} else {
			temp += string(element)
		}
	}
	if temp != "\n" {
		tab = append(tab, temp)
	}

	//We separate the string from the argument every \\n in an array
	input := strings.Split(flag.Word, "\\n")
	var fstring string
	var spacetab []int
	letters := 0

	//We do as many loops as there are in the table
	for _, table := range input {
		//If the current table entry is empty, we make a line break
		if table != "" {
			//We do a double loop which will add the letters in the correct order to the string variable.
			for i := 1; i < 9; i++ {
				//We loop as many times as there are letters in the current entry in the table
				for x, char := range table {
					/*We virtually delete the first 32 ascii then add the correct line from the clone table of
					the .txt file in the variable containing the new sentence to display*/
					ia := int(char) - 32

					for v, i := range flag.Color {
						if i == "" {
							i = flag.Word
						}

						if len(table[x:]) >= len(i[letters:]) && VerifString(table[x:], i[letters:]) {
							if letters == len(i)-1 {
								letters = 0
							} else {
								letters++
							}
							fstring += color.Color(v)
						}
					}

					fstring += tab[(ia*9+i)] + color.Color("reset")

					//if we have a space we add its position (equal to the length of fstring)
					if char == ' ' {
						spacetab = append(spacetab, len(fstring))
					}
				}
				//At the end of each line we do a newline
				fstring += "\n"

				if flag.Output != "" {
					fmt.Print(color.RemoveAnsii(fstring))
				} else {
					//we write following the alignment ask
					if flag.Justify == "left" || flag.Justify == "" {
						align.Left(fstring)
					} else if flag.Justify == "center" {
						align.Center(fstring)
					} else if flag.Justify == "right" {
						align.Right(fstring)
					} else if flag.Justify == "justify" {
						align.Justify(fstring, spacetab)
					}
					spacetab = []int{}
				}

				fstring = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
}

func VerifString(s, tab string) bool {
	if len(tab) <= len(s) && tab == s[:len(tab)] {
		return true
	}
	return false
}
