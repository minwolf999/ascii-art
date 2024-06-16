package help

import "fmt"

// write the utilisation message
func Help() {
	fmt.Println("Usage: go run . [FLAG] [STRING] [BANNER]")
	fmt.Println("EX: go run . something standard")
	fmt.Println()
	fmt.Println("The avalable BANNER are: standard, shadow, thinkertoy")
	fmt.Println("The avalable flags are:")
	fmt.Println("-h, --help to see this menu")
	fmt.Println("--justify=[JUSTIFY] to justify the result in the terminal (incompatible with the output flag). JUSTIFY can only be replace by \"left\", \"center\", \"right\", \"justify\"")
	fmt.Println("--output=[NAME_RESULT_FILE] to write the result in a file and not in the terminal (incompatible with all others flags). The NAME_RESULT_FILE need to be a standard text file.")
	fmt.Println("--color=[COLOR_NAME] [SUITE_OF_LETTERS] to color the SUITE_OF_LETTERS (or the complete word if ther isn't SUITE_OF_LETTERS) with the color (work with rgb, hsl, hexa)")
}