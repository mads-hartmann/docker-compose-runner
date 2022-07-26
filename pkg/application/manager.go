package application

import (
	"strings"

	"github.com/mads-hartmann/docker-compose-runner/pkg/shell"
)

type Manager struct {
	ProjectName       string
	DockerComposePath string
}

func (m *Manager) Run() error {
	cmd := strings.Join([]string{
		"docker-compose",
		"--file",
		m.DockerComposePath,
		"--project-name",
		m.ProjectName,
		"up",
	}, " ")
	_, _, err := shell.Command(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Stop() error {
	cmd := strings.Join([]string{
		"docker-compose",
		"--file",
		m.DockerComposePath,
		"--project-name",
		m.ProjectName,
		"down",
	}, " ")
	_, _, err := shell.Command(cmd)
	if err != nil {
		return err
	}
	return nil
}
