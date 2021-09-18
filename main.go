package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const Version = "1.0.0"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	fmt.Fprintf(c.errStream, "gobook version %s \n", Version)
	return 0
}

func TestRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("gobook -version", " ")

	cli.Run(args)
	expected := fmt.Sprintf("gobook version %s", Version)
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
