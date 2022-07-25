//go:build integration

package shell

import (
	"strings"
	"testing"
)

func TestCommand(t *testing.T) {
	stdout, stderr, err := Command("file testdata/example.md")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if stderr != "" {
		t.Error("Expected empty stderr, got", stderr)
	}
	if strings.TrimSpace(stdout) != "testdata/example.md: ASCII text" {
		t.Error("Expected foo got", stdout)
	}
}
