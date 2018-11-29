package amuro

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type AmuroCmd struct {
	mobilesuit string
}

func (*AmuroCmd) Name() string {
	return "Newtype"
}

func (*AmuroCmd) Synopsis() string {
	return "Wake up newtype of in you"
}

func (*AmuroCmd) Usage() string {
	return "Usage: fulhurontal [option]"
}

func (c *AmuroCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.mobilesuit, "MobileSuit", "", "this is working")
}

func (c *AmuroCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
