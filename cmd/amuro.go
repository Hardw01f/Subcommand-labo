package amuro

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type AmuroCmd struct {
	mobilesuit string
}

func (*AmuroCmd) Name() string {
	return "Amuro"
}

func (*AmuroCmd) Synopsis() string {
	return "The best of one NewType"
}

func (*AmuroCmd) Usage() string {
	return "Usage: fulhurontal [option]"
}

func (c *AmuroCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.mobilesuit, "MobileSuit", "", "this is working")
}

func (c *AmuroCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("Amuro, GAMDOM!!!  Rift Off!!!")
	return subcommands.ExitSuccess
}
