package main

import (
	"ascii-art/art"
	getflags "ascii-art/getFlags"
	"ascii-art/help"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	flag := getflags.GetFlags(args)

	if flag.Help {
		help.Help()
		return
	}

	stdout := os.Stdout
	if flag.Output != "" {
		file, err := os.Create(flag.Output)
		if err != nil {
			fmt.Println(err)
			return
		}

		os.Stdout = file
	}

	art.Ascii_Art(flag)

	os.Stdout = stdout
}
