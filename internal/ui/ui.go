package ui

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	Cyan    = color.New(color.FgCyan, color.Bold)
	Green   = color.New(color.FgGreen, color.Bold)
	Red     = color.New(color.FgRed, color.Bold)
	Yellow  = color.New(color.FgYellow, color.Bold)
	Magenta = color.New(color.FgMagenta, color.Bold)
)

func PrintHeader(title string) {
	fmt.Println()
	Green.Println(title)
	fmt.Println()
}

func PrintError(err error) {
	Red.Printf("❌ Error: %v\n", err)
}

func PrintSuccess(msg string) {
	Green.Printf("✅ %s\n", msg)
}

func PrintInfo(msg string) {
	Cyan.Printf("ℹ️  %s\n", msg)
}

func PrintWarning(msg string) {
	Yellow.Printf("⚠️  %s\n", msg)
}

func PrintBox(msg string) {
	width := len(msg) + 4
	top := "╔" + strings.Repeat("═", width) + "╗"
	middle := "║  " + msg + "  ║"
	bottom := "╚" + strings.Repeat("═", width) + "╝"

	Magenta.Println(top)
	Magenta.Println(middle)
	Magenta.Println(bottom)
}
