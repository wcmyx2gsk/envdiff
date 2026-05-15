package diff_test

import (
	"strings"
	"testing"

	"github.com/user/envdiff/internal/diff"
)

func TestFormatText_Clean(t *testing.T) {
	result := diff.Result{}
	out := diff.FormatText(result, "a.env", "b.env")
	if !strings.Contains(out, "No differences") {
		t.Errorf("expected 'No differences', got: %s", out)
	}
}

func TestFormatText_MissingInRight(t *testing.T) {
	result := diff.Result{
		MissingInRight: []string{"SECRET_KEY", "DB_HOST"},
	}
	out := diff.FormatText(result, "prod.env", "staging.env")
	if !strings.Contains(out, "Missing in staging.env") {
		t.Errorf("expected missing-in-right header, got: %s", out)
	}
	if !strings.Contains(out, "SECRET_KEY") {
		t.Errorf("expected SECRET_KEY in output, got: %s", out)
	}
	if !strings.Contains(out, "DB_HOST") {
		t.Errorf("expected DB_HOST in output, got: %s", out)
	}
}

func TestFormatText_MissingInLeft(t *testing.T) {
	result := diff.Result{
		MissingInLeft: []string{"NEW_FEATURE_FLAG"},
	}
	out := diff.FormatText(result, "prod.env", "staging.env")
	if !strings.Contains(out, "Missing in prod.env") {
		t.Errorf("expected missing-in-left header, got: %s", out)
	}
	if !strings.Contains(out, "NEW_FEATURE_FLAG") {
		t.Errorf("expected NEW_FEATURE_FLAG in output, got: %s", out)
	}
}

func TestFormatText_Mismatched(t *testing.T) {
	result := diff.Result{
		Mismatched: []diff.Mismatch{
			{Key: "APP_ENV", LeftVal: "production", RightVal: "staging"},
		},
	}
	out := diff.FormatText(result, "prod.env", "staging.env")
	if !strings.Contains(out, "Mismatched values") {
		t.Errorf("expected 'Mismatched values' header, got: %s", out)
	}
	if !strings.Contains(out, "APP_ENV") {
		t.Errorf("expected APP_ENV in output, got: %s", out)
	}
	if !strings.Contains(out, "production") {
		t.Errorf("expected left value in output, got: %s", out)
	}
	if !strings.Contains(out, "staging") {
		t.Errorf("expected right value in output, got: %s", out)
	}
}

func TestFormatText_DefaultNames(t *testing.T) {
	result := diff.Result{
		MissingInRight: []string{"KEY"},
	}
	out := diff.FormatText(result, "", "")
	if !strings.Contains(out, "Missing in right") {
		t.Errorf("expected default name 'right', got: %s", out)
	}
}
