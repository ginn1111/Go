package myflag

import (
	"flag"
	"fmt"
	"os"
)

type GreetingCommand struct {
	fs   *flag.FlagSet
	name string
}

func NewGreetingCommand() *GreetingCommand {
	gc := &GreetingCommand{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "", "Name")

	return gc
}

func (g *GreetingCommand) Name() string {
	return g.fs.Name()
}

func (g *GreetingCommand) Init(args []string) {
	g.fs.Parse(args)
}

func (g *GreetingCommand) Run() {
	fmt.Println("Greeting", g.name)
}

type Runner interface {
	Init([]string)
	Name() string
	Run()
}

func SubCommand() {

	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Please provide the subcommand")
		return
	}

	cmds := []Runner{
		NewGreetingCommand(),
	}

	subCommand := os.Args[1]

	for _, cmd := range cmds {

		if cmd.Name() == subCommand {
			cmd.Init(os.Args[2:])
			cmd.Run()
		}

	}

}
