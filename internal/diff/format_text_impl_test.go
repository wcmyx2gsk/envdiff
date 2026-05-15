package diff

import (
	"strings"
	"testing"
)

// internal whitebox tests for FormatText helper logic

func TestFormatText_AllSections(t *testing.T) {
	result := Result{
		MissingInLeft:  []string{"ONLY_IN_RIGHT"},
		MissingInRight: []string{"ONLY_IN_LEFT"},
		Mismatched: []Mismatch{
			{Key: "PORT", LeftVal: "8080", RightVal: "9090"},
		},
	}
	out := FormatText(result, "left.env", "right.env")

	sections := []string{
		"Missing in right.env",
		"Missing in left.env",
		"Mismatched values",
		"ONLY_IN_RIGHT",
		"ONLY_IN_LEFT",
		"PORT",
		"8080",
		"9090",
	}
	for _, s := range sections {
		if !strings.Contains(out, s) {
			t.Errorf("expected %q in output:\n%s", s, out)
		}
	}
}

func TestFormatText_SummaryLine(t *testing.T) {
	result := Result{
		MissingInLeft:  []string{"A"},
		MissingInRight: []string{"B", "C"},
		Mismatched: []Mismatch{
			{Key: "D", LeftVal: "x", RightVal: "y"},
		},
	}
	out := FormatText(result, "l.env", "r.env")
	if !strings.Contains(out, "1 missing in left") {
		t.Errorf("expected summary with missing-in-left count, got:\n%s", out)
	}
	if !strings.Contains(out, "2 missing in right") {
		t.Errorf("expected summary with missing-in-right count, got:\n%s", out)
	}
	if !strings.Contains(out, "1 mismatched") {
		t.Errorf("expected summary with mismatched count, got:\n%s", out)
	}
}
