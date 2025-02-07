package tool

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/khulnasoft/binpack"
	"github.com/khulnasoft/binpack/tool/git"
	"github.com/khulnasoft/binpack/tool/githubrelease"
	"github.com/khulnasoft/binpack/tool/goproxy"
)

func VersionResolverMethods() []string {
	return []string{
		githubrelease.ResolveMethod,
		goproxy.ResolveMethod,
		git.ResolveMethod,
	}
}

func ResolveVersion(tool binpack.VersionResolver, intent binpack.VersionIntent) (string, error) {
	want := intent.Want
	constraint := intent.Constraint

	var resolvedVersion string

	resolvedVersion, err := tool.ResolveVersion(want, constraint)
	if err != nil {
		return "", fmt.Errorf("failed to resolve version: %w", err)
	}

	if constraint != "" {
		ver, err := semver.NewVersion(resolvedVersion)
		if err == nil {
			constraintObj, err := semver.NewConstraint(constraint)
			if err != nil {
				return resolvedVersion, fmt.Errorf("invalid version constraint: %v", err)
			}

			if !constraintObj.Check(ver) {
				return resolvedVersion, fmt.Errorf("resolved version %q is unsatisfied by constraint %q. Remove the constraint or run 'update' to re-pin a valid version", resolvedVersion, constraint)
			}
		}
	}
	return resolvedVersion, nil
}
