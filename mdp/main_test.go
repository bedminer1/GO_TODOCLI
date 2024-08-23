package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	inputFile = "./testdata/test1.md"
	goldenFile = "./testdata/test1.md.html"
)

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer
	if err := run(inputFile, &mockStdOut, "", false); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStdOut.String())
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Error("result content does not match")
	}
}