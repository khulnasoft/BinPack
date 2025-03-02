package goinstall

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"github.com/khulnasoft/binpack"
	"github.com/khulnasoft/binpack/internal/log"
)

var _ binpack.Installer = (*Installer)(nil)

type InstallerParameters struct {
	Module     string   `json:"module" yaml:"module" mapstructure:"module"`
	Entrypoint string   `json:"entrypoint" yaml:"entrypoint" mapstructure:"entrypoint"`
	LDFlags    []string `json:"ldflags" yaml:"ldflags" mapstructure:"ldflags"`
	Args       []string `json:"args" yaml:"args" mapstructure:"args"`
	Env        []string `json:"env" yaml:"env" mapstructure:"env"`
}

type Installer struct {
	config          InstallerParameters
	goInstallRunner func(spec, ldflags string, args []string, env []string, destDir string) error
}

func NewInstaller(cfg InstallerParameters) Installer {
	return Installer{
		config:          cfg,
		goInstallRunner: runGoInstall,
	}
}

func (i Installer) InstallTo(version, destDir string) (string, error) {
	path := i.config.Module
	if i.config.Entrypoint != "" {
		path += "/" + i.config.Entrypoint
	}
	fields := strings.Split(path, "/")
	binName := fields[len(fields)-1]
	binPath := filepath.Join(destDir, binName)

	spec := fmt.Sprintf("%s@%s", path, version)
	if strings.HasPrefix(i.config.Module, ".") || strings.HasPrefix(i.config.Module, "/") {
		spec = path
		log.WithFields("module", i.config.Module, "version", version).Debug("installing go module (local)")
	} else {
		log.WithFields("module", i.config.Module, "version", version).Debug("installing go module (remote)")
	}

	ldflags, err := templateFlags(i.config.LDFlags, version)
	if err != nil {
		return "", fmt.Errorf("failed to template ldflags: %v", err)
	}

	args, err := templateSlice(i.config.Args, version)
	if err != nil {
		return "", fmt.Errorf("failed to template args: %v", err)
	}

	if err := validateEnvSlice(i.config.Env); err != nil {
		return "", err
	}

	env, err := templateSlice(i.config.Env, version)
	if err != nil {
		return "", fmt.Errorf("failed to template env: %v", err)
	}

	if err := i.goInstallRunner(spec, ldflags, args, env, destDir); err != nil {
		return "", fmt.Errorf("failed to install: %v", err)
	}

	return binPath, nil
}

func validateEnvSlice(env []string) error {
	for _, e := range env {
		if !strings.Contains(e, "=") {
			return fmt.Errorf("invalid env format: %q", e)
		}
	}
	return nil
}

func templateSlice(in []string, version string) ([]string, error) {
	ret := make([]string, len(in))
	for i, arg := range in {
		rendered, err := templateString(arg, version)
		if err != nil {
			return nil, err
		}

		ret[i] = rendered
	}
	return ret, nil
}

func templateFlags(ldFlags []string, version string) (string, error) {
	flags := strings.Join(ldFlags, " ")

	return templateString(flags, version)
}

func templateString(in string, version string) (string, error) {
	tmpl, err := template.New("ldflags").Funcs(sprig.FuncMap()).Parse(in)
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, map[string]string{
		"Version": version,
	})

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func runGoInstall(spec, ldflags string, userArgs, userEnv []string, destDir string) error {
	args := []string{"install"}
	args = append(args, userArgs...)

	if ldflags != "" {
		args = append(args, fmt.Sprintf("-ldflags=%s", ldflags))
	}
	args = append(args, spec)

	log.WithFields("env-vars", len(userEnv)).Trace("running: go " + strings.Join(args, " "))

	cmd := exec.Command("go", args...)

	// set env vars...
	env := os.Environ()
	env = append(env, userEnv...)
	// always override any conflicting env vars
	env = append(env, "GOBIN="+destDir)
	cmd.Env = env

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("installation failed: %v\nOutput: %s", err, output)
	}
	return nil
}
