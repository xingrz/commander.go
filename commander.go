package commander

import (
	"os"
	"fmt"
)

/**
 * `Option` type, contains data for a specific option
 */

type Option struct {
	Name string
	Tiny byte
	Verbose string
	Description string
	Required bool
	Callback func()
}

/**
 * Commander type, contains all program data
 */

type Commander struct {
	Name string
	Version string
	Options []Option
}

/**
 * `Parse` arguments
 */

func (commander *Commander) Parse() {
	args := commander.explode(os.Args[1:])

	if(len(*args) == 0) {
		commander.Usage();
	}
}

func (commander *Commander) explode(args []string) *[]string {
	for i := range args {
		fmt.Println(args[i])
	}

	return &args
}

/**
 * `Add` `option` to the commander instance
 */

func (commander *Commander) Add(option *Option) {
	commander.Options = append(commander.Options, *option)
}

/**
 * Display the usage of `commander`
 */

func (commander *Commander) Usage() {
	fmt.Fprintf(os.Stderr, "\n  Usage: %s [options]\n\n", commander.Name)
	fmt.Fprintf(os.Stderr, "  Options:\n");

	options := commander.Options
	for i := range options {
		fmt.Fprintf(os.Stderr, "    -%c, --%s %s",
			options[i].Tiny, options[i].Verbose, options[i].Description)
	}
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

/**
 * Return the total number of arguments
 */

func (commander *Commander) Len() int {
	return len(os.Args)
}
