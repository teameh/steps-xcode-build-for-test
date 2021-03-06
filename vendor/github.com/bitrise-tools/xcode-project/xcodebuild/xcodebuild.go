package xcodebuild

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-tools/xcode-project/serialized"
)

func parseShowBuildSettingsOutput(out string) (serialized.Object, error) {
	settings := serialized.Object{}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Build settings") {
			continue
		}

		if strings.HasPrefix(line, "User defaults from command line") {
			continue
		}

		if line == "" {
			continue
		}

		split := strings.Split(line, " = ")

		if len(split) < 2 {
			return nil, fmt.Errorf("unknown build settings: %s", line)
		}

		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(strings.Join(split[1:], " = "))

		settings[key] = value
	}

	return settings, nil
}

// ShowProjectBuildSettings ...
func ShowProjectBuildSettings(project, target, configuration string, customOptions ...string) (serialized.Object, error) {
	args := []string{"-project", project, "-target", target, "-configuration", configuration}
	args = append(args, "-showBuildSettings")
	args = append(args, customOptions...)

	cmd := command.New("xcodebuild", args...)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s failed: %s", cmd.PrintableCommandArgs(), err)
	}

	return parseShowBuildSettingsOutput(out)
}

// ShowWorkspaceBuildSettings ...
func ShowWorkspaceBuildSettings(workspace, scheme, configuration string, customOptions ...string) (serialized.Object, error) {
	args := []string{"-workspace", workspace, "-scheme", scheme, "-configuration", configuration}
	args = append(args, "-showBuildSettings")
	args = append(args, customOptions...)

	cmd := command.New("xcodebuild", args...)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s failed: %s", cmd.PrintableCommandArgs(), err)
	}

	return parseShowBuildSettingsOutput(out)
}
