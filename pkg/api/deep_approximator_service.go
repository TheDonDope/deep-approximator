package api

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	"gitlab.com/TheDonDope/deep-approximator/pkg/types"
	"gitlab.com/TheDonDope/deep-approximator/pkg/util/configs"
	"gitlab.com/TheDonDope/deep-approximator/pkg/util/errors"
	"gitlab.com/TheDonDope/deep-approximator/pkg/util/files"
	"gitlab.com/TheDonDope/deep-approximator/pkg/util/logs"
	"github.com/dariubs/percent"
	"go.uber.org/zap"
)

// DeepApproximatorService implements the DeepApproximator interface
type DeepApproximatorService struct{}

// Learn starts the learning process of the neural network
func (impl DeepApproximatorService) Learn() {
	random := newRandom()
	network := createNetwork()
	maxValue := 2.
	minValue := 0.
	for i := 0; i < configs.Opts.Rounds; i++ {
		x := minValue + random.Float64()*(maxValue-minValue)
		y := minValue + random.Float64()*(maxValue-minValue)

		expression, err := configs.EvaluateExpression()
		errors.HandleError(err, fmt.Sprintf("Unable to interpret Expression (-e): %v", err))
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = x
		parameters["y"] = y
		result, err := expression.Evaluate(parameters)
		errors.HandleError(err, fmt.Sprintf("Unable to evaluate Expression (-e): %v", err))

		// actual learning process
		// learn.Learn(network, []float64{x, y}, []float64{math.Sin(x + y)}, configs.Opts.Speed)
		learn.Learn(network, []float64{x, y}, []float64{downscale(math.Sin(math.Abs(math.Sin(result.(float64)))))}, configs.Opts.Speed)

		if i%(configs.Opts.Rounds/10) == 0 {
			// logs.Logger.Info("Current progress", zap.Int("curr", i), zap.Int("all", configs.Opts.Rounds), zap.Float64("%", percent.PercentOf(i, configs.Opts.Rounds)))
			logs.Logger.Info("Current progress", zap.Float64("%", percent.PercentOf(i, configs.Opts.Rounds)))
		}
	}
	persist.ToFile(configs.Opts.Output, network)
}

// Calculate starts the calculation process of the nerual network
func (impl DeepApproximatorService) Calculate() {
	network := createNetwork()
	coordinates := make([]types.Coordinate, 0)
	maxValue := 2.
	minValue := 0.
	step := maxValue / 10.
	for x := minValue; x <= maxValue; x += step {
		for y := minValue; y <= maxValue; y += step {
			z := upscale(network.Calculate([]float64{x, y})[0])
			coordinate := types.Coordinate{X: x, Y: y, Z: z}
			// logs.Logger.Info("Current coordinate", zap.Any("coord", coordinate))
			coordinates = append(coordinates, coordinate)
		}
	}
	persistCoordinatesToJSON(coordinates)
}

func createNetwork() *neural.Network {
	var network *neural.Network
	if configs.Opts.Input == "" {
		layers := []int{2}
		for i := 0; i < configs.Opts.HiddenLayers; i++ {
			layers = append(layers, configs.Opts.Nodes)
		}
		layers = append(layers, 1)
		network = neural.NewNetwork(2, layers)
		network.RandomizeSynapses()
	} else {
		network = persist.FromFile(configs.Opts.Input)
	}
	return network
}

func newRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func persistCoordinatesToJSON(coordinates []types.Coordinate) {
	resultJSONBytes, resultJSONBytesError := json.Marshal(coordinates)
	errors.HandleError(resultJSONBytesError, "Error marshaling results")
	files.WriteToJSON("coordinates.json", resultJSONBytes)
}

func downscale(number float64) float64 {
	return number / 10.
}

func upscale(number float64) float64 {
	return number * 10.
}
