package main

import (
	"os"
	"time"

	"gitlab.com/TheDonDope/deep-approximator/pkg/api"
	"gitlab.com/TheDonDope/deep-approximator/pkg/util/configs"
	"gitlab.com/TheDonDope/deep-approximator/pkg/util/logs"
	"go.uber.org/zap"
)

func main() {
	configs.ParseArguments(os.Args)
	allStart := time.Now()
	logs.Logger.Info("Starting deep-approximator", zap.String("allStart", allStart.Format(time.RFC3339)))
	approximatorImpl := &api.DeepApproximatorService{}
	if configs.Opts.Learn {
		learnStart := time.Now()
		approximatorImpl.Learn()
		logs.Logger.Info("Approximator learned successfully.", zap.Duration("learnEnd", time.Since(learnStart)))
	}
	if configs.Opts.Calculate {
		calcStart := time.Now()
		approximatorImpl.Calculate()
		logs.Logger.Info("Approximator calculated successfully.", zap.Duration("calcEnd", time.Since(calcStart)))
	}
	allEnd := time.Now()
	logs.Logger.Info("Program arguments", zap.Any("args", configs.Opts))
	logs.Logger.Info("Finished deep-approximator", zap.String("allEnd", allEnd.Format(time.RFC3339)))
	logs.Logger.Info("Overall time spent", zap.Duration("overall", time.Since(allStart)))
}
