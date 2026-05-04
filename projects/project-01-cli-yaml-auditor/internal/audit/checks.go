package audit

import (
	"strings"

	"github.com/example/go-k8s-sec-tooling-course/projects/project-01-cli-yaml-auditor/internal/parser"
)

type Finding struct {
	File     string
	Rule     string
	Severity string
	Message  string
}

func RunChecks(manifests []parser.Manifest) []Finding {
	var out []Finding
	for _, m := range manifests {
		for _, c := range m.Containers {
			if c.Privileged {
				out = append(out, Finding{File: m.File, Rule: "no-privileged", Severity: "high", Message: c.Name + " runs privileged"})
			}
			if !c.HasLimits {
				out = append(out, Finding{File: m.File, Rule: "require-limits", Severity: "medium", Message: c.Name + " missing cpu/memory limits"})
			}
			if strings.HasSuffix(c.Image, ":latest") {
				out = append(out, Finding{File: m.File, Rule: "no-latest-tag", Severity: "low", Message: c.Name + " uses latest tag"})
			}
		}
	}
	return out
}
