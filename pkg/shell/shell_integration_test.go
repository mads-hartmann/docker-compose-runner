//go:build integration

package shell

import (
	"strings"
	"testing"
)

func TestCommand(t *testing.T) {

	data := []struct {
		name    string
		command string
		stdout  string
		stderr  string
		errMsg  string
	}{
		{"Filesystem access", "file testdata/example.md", "testdata/example.md: ASCII text", "", ""},
		{"Echo stdout", "echo Hello to stdout", "Hello to stdout", "", ""},
		{"Echo stderr", ">&2 echo Hello to stderr", "", "Hello to stderr", ""},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			stdout, stderr, err := Command(d.command)

			if strings.TrimSpace(stdout) != d.stdout {
				t.Errorf("Expected stdout of '%s' got '%s'", d.stdout, stdout)
			}

			if strings.TrimSpace(stderr) != d.stderr {
				t.Errorf("Expected stderr of '%s' got '%s'", d.stderr, stderr)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if strings.TrimSpace(errMsg) != d.errMsg {
				t.Errorf("Expected error message '%s', got '%s'", d.errMsg, errMsg)
			}
		})
	}
}
