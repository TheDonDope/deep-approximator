package api

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	"github.com/TheDonDope/deep-approximator/pkg/util/configs"
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
		learn.Learn(network, []float64{x, y}, []float64{math.Sin(x + y)}, configs.Opts.Speed)
		if i%(configs.Opts.Rounds/10) == 0 {
			log.Println(fmt.Sprintf("%v / %v (%v %%)", i, configs.Opts.Rounds, percent.PercentOf(i, configs.Opts.Rounds)))
		}
	}
	persist.ToFile(configs.Opts.Output, network)
}

// Calculate starts the calculation process of the nerual network
func (impl DeepApproximatorService) Calculate() {
	random := newRandom()
	network := createNetwork()
	x := random.Float64()
	y := random.Float64()
	result := network.Calculate([]float64{x, y})
	idealResult := math.Sin(x + y)
	fmt.Println(fmt.Sprintf("%v should be: %v", result[0], idealResult))
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
