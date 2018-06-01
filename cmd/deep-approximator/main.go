package main

import (
	"os"
	"time"

	"github.com/TheDonDope/deep-approximator/pkg/util/configs"
	"github.com/TheDonDope/deep-approximator/pkg/util/logs"
)

func main() {
	configs.ParseArguments(os.Args)

	start := time.Now()
	logs.Printfln("Starting deep-approximator @ %v", start.Format(time.RFC3339))
}
