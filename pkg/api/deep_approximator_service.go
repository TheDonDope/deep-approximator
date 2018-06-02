package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	"github.com/TheDonDope/deep-approximator/pkg/types"
	"github.com/TheDonDope/deep-approximator/pkg/util/configs"
	"github.com/TheDonDope/deep-approximator/pkg/util/errors"
	"github.com/TheDonDope/deep-approximator/pkg/util/files"
	"github.com/dariubs/percent"
)

// DeepApproximatorService implements the DeepApproximator interface
type DeepApproximatorService struct{}

// Learn starts the learning process of the neural network
func (impl DeepApproximatorService) Learn() {
	random := newRandom()
	network := createNetwork()
	for i := 0; i < configs.Opts.Rounds; i++ {
		x := random.Float64()
		y := random.Float64()
		// learn.Learn(network, []float64{x, y}, []float64{math.Sin(x + y)}, configs.Opts.Speed)
		learn.Learn(network, []float64{x, y}, []float64{downscale(x + y)}, configs.Opts.Speed)
		if i%(configs.Opts.Rounds/10) == 0 {
			log.Println(fmt.Sprintf("%v / %v (%v %%)", i, configs.Opts.Rounds, percent.PercentOf(i, configs.Opts.Rounds)))
		}
	}
	persist.ToFile(configs.Opts.Output, network)
}

// Calculate starts the calculation process of the nerual network
func (impl DeepApproximatorService) Calculate() {
	network := createNetwork()
	coordinates := make([]types.Coordinate, 0)
	for x := -1.; x <= 1.; x += 0.3 {
		for y := -1.; y <= 1.; y += 0.3 {
			z := upscale(network.Calculate([]float64{x, y})[0])
			coordinate := types.Coordinate{X: x, Y: y, Z: z}
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
	return number / 1000.
}

func upscale(number float64) float64 {
	return number * 1000.
}
