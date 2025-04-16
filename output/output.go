package output

import "github.com/fatih/color"

func PrintError(str string, err error) {
	color.Red("%s: %s", str, err.Error())
}

func PrintWarning(str string, err error) {
	color.Yellow("%s: %s", str, err.Error())
}

func PrintInfo(str string) {
	color.Green("%s: %s", str)
}
