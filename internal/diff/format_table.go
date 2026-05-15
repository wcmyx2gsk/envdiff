package diff

import (
	"fmt"
	"io"
	"strings"
)

// FormatTable writes a human-readable table representation of the diff result
// to the provided writer. Columns are padded for alignment.
func FormatTable(w io.Writer, result Result, leftName, rightName string) error {
	if leftName == "" {
		leftName = "left"
	}
	if rightName == "" {
		rightName = "right"
	}

	if result.Clean() {
		_, err := fmt.Fprintln(w, "No differences found.")
		return err
	}

	const colWidth = 30
	pad := func(s string) string {
		if len(s) >= colWidth {
			return s[:colWidth-1] + "…"
		}
		return s + strings.Repeat(" ", colWidth-len(s))
	}

	header := fmt.Sprintf("%-20s  %-6s  %s  %s", "KEY", "STATUS", pad(leftName), pad(rightName))
	sep := strings.Repeat("-", len(header))

	if _, err := fmt.Fprintln(w, header); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, sep); err != nil {
		return err
	}

	for _, k := range result.MissingInRight {
		line := fmt.Sprintf("%-20s  %-6s  %s  %s", k, "LEFT", pad("(present)"), pad("(missing)"))
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}

	for _, k := range result.MissingInLeft {
		line := fmt.Sprintf("%-20s  %-6s  %s  %s", k, "RIGHT", pad("(missing)"), pad("(present)"))
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}

	for _, m := range result.Mismatched {
		line := fmt.Sprintf("%-20s  %-6s  %s  %s", m.Key, "DIFF", pad(m.LeftValue), pad(m.RightValue))
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}

	return nil
}
