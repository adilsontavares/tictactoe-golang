package log

import "github.com/fatih/color"

func Message(format string, a ...interface{}) {
	color.White(format, a...)
}

func Warning(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

func Error(format string, a ...interface{}) {
	color.Red(format, a...)
}