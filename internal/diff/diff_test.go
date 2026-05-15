package diff_test

import (
	"testing"

	"github.com/user/envdiff/internal/diff"
)

func TestCompare_Clean(t *testing.T) {
	left := map[string]string{"A": "1", "B": "2"}
	right := map[string]string{"A": "1", "B": "2"}

	r := diff.Compare(left, right)
	if !r.IsClean() {
		t.Errorf("expected clean result, got %+v", r)
	}
}

func TestCompare_MissingInRight(t *testing.T) {
	left := map[string]string{"A": "1", "B": "2"}
	right := map[string]string{"A": "1"}

	r := diff.Compare(left, right)
	if len(r.MissingInRight) != 1 || r.MissingInRight[0] != "B" {
		t.Errorf("expected B missing in right, got %v", r.MissingInRight)
	}
}

func TestCompare_MissingInLeft(t *testing.T) {
	left := map[string]string{"A": "1"}
	right := map[string]string{"A": "1", "C": "3"}

	r := diff.Compare(left, right)
	if len(r.MissingInLeft) != 1 || r.MissingInLeft[0] != "C" {
		t.Errorf("expected C missing in left, got %v", r.MissingInLeft)
	}
}

func TestCompare_Mismatched(t *testing.T) {
	left := map[string]string{"A": "hello"}
	right := map[string]string{"A": "world"}

	r := diff.Compare(left, right)
	if len(r.Mismatched) != 1 {
		t.Fatalf("expected 1 mismatch, got %d", len(r.Mismatched))
	}
	m := r.Mismatched[0]
	if m.Key != "A" || m.LeftValue != "hello" || m.RightValue != "world" {
		t.Errorf("unexpected mismatch entry: %+v", m)
	}
}

func TestCompare_SortedOutput(t *testing.T) {
	left := map[string]string{"Z": "1", "A": "1", "M": "1"}
	right := map[string]string{}

	r := diff.Compare(left, right)
	expected := []string{"A", "M", "Z"}
	for i, k := range r.MissingInRight {
		if k != expected[i] {
			t.Errorf("position %d: want %s, got %s", i, expected[i], k)
		}
	}
}

func TestIsClean_False(t *testing.T) {
	r := diff.Result{
		MissingInRight: []string{"X"},
	}
	if r.IsClean() {
		t.Error("expected not clean")
	}
}
