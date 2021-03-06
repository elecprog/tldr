// Copyright © 2020 Evert Provoost
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package pages

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	normal      = colDefault
	heading     = normal | modBold
	note        = normal
	description = normal
	verbatim    = colBrightRed
	example     = normal | modItalic
)

func emptyDatabase() {
	// The database is empty
	fmt.Fprint(os.Stderr, "\n  ", "The database is empty.")
	fmt.Fprint(os.Stderr, "\n  ", "You can try updating the database using ", colorize("tldr --update", verbatim), ".\n\n")
}

func pageUnavailable(command string) {
	// The page is not in the database
	fmt.Print("\n  ", colorize(command, heading), " documentation is not available.")
	fmt.Print("\n  ", "You can try updating the database using ", colorize("tldr --update", verbatim), ".")
	fmt.Print("\n  ", "Or add a page yourself to https://github.com/tldr-pages/tldr.", "\n\n")
}

func prettyPrint(page []byte) {
	// Don't pretty print to TTY
	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Print(string(page))
		return
	}

	// Add an blank line in front of the page
	fmt.Println()

	// Pretty print the lines in the page
	for _, lineB := range bytes.Split(page, []byte{'\n'}) {
		line := string(bytes.TrimSpace(lineB))

		if len(line) == 0 {
			// Skip empty lines
			continue
		}

		switch line[0] {
		case '#':
			fmt.Print("  ")
			processLine(line[1:], heading)

		case '>':
			fmt.Print("  ")
			processLine(line[1:], note)

		case '-':
			fmt.Print("\n- ")
			processLine(line[1:], description)

		default:
			fmt.Print("  ")
			processLine(line, normal)
		}
	}

	// Add an extra blank line at the end of the page
	fmt.Println()
}

func processLine(line string, defaultStyle color) {
	// Remove unneeded spaces
	line = strings.TrimSpace(line)

	// Our parsing method would fail on ``, but as
	// these a no-ops we can safely remove them.
	line = strings.Replace(line, "``", "", -1)

	inVerbatim := line[0] == '`'
	for _, part := range strings.Split(line, "`") {
		// Skip empty strings
		if len(part) == 0 {
			continue
		}

		if inVerbatim {
			// Verbatim
			processVerbatim(part)

		} else {
			// Normal text
			fmt.Print(colorize(part, defaultStyle))
		}

		inVerbatim = !inVerbatim
	}

	// As you might have noticed, we never check if the backticks are balanced.
	// But that check is not regular, and pages should be valid,
	// so in theory we never have a case where the backticks aren't balanced.

	// Go to the next line
	fmt.Println()
}

func processVerbatim(line string) {
	// Our parsing method would fail on {{}} or }}{{, but as
	// these a no-ops we can safely remove them.
	line = strings.Replace(line, "{{}}", "", -1)
	line = strings.Replace(line, "}}{{", "", -1)

	inExample := strings.HasPrefix(line, "{{")
	for _, segment := range strings.Split(line, "{{") {
		for _, part := range strings.Split(segment, "}}") {
			if inExample {
				// Optional
				fmt.Print(colorize(part, example))

			} else {
				// Verbatim
				fmt.Print(colorize(part, verbatim))
			}

			inExample = !inExample
		}
	}

	// As you might have noticed, we never check if the braces are balanced.
	// But that check is not regular, and pages should be valid,
	// so in theory we never have a case where the braces aren't balanced.
}
