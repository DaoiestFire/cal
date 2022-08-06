package options

import "github.com/spf13/pflag"

type Options struct {
	ArgOne int
	ArgTwo int
}

func (o *Options) InitFlagSet() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.IntVar(&o.ArgOne, "argone", 0, "assign a value to arg1")
	flagSet.IntVar(&o.ArgTwo, "argtwo", 0, "assign a value to arg2")
	return flagSet
}
