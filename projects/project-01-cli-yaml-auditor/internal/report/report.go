package report

import (
	"fmt"

	"github.com/example/go-k8s-sec-tooling-course/projects/project-01-cli-yaml-auditor/internal/audit"
)

func Print(findings []audit.Finding) {
	if len(findings) == 0 {
		fmt.Println("PASS: no findings")
		return
	}
	fmt.Printf("FAIL: %d findings\n", len(findings))
	for _, f := range findings {
		fmt.Printf("- [%s] %s (%s): %s\n", f.Severity, f.Rule, f.File, f.Message)
	}
}
