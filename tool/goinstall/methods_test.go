package goinstall

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/binpack/tool/goproxy"
)

func TestMethods(t *testing.T) {
	tests := []struct {
		name    string
		methods []string
		want    bool
	}{
		{
			name:    "valid",
			methods: []string{"go-install", "go", "go install", "goinstall", "golang"},
			want:    true,
		},
		{
			name:    "invalid",
			methods: []string{"made up"},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, method := range tt.methods {
				t.Run(method, func(t *testing.T) {
					t.Run("IsInstallMethod", func(t *testing.T) {
						assert.Equal(t, tt.want, IsInstallMethod(method))
					})
				})
			}
		})
	}
}

func TestDefaultVersionResolverConfig(t *testing.T) {
	tests := []struct {
		name          string
		installParams any
		wantMethod    string
		wantParams    any
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "valid",
			installParams: InstallerParameters{
				Module:     "github.com/khulnasoft/binpack",
				Entrypoint: "cmd/binpack",
				LDFlags:    []string{"-X main.version=1.0.0"},
			},
			wantMethod: goproxy.ResolveMethod,
			wantParams: goproxy.VersionResolutionParameters{
				Module: "github.com/khulnasoft/binpack",
			},
		},
		{
			name: "invalid",
			installParams: map[string]string{
				"module": "github.com/khulnasoft/binpack",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				tt.wantErr = assert.NoError
			}
			method, params, err := DefaultVersionResolverConfig(tt.installParams)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.wantMethod, method)
			assert.Equal(t, tt.wantParams, params)
		})
	}
}
