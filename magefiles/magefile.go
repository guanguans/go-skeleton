// This file is part of the guanguans/go-skeleton.
// (c) 2022. guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/magefile/mage/sh"
)

// Command aliases.
var Aliases = map[string]interface{}{
	"Speak": Say,
}

// Say says something.
func Say(msg string) error {
	_, err := fmt.Printf("%v(%T)\n", msg, msg)
	return err
}

var releaseTag = regexp.MustCompile(`^v1\.[0-9]+\.[0-9]+$`)

// Generates a new release. Expects a version tag in v1.x.x format.
func Release(tag string) (err error) {
	if _, err := exec.LookPath("goreleaser"); err != nil {
		return fmt.Errorf("can't find goreleaser: %w", err)
	}
	if !releaseTag.MatchString(tag) {
		return errors.New("TAG environment variable must be in semver v1.x.x format, but was " + tag)
	}

	if err := sh.RunV("git", "tag", "-a", tag, "-m", tag); err != nil {
		return err
	}
	if err := sh.RunV("git", "push", "origin", tag); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			sh.RunV("git", "tag", "--delete", tag)
			sh.RunV("git", "push", "--delete", "origin", tag)
		}
	}()
	return sh.RunV("goreleaser")
}

// Remove the temporarily generated files from Release.
func Clean() error {
	return sh.Rm("dist")
}
