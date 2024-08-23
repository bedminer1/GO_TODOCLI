package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
		<title>Markdown Preview Tool</title>
	</head>
	<body>
	`
	footer = `
	</body>
</html>
	`
)

func main() {
	fileName := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*fileName, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string, out io.Writer) error {
	// read data from file
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)

	return saveHTML(outName, htmlData)
}

// convert Markdown to HTML
func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	var buffer bytes.Buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

// writes htmlData to a file
func saveHTML(outName string, htmlData []byte) error{
	return os.WriteFile(outName, htmlData, 0644)
}

// takes temp file and opens in a browser
func preview(fname string) error {
	cName := ""
	cParams := []string{}

	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("Os not supported")
	}

	cParams = append(cParams, fname)
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	return exec.Command(cPath, cParams...).Run()
}
