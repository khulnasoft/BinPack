package option

import (
	"fmt"

	"github.com/khulnasoft/binpack/tool"
	"github.com/khulnasoft/gob"
)

type VersionResolution struct {
	Constraint string `json:"constraint" yaml:"constraint" mapstructure:"constraint"`
	Method     string `json:"method" yaml:"method" mapstructure:"method"`
}

func (o *VersionResolution) AddFlags(flags gob.FlagSet) {
	flags.StringVarP(&o.Constraint, "constraint", "", "Version constraint (e.g. '<2.0' or '>=1.0.0')")
	flags.StringVarP(&o.Method, "version-from", "f", fmt.Sprintf("The method to use to resolve the version (available: %+v)", tool.VersionResolverMethods()))
}
