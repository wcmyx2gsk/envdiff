package diff

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONReport is the structured representation of a diff result for JSON output.
type JSONReport struct {
	Clean      bool               `json:"clean"`
	MissingIn  map[string][]string `json:"missing_in,omitempty"`
	Mismatched []MismatchedJSON   `json:"mismatched,omitempty"`
}

// MismatchedJSON holds a single key whose values differ between two env files.
type MismatchedJSON struct {
	Key    string `json:"key"`
	Values map[string]string `json:"values"`
}

// FormatJSON writes the diff result as a JSON document to w.
// The keys in MissingIn are the file names provided in the Result.
func FormatJSON(w io.Writer, r Result, names []string) error {
	if len(names) != 2 {
		return fmt.Errorf("FormatJSON: expected exactly 2 names, got %d", len(names))
	}

	report := JSONReport{
		Clean:     r.IsClean(),
		MissingIn: make(map[string][]string),
	}

	if len(r.MissingInRight) > 0 {
		report.MissingIn[names[1]] = r.MissingInRight
	}
	if len(r.MissingInLeft) > 0 {
		report.MissingIn[names[0]] = r.MissingInLeft
	}
	if len(report.MissingIn) == 0 {
		report.MissingIn = nil
	}

	for _, m := range r.Mismatched {
		report.Mismatched = append(report.Mismatched, MismatchedJSON{
			Key: m.Key,
			Values: map[string]string{
				names[0]: m.LeftValue,
				names[1]: m.RightValue,
			},
		})
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(report)
}
