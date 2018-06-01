package main

import (
	"log"
	"os"
	"time"

	"github.com/TheDonDope/deep-approximator/pkg/api"
	"github.com/TheDonDope/deep-approximator/pkg/util/configs"
	"github.com/TheDonDope/deep-approximator/pkg/util/logs"
)

func main() {
	configs.ParseArguments(os.Args)
	start := time.Now()
	logs.Printfln("Starting deep-approximator @ %v", start.Format(time.RFC3339))
	approximatorImpl := &api.DeepApproximatorService{}
	approximatorImpl.InitNetwork()
	if configs.Opts.Learn {
		approximatorImpl.Learn()
		log.Println("Approximator learned successfully.")
	}
	if configs.Opts.Calculate {
		approximatorImpl.Calculate()
		log.Println("Approximator calculated successfully.")
	}
	logs.Printfln("Program arguments: %+v", configs.Opts)
	logs.Printfln("Finishing deep-approximator @ %v", time.Now().Format(time.RFC3339))
	logs.Printfln("Overall time spent: %v", time.Since(start))
}
