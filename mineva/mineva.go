package mineva

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type MinevaCmd struct {
	zabifamily string
}

func (*MinevaCmd) Name() string {
	return "Mineva Lao Zabi"
}

func (*MinevaCmd) Synopsis() string {
	return "Neo Zion"
}

func (*MinevaCmd) Usage() string {
	return "Usage: fulhurontal [option]"
}

func (c *MinevaCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.zabifamily, "zabifamily", "", "the Last one of Zabi family")
}

func (c *MinevaCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("I'm Mineva Zabi , I'm don't escape !! don't hide !! That's My Way !!")
	return subcommands.ExitSuccess
}
