package amuro

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type amuroCmd struct {
	mobilesuit string
}

func (*amuroCmd) Name() string {
	return "Newtype"
}

func (*amuroCmd) Synopsis() string {
	return "Wake up newtype of in you"
}

func (*amuroCmd) Usage() string {
	return "Usage: fulhurontal [option]"
}

func (c *amuroCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.goust, "goust", false, "goust output")
}

func (c *amuroCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
