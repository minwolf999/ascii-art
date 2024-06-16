package getflags

import (
	"ascii-art/structure"
	"strings"
)

// Split all the arguments in the flag structure
func GetFlags(args []string) structure.Flag {
	var flag structure.Flag
	flag.Font = "standard"
	flag.Color = make(map[string]string)

	if len(args) == 0 {
		flag.Help = true
	}

	for i, v := range args {
		if v == "-h" || v == "--help" {
			flag.Help = true
			continue

		} else if strings.HasPrefix(v, "--color=") {
			if strings.HasPrefix(args[i+1], "--color=") || strings.HasPrefix(args[i+1], "--output=") || strings.HasPrefix(args[i+1], "--justify=") || args[i+1] == "-h" || args[i+1] == "--help" {
				flag.Color[strings.ReplaceAll(v, "--color=", "")] = ""
			} else {
				flag.Color[strings.ReplaceAll(v, "--color=", "")] = args[i+1]
			}

			continue

		} else if strings.HasPrefix(v, "--output=") {
			flag.Output = strings.ReplaceAll(v, "--output=", "")
			continue

		} else if strings.HasPrefix(v, "--justify=") {
			if strings.ReplaceAll(v, "--justify=", "") != "left" && strings.ReplaceAll(v, "--justify=", "") != "right" && strings.ReplaceAll(v, "--justify=", "") != "center" && strings.ReplaceAll(v, "--justify=", "") != "justify" {
				flag.Help = true
			} else {
				flag.Justify = strings.ReplaceAll(v, "--justify=", "")
			}
			continue

		} else if i == len(args)-1 && (v == "standard" || v == "shadow" || v == "thinkertoy") {
			flag.Font = v
			continue

		} else if i == len(args)-1 || i == len(args)-2 && (args[len(args)-1] == "standard" || args[len(args)-1] == "shadow" || args[len(args)-1] == "thinkertoy") && !strings.HasPrefix(v, "--color=") && !strings.HasPrefix(v, "--justify=") && !strings.HasPrefix(v, "--output=") && v != "-h" && v != "--help" {
			if flag.Word != "" {
				flag.Help = true
			}

			flag.Word = v
			continue
		}
	}

	if flag.Word == "" || flag.Justify != "" && flag.Output != "" || len(flag.Color) != 0 && flag.Output != "" {
		flag.Help = true
	}

	return flag
}
