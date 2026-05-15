package diff

import (
	"strings"
	"testing"
)

func TestFormatTable_Clean(t *testing.T) {
	result := Result{}
	var sb strings.Builder
	if err := FormatTable(&sb, result, "dev", "prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(sb.String(), "No differences found.") {
		t.Errorf("expected clean message, got: %s", sb.String())
	}
}

func TestFormatTable_MissingInRight(t *testing.T) {
	result := Result{
		MissingInRight: []string{"SECRET_KEY"},
	}
	var sb strings.Builder
	if err := FormatTable(&sb, result, "dev", "prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := sb.String()
	if !strings.Contains(out, "SECRET_KEY") {
		t.Errorf("expected key in output, got: %s", out)
	}
	if !strings.Contains(out, "LEFT") {
		t.Errorf("expected LEFT status in output, got: %s", out)
	}
}

func TestFormatTable_MissingInLeft(t *testing.T) {
	result := Result{
		MissingInLeft: []string{"NEW_FLAG"},
	}
	var sb strings.Builder
	if err := FormatTable(&sb, result, "dev", "prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := sb.String()
	if !strings.Contains(out, "NEW_FLAG") {
		t.Errorf("expected key in output, got: %s", out)
	}
	if !strings.Contains(out, "RIGHT") {
		t.Errorf("expected RIGHT status in output, got: %s", out)
	}
}

func TestFormatTable_Mismatched(t *testing.T) {
	result := Result{
		Mismatched: []MismatchedKey{
			{Key: "DB_HOST", LeftValue: "localhost", RightValue: "db.prod.internal"},
		},
	}
	var sb strings.Builder
	if err := FormatTable(&sb, result, "dev", "prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := sb.String()
	if !strings.Contains(out, "DB_HOST") {
		t.Errorf("expected key in output, got: %s", out)
	}
	if !strings.Contains(out, "DIFF") {
		t.Errorf("expected DIFF status in output, got: %s", out)
	}
	if !strings.Contains(out, "localhost") {
		t.Errorf("expected left value in output, got: %s", out)
	}
}

func TestFormatTable_DefaultNames(t *testing.T) {
	result := Result{
		MissingInLeft: []string{"SOME_KEY"},
	}
	var sb strings.Builder
	if err := FormatTable(&sb, result, "", ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := sb.String()
	if !strings.Contains(out, "left") || !strings.Contains(out, "right") {
		t.Errorf("expected default names in header, got: %s", out)
	}
}
