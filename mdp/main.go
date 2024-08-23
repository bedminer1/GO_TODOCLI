package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

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
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*fileName, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string, out io.Writer, skipPreview bool) error {
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

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outName)

	return preview(outName)
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
		return fmt.Errorf("os not supported")
	}

	cParams = append(cParams, fname)
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	err = exec.Command(cPath, cParams...).Run()
	time.Sleep(2 * time.Second) // give browser time to open before deleting
	return err
}
