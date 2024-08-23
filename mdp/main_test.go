package main

import (
	"fmt"
	"os"
	"testing"
)

const (
	inputFile = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result := parseContent(input)
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !compareIgnoringWhitespace(result, expected) {
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)
}

func compareIgnoringWhitespace(result, golden []byte) bool {
	ri, gi := 0, 0

	for ri < len(result) && gi < len(golden) {

		for ri < len(result) && (result[ri] == ' ' || result[ri] == '\n' || result[ri] == '\r' || result[ri] == '\t') {
			ri++
		}
		for gi < len(golden) && (golden[gi] == ' ' || golden[gi] == '\n' || golden[gi] == '\r' || golden[gi] == '\t') {
			gi++
		}
		// Compare non-whitespace characters
		if ri < len(result) && gi < len(golden) && result[ri] != golden[gi] {
			fmt.Printf("Mismatch at byte %d:\n result: %d\n golden: %d\n", ri, result[ri], golden[gi])
			return false
		}
		if ri >= len(result) || gi >= len(golden) {
			break
		}

		ri++
		gi++
	}

	// Check if either file still has remaining non-whitespace content
	if ri != len(result) || gi != len(golden) {
		fmt.Println("Lengths do not match after trimming whitespace.")
		return false
	}

	return true
}