package color

import "runtime"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func initColors() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func RedString(str string) string {
	return Red + str + Reset
}

func GreenString(str string) string {
	return Green + str + Reset
}

func YellowString(str string) string {
	return Yellow + str + Reset
}

func BlueString(str string) string {
	return Blue + str + Reset
}

func PurpleString(str string) string {
	return Purple + str + Reset
}

func CyanString(str string) string {
	return Cyan + str + Reset
}

func GrayString(str string) string {
	return Gray + str + Reset
}

func WhiteString(str string) string {
	return White + str + Reset
}
