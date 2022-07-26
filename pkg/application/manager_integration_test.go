//go:build integration

package application

import (
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	manager := Manager{
		ProjectName:       "manager-test",
		DockerComposePath: "./testdata/docker-compose.yml",
	}
	go func() {
		err := manager.Run()
		if err != nil {
			t.Error("Expected no error but got", err)
		}
	}()

	// TODO: Find a better way to check for "readiness" and then perform a HTTP call to test
	// that it works as expected.
	t.Log("Sleeping 30s")
	time.Sleep(30 * time.Second)

	err := manager.Stop()
	if err != nil {
		t.Error("Expected no error but got", err)
	}

}
