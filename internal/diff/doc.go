// Package diff provides functionality for comparing two parsed .env file maps
// and surfacing differences between them.
//
// The primary entry point is Compare, which accepts two map[string]string values
// (as returned by the parser package) and returns a Result describing:
//
//   - Keys present in the left file but missing from the right.
//   - Keys present in the right file but missing from the left.
//   - Keys present in both files whose values differ.
//
// All slices in the returned Result are sorted alphabetically for deterministic
// output, making results easy to display or test.
package diff
