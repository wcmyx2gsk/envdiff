package diff

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestFormatJSON_Clean(t *testing.T) {
	r := Result{}
	var buf bytes.Buffer
	if err := FormatJSON(&buf, r, []string{"dev", "prod"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var report JSONReport
	if err := json.Unmarshal(buf.Bytes(), &report); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if !report.Clean {
		t.Errorf("expected clean=true")
	}
	if len(report.MissingIn) != 0 {
		t.Errorf("expected no missing keys, got %v", report.MissingIn)
	}
	if len(report.Mismatched) != 0 {
		t.Errorf("expected no mismatches, got %v", report.Mismatched)
	}
}

func TestFormatJSON_MissingKeys(t *testing.T) {
	r := Result{
		MissingInRight: []string{"SECRET"},
		MissingInLeft:  []string{"NEW_FLAG"},
	}
	var buf bytes.Buffer
	if err := FormatJSON(&buf, r, []string{"dev", "prod"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var report JSONReport
	if err := json.Unmarshal(buf.Bytes(), &report); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if report.Clean {
		t.Errorf("expected clean=false")
	}
	if got := report.MissingIn["prod"]; len(got) != 1 || got[0] != "SECRET" {
		t.Errorf("expected prod missing SECRET, got %v", got)
	}
	if got := report.MissingIn["dev"]; len(got) != 1 || got[0] != "NEW_FLAG" {
		t.Errorf("expected dev missing NEW_FLAG, got %v", got)
	}
}

func TestFormatJSON_Mismatched(t *testing.T) {
	r := Result{
		Mismatched: []Mismatch{{Key: "PORT", LeftValue: "3000", RightValue: "8080"}},
	}
	var buf bytes.Buffer
	if err := FormatJSON(&buf, r, []string{"dev", "prod"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var report JSONReport
	if err := json.Unmarshal(buf.Bytes(), &report); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(report.Mismatched) != 1 {
		t.Fatalf("expected 1 mismatch, got %d", len(report.Mismatched))
	}
	m := report.Mismatched[0]
	if m.Key != "PORT" {
		t.Errorf("expected key PORT, got %s", m.Key)
	}
	if m.Values["dev"] != "3000" || m.Values["prod"] != "8080" {
		t.Errorf("unexpected values: %v", m.Values)
	}
}

func TestFormatJSON_InvalidNames(t *testing.T) {
	var buf bytes.Buffer
	if err := FormatJSON(&buf, Result{}, []string{"only-one"}); err == nil {
		t.Error("expected error for wrong number of names")
	}
}
