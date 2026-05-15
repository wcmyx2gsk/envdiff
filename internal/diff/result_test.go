package diff

import (
	"strings"
	"testing"
)

func TestResult_Clean_True(t *testing.T) {
	r := Result{}
	if !r.Clean() {
		t.Error("expected Clean() to return true for empty result")
	}
}

func TestResult_Clean_False(t *testing.T) {
	cases := []Result{
		{MissingInRight: []string{"KEY"}},
		{MissingInLeft: []string{"KEY"}},
		{Mismatched: []MismatchedKey{{Key: "K", LeftValue: "a", RightValue: "b"}}},
	}
	for _, r := range cases {
		if r.Clean() {
			t.Errorf("expected Clean() false for result: %+v", r)
		}
	}
}

func TestResult_Summary_Clean(t *testing.T) {
	r := Result{}
	if r.Summary() != "files are identical" {
		t.Errorf("unexpected summary: %s", r.Summary())
	}
}

func TestResult_Summary_Single(t *testing.T) {
	r := Result{MissingInRight: []string{"KEY"}}
	if r.Summary() != "1 difference found" {
		t.Errorf("unexpected summary: %s", r.Summary())
	}
}

func TestResult_Summary_Multiple(t *testing.T) {
	r := Result{
		MissingInRight: []string{"A", "B"},
		Mismatched:     []MismatchedKey{{Key: "C"}},
	}
	s := r.Summary()
	if !strings.Contains(s, "3") {
		t.Errorf("expected '3' in summary, got: %s", s)
	}
}
