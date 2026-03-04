package eas

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Build struct {
	ID              string         `json:"id"`
	Status          string         `json:"status"`
	Platform        string         `json:"platform"`
	BuildProfile    string         `json:"buildProfile"`
	CreatedAt       string         `json:"createdAt"`
	Artifacts       BuildArtifacts `json:"artifacts,omitzero"`
	AppVersion      string         `json:"appVersion"`
	AppBuildVersion string         `json:"appBuildVersion"`
}

type BuildArtifacts struct {
	BuildUrl string `json:"buildUrl"`
}

type CommandRunner interface {
	Run(name string, args ...string) ([]byte, error)
}

type RealCommandRunner struct{}

func (r RealCommandRunner) Run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("command failed: %s", string(exitErr.Stderr))
		}
		return nil, err
	}
	return output, nil
}

func ListBuilds(runner CommandRunner) ([]Build, error) {
	output, err := runner.Run("eas", "build:list", "--json", "--non-interactive", "--limit", "20")
	if err != nil {
		return nil, fmt.Errorf("eas build:list failed: %w", err)
	}

	var builds []Build
	if err := json.Unmarshal(output, &builds); err != nil {
		return nil, fmt.Errorf("failed to parse builds: %w", err)
	}

	return builds, nil
}
