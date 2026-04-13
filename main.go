package main

import (
	"fmt"
	"os"
	"strconv"

	tea "charm.land/bubbletea/v2"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: hya-7 [my-age [other-age]]

A TUI for exploring the "Half Your Age Plus 7" dating age rule.

Arguments:
  my-age      Your age (14-120)
  other-age   Optional partner age to check (14-120)

Controls:
  ←→/hljk     Adjust age by 1
  Shift+←→    Adjust age by 5
  Tab          Switch between your age and partner age
  Backspace    Clear partner age
  q/Ctrl+C     Quit

Free software released under the MIT License.
https://github.com/ConAcademy/hya-7
`)
}

func main() {
	m := initialModel()

	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" || arg == "-help" {
			usage()
			os.Exit(0)
		}
	}

	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Fprintf(os.Stderr, "Error: expected at most 2 arguments, got %d\n", len(args))
		usage()
		os.Exit(1)
	}

	if len(args) >= 1 {
		age, err := strconv.Atoi(args[0])
		if err != nil || age < 14 || age > 120 {
			fmt.Fprintf(os.Stderr, "Error: my-age must be a number between 14 and 120\n")
			os.Exit(1)
		}
		m.yourAge = age
	}

	if len(args) >= 2 {
		age, err := strconv.Atoi(args[1])
		if err != nil || age < 14 || age > 120 {
			fmt.Fprintf(os.Stderr, "Error: other-age must be a number between 14 and 120\n")
			os.Exit(1)
		}
		m.partnerAge = age
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
