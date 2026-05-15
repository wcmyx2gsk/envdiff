package diff

import (
	"fmt"
	"io"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

// FormatText writes a human-readable, plain-text summary of the diff result
// to w. If color is true, ANSI escape codes are used to highlight differences.
func FormatText(w io.Writer, r Result, leftName, rightName string, color bool) {
	if r.IsClean() {
		fmt.Fprintln(w, "No differences found.")
		return
	}

	if len(r.MissingInRight) > 0 {
		fmt.Fprintf(w, "Missing in %s:\n", rightName)
		for _, k := range r.MissingInRight {
			if color {
				fmt.Fprintf(w, "  %s- %s%s\n", colorRed, k, colorReset)
			} else {
				fmt.Fprintf(w, "  - %s\n", k)
			}
		}
	}

	if len(r.MissingInLeft) > 0 {
		fmt.Fprintf(w, "Missing in %s:\n", leftName)
		for _, k := range r.MissingInLeft {
			if color {
				fmt.Fprintf(w, "  %s+ %s%s\n", colorGreen, k, colorReset)
			} else {
				fmt.Fprintf(w, "  + %s\n", k)
			}
		}
	}

	if len(r.Mismatched) > 0 {
		fmt.Fprintln(w, "Mismatched values:")
		for _, m := range r.Mismatched {
			line := fmt.Sprintf("  ~ %s  (%s: %q  |  %s: %q)",
				m.Key, leftName, m.LeftValue, rightName, m.RightValue)
			if color {
				fmt.Fprintf(w, "%s%s%s\n", colorYellow, line, colorReset)
			} else {
				fmt.Fprintln(w, line)
			}
		}
	}

	_ = strings.NewReplacer() // ensure strings import is used via package-level use
}
