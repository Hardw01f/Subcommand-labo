package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Snow-HardWolf/Subcommand-labo/cmd/amuro"

	"github.com/google/subcommands"
)

type printCmd struct {
	capitalize bool
}

func (*printCmd) Name() string     { return "print" }
func (*printCmd) Synopsis() string { return "Print args to stdout." }
func (*printCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *printCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *printCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

type NewtypeCmd struct {
	goust     bool
	actername string
}

func (*NewtypeCmd) Name() string {
	return "Newtype"
}

func (*NewtypeCmd) Synopsis() string {
	return "Wake up newtype of in you"
}

func (*NewtypeCmd) Usage() string {
	return "Usage: fulhurontal [option]"
}

func (c *NewtypeCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.goust, "goust", false, "goust output")
	f.StringVar(&c.actername, "banaji", "", "banaji is Newtype")
}

func (c *NewtypeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	arg := f.Args()
	fmt.Printf("%+v\n", arg)

	if arg[0] == "banaji" {
		fmt.Println("Belive A beast named possibility")
	} else {
		fmt.Println("----------~~~~~~~^^^^^^^^^~~~~~~~------------")
		fmt.Println("the mobileSuits is approaching at 3 times normal speed")
		fmt.Println("There is nothing to say without hitting it")
		fmt.Println("Is the whole body made of psycho frame?")
	}
	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&printCmd{}, "")
	subcommands.Register(&NewtypeCmd{}, "")
	subcommands.Register(&amuro.AmuroCmd{}, "Amuro")
	//subcommands.Register(&mineva.MinevaCmd{},"Mineva")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
