// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2021 Datadog, Inc.

package api

import (
	"strings"

	chaostypes "github.com/DataDog/chaos-controller/types"
)

// DisruptionKind contains all methods required for a disruption sub-specification (Network, DNS, CPUPressure, etc.)
type DisruptionKind interface {
	// generates CLI args for the given disruption sub-specification
	GenerateArgs() []string

	// validates schema for the given disruption sub-specification
	Validate() error
}

// AppendCommonArgs is a helper function generating common args and appending them to the given args array
func AppendCommonArgs(args []string, level chaostypes.DisruptionLevel, containerIDs []string, sink string, dryRun bool,
	disruptionName string, disruptionNamespace string, targetName string) []string {
	args = append(args,
		// basic args
		"--metrics-sink", sink,
		"--level", string(level),
		"--containers-id", strings.Join(containerIDs, ","),

		// log context args
		"--log-context-disruption-name", disruptionName,
		"--log-context-disruption-namespace", disruptionNamespace,
		"--log-context-target-name", targetName,
	)

	// enable dry-run mode
	if dryRun {
		args = append(args,
			"--dry-run",
		)
	}

	return args
}