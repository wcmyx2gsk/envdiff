package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// buildBinary compiles the CLI binary into a temp dir and returns its path.
func buildBinary(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	bin := filepath.Join(dir, "envdiff")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return bin
}

func writeTempEnv(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "*.env")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

func TestCLI_CleanExit0(t *testing.T) {
	bin := buildBinary(t)
	left := writeTempEnv(t, "KEY=value\nFOO=bar\n")
	right := writeTempEnv(t, "KEY=value\nFOO=bar\n")
	cmd := exec.Command(bin, left, right)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected exit 0, got error: %v\noutput: %s", err, out)
	}
}

func TestCLI_DiffExit1(t *testing.T) {
	bin := buildBinary(t)
	left := writeTempEnv(t, "KEY=value\nEXTRA=only_left\n")
	right := writeTempEnv(t, "KEY=value\n")
	cmd := exec.Command(bin, left, right)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected exit 1 for diff, got exit 0")
	}
	if !strings.Contains(string(out), "EXTRA") {
		t.Errorf("expected output to mention EXTRA key, got: %s", out)
	}
}

func TestCLI_JSONFormat(t *testing.T) {
	bin := buildBinary(t)
	left := writeTempEnv(t, "KEY=value\n")
	right := writeTempEnv(t, "KEY=value\n")
	cmd := exec.Command(bin, "-format=json", left, right)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %v\noutput: %s", err, out)
	}
	if !strings.Contains(string(out), "\"clean\"") {
		t.Errorf("expected JSON output with 'clean' field, got: %s", out)
	}
}

func TestCLI_MissingFile(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin, "nonexistent.env", "also_missing.env")
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected non-zero exit for missing file")
	}
}
