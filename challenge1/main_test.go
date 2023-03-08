package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("testString")

	wg.Wait()

	if msg != "testString" {
		t.Errorf("Expected to find testString, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "testString"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "testString") {
		t.Errorf("Expected to find testString, but it is not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find Hello, universe!, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find Hello, cosmos!, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find Hello, world!, but it is not there")
	}
}
