package main

import (
	"os"
	"strings"
)

var isAdbEnabled = strings.ToLower(os.Getenv("ENABLE_ARANGODB"))
var isConsoleEnabled = strings.ToLower(os.Getenv("ENABLE_CONSOLE_OUTPUT"))

type TargetStrategy interface {
	push(match *CertStreamMessage) error
}

type Targets struct {
	targets []TargetStrategy
}

func initTargets() *Targets {
	var targets = []TargetStrategy{}

	if isAdbEnabled == "true" {
		adbClient = *AdbConnect()
		targets = append(targets, &AdbStrategy{})
	}
	if isConsoleEnabled == "true" {
		targets = append(targets, &ConsoleOutputStrategy{})
	}

	// Default to console output if no target strategy is defined
	if len(targets) == 0 {
		targets = append(targets, &ConsoleOutputStrategy{})
	}

	return &Targets{
		targets: targets,
	}
}

func (targets *Targets) push(match *CertStreamMessage) {
	for _, t := range targets.targets {
		t.push(match)
	}
}
