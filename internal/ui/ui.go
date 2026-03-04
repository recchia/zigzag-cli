package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/recchia/zigzag-cli/internal/eas"
)

func SelectBuild(builds []eas.Build) (*eas.Build, error) {
	if len(builds) == 0 {
		return nil, fmt.Errorf("no builds found")
	}

	options := make([]string, len(builds))
	for i, b := range builds {
		options[i] = fmt.Sprintf("[%s] %s - %s (%s) - %s", b.Platform, b.BuildProfile, b.AppVersion, b.AppBuildVersion, b.CreatedAt)
	}

	var selectedIndex int
	prompt := &survey.Select{
		Message: "Select a build to download:",
		Options: options,
	}

	err := survey.AskOne(prompt, &selectedIndex)
	if err != nil {
		return nil, err
	}

	return &builds[selectedIndex], nil
}

func AskFileName(defaultName string) (string, error) {
	var name string
	prompt := &survey.Input{
		Message: "Enter the file name:",
		Default: defaultName,
	}
	err := survey.AskOne(prompt, &name)
	return name, err
}

func AskOutputDir() (string, error) {
	var dir string
	prompt := &survey.Input{
		Message: "Enter the base output directory:",
		Default: ".",
	}
	err := survey.AskOne(prompt, &dir)
	return dir, err
}
