package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadManifests(t *testing.T) {
	d := t.TempDir()
	f := filepath.Join(d, "dep.yaml")
	raw := `apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
        - name: app
          image: nginx:latest
          securityContext:
            privileged: true
          resources:
            limits:
              cpu: "200m"`
	if err := os.WriteFile(f, []byte(raw), 0o644); err != nil {
		t.Fatal(err)
	}

	ms, err := LoadManifests(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(ms) != 1 || len(ms[0].Containers) != 1 {
		t.Fatalf("unexpected parse result: %+v", ms)
	}
	c := ms[0].Containers[0]
	if !c.Privileged || !c.HasLimits || c.Image != "nginx:latest" {
		t.Fatalf("container parse mismatch: %+v", c)
	}
}
