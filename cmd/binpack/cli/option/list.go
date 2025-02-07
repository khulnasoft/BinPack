package option

import "github.com/khulnasoft/gob"

type List struct {
	Updates       bool     `json:"updates" yaml:"updates" mapstructure:"updates"`
	IncludeFilter []string `json:"includeFilter" yaml:"includeFilter" mapstructure:"includeFilter"`
}

func (o *List) AddFlags(flags gob.FlagSet) {
	flags.BoolVarP(&o.Updates, "updates", "", "List only tool installations that need to be updated (relative to what is currently installed)")
}
