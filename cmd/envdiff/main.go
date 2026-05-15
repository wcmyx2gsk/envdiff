package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/user/envdiff/internal/diff"
	"github.com/user/envdiff/internal/parser"
)

func main() {
	format := flag.String("format", "text", "Output format: text, json, table")
	leftName := flag.String("left-name", "", "Label for the left file (default: filename)")
	rightName := flag.String("right-name", "", "Label for the right file (default: filename)")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: envdiff [flags] <left.env> <right.env>")
		flag.PrintDefaults()
		os.Exit(2)
	}

	leftPath := args[0]
	rightPath := args[1]

	leftEnv, err := parser.ParseFile(leftPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", leftPath, err)
		os.Exit(1)
	}

	rightEnv, err := parser.ParseFile(rightPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", rightPath, err)
		os.Exit(1)
	}

	left := leftPath
	if *leftName != "" {
		left = *leftName
	}
	right := rightPath
	if *rightName != "" {
		right = *rightName
	}

	result := diff.Compare(leftEnv, rightEnv)

	var output string
	switch *format {
	case "json":
		output, err = diff.FormatJSON(result, left, right)
	case "table":
		output, err = diff.FormatTable(result, left, right)
	case "text":
		output, err = diff.FormatText(result, left, right), nil
	default:
		fmt.Fprintf(os.Stderr, "unknown format %q; choose text, json, or table\n", *format)
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "formatting error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(output)

	if !result.Clean() {
		os.Exit(1)
	}
}
