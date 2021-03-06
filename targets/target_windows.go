// Copyright © 2019 Evert Provoost
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

package targets

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// OsName is the name of the current platform
const OsName = "Windows"

// OsDir is the directory in the tldr pages containing
// the pages for this platform
var OsDir = "windows"

// CurrentLanguage is the current users language
// TODO: Detect language on Windows
var CurrentLanguage = "en"

// Windows by default ignores ASCII escape codes,
// however we can change this using this.
// Why is this not the default? No idea...
func init() {
	// Get the current console settings
	var consoleMode uint32 = 0
	err := windows.GetConsoleMode(windows.Stdout, &consoleMode)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get current terminal mode, please open an issue, this should work...")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Add support for escape codes to those settings
	consoleMode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	err = windows.SetConsoleMode(windows.Stdout, consoleMode)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to enable ASCII escape sequences, please open an issue, this should work...")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
