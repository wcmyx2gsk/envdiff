package diff

// MismatchedKey holds a key and its differing values between two env files.
type MismatchedKey struct {
	Key        string `json:"key"`
	LeftValue  string `json:"left_value"`
	RightValue string `json:"right_value"`
}

// Result holds the full comparison output between two env files.
type Result struct {
	// MissingInRight contains keys present in the left file but absent in the right.
	MissingInRight []string `json:"missing_in_right"`

	// MissingInLeft contains keys present in the right file but absent in the left.
	MissingInLeft []string `json:"missing_in_left"`

	// Mismatched contains keys present in both files but with different values.
	Mismatched []MismatchedKey `json:"mismatched"`
}

// Clean returns true when there are no differences between the two env files.
func (r Result) Clean() bool {
	return len(r.MissingInRight) == 0 &&
		len(r.MissingInLeft) == 0 &&
		len(r.Mismatched) == 0
}

// Summary returns a brief one-line description of the result.
func (r Result) Summary() string {
	if r.Clean() {
		return "files are identical"
	}
	total := len(r.MissingInRight) + len(r.MissingInLeft) + len(r.Mismatched)
	switch {
	case total == 1:
		return "1 difference found"
	default:
		return fmt.Sprintf("%d differences found", total)
	}
}
