package audit

import (
	"testing"

	"github.com/example/go-k8s-sec-tooling-course/projects/project-01-cli-yaml-auditor/internal/parser"
)

func TestRunChecks(t *testing.T) {
	m := parser.Manifest{File: "test.yaml", Containers: []parser.Container{{Name: "app", Image: "nginx:latest", HasLimits: false, Privileged: true}}}
	findings := RunChecks([]parser.Manifest{m})
	if len(findings) != 3 {
		t.Fatalf("expected 3 findings, got %d", len(findings))
	}
}
