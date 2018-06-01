package configs

import (
	flags "github.com/jessevdk/go-flags"
)

// ParseArguments parses the program arguments
func ParseArguments(args []string) {
	_, argsError := flags.ParseArgs(&Opts, args)
	if argsError != nil {
		panic(argsError)
	}
}
