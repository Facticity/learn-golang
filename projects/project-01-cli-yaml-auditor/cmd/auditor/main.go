package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/example/go-k8s-sec-tooling-course/projects/project-01-cli-yaml-auditor/internal/audit"
	"github.com/example/go-k8s-sec-tooling-course/projects/project-01-cli-yaml-auditor/internal/parser"
	"github.com/example/go-k8s-sec-tooling-course/projects/project-01-cli-yaml-auditor/internal/report"
)

func main() {
	input := flag.String("input", "", "directory with kubernetes yaml manifests")
	flag.Parse()
	if *input == "" {
		fmt.Println("--input is required")
		os.Exit(1)
	}

	manifests, err := parser.LoadManifests(*input)
	if err != nil {
		fmt.Println("error loading manifests:", err)
		os.Exit(1)
	}

	findings := audit.RunChecks(manifests)
	report.Print(findings)
	if len(findings) > 0 {
		os.Exit(2)
	}
}
