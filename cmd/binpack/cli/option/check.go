package option

import (
	"github.com/khulnasoft/gob"
)

type Check struct {
	VerifySHA256Digest bool `json:"verify-sha256" yaml:"verify-sha256" mapstructure:"verify-sha256"`
}

func (o *Check) AddFlags(flags gob.FlagSet) {
	flags.BoolVarP(&o.VerifySHA256Digest, "verify-sha256", "", "Verifying the sha256 digest of already installed tools (by default xxh64 is used)")
}
