package gostl

import (
	"fmt"
	"os"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// Error prints an error message and exits the program
func Error(s string) {
	fmt.Println(Red + "[ERROR] " + Reset + s)
	os.Exit(1)
}

// Info prints an info message
func Info(s string) {
	fmt.Println(Cyan + "[INFO] " + Reset + s)
}

// Warn prints a warning message
func Warn(s string) {
	fmt.Println(Yellow + "[WARN] " + Reset + s)
	os.Exit(1)
}
