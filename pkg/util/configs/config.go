package configs

import (
	"github.com/Knetic/govaluate"
	flags "github.com/jessevdk/go-flags"
)

// ParseArguments parses the program arguments
func ParseArguments(args []string) {
	_, argsError := flags.ParseArgs(&Opts, args)
	if argsError != nil {
		panic(argsError)
	}
}

// EvaluateExpression evaluates the given expression
func EvaluateExpression() (*govaluate.EvaluableExpression, error) {
	return govaluate.NewEvaluableExpression(Opts.Expression)
}
