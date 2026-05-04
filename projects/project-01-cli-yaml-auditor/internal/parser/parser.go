package parser

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Manifest struct {
	File       string
	Kind       string
	Containers []Container
}

type Container struct {
	Name       string
	Image      string
	Privileged bool
	HasLimits  bool
}

var (
	reKind       = regexp.MustCompile(`(?m)^kind:\s*([A-Za-z0-9]+)\s*$`)
	reContainer  = regexp.MustCompile(`(?m)^\s*-\s*name:\s*([A-Za-z0-9-_.]+)\s*$`)
	reImage      = regexp.MustCompile(`(?m)^\s*image:\s*([^\s]+)\s*$`)
	rePrivileged = regexp.MustCompile(`(?m)^\s*privileged:\s*true\s*$`)
	reLimits     = regexp.MustCompile(`(?m)^\s*limits:\s*$`)
)

func LoadManifests(root string) ([]Manifest, error) {
	if strings.TrimSpace(root) == "" {
		return nil, errors.New("root path is required")
	}

	var manifests []Manifest
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".yml") {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		m := parseManifest(path, string(b))
		manifests = append(manifests, m)
		return nil
	})
	return manifests, err
}

func parseManifest(filePath, raw string) Manifest {
	m := Manifest{File: filePath}
	if kind := reKind.FindStringSubmatch(raw); len(kind) == 2 {
		m.Kind = kind[1]
	}

	containerNames := reContainer.FindAllStringSubmatchIndex(raw, -1)
	for idx, c := range containerNames {
		name := raw[c[2]:c[3]]
		start := c[1]
		end := len(raw)
		if idx+1 < len(containerNames) {
			end = containerNames[idx+1][0]
		}
		block := raw[start:end]

		container := Container{Name: name}
		if img := reImage.FindStringSubmatch(block); len(img) == 2 {
			container.Image = img[1]
		}
		container.Privileged = rePrivileged.MatchString(block)
		container.HasLimits = reLimits.MatchString(block)
		m.Containers = append(m.Containers, container)
	}
	return m
}
