package api

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	"github.com/TheDonDope/deep-approximator/pkg/util/configs"
)

// DeepApproximatorService implements the DeepApproximator interface
type DeepApproximatorService struct {
	network *neural.Network
	random  *rand.Rand
}

// Learn starts the learning process of the neural network
func (impl DeepApproximatorService) Learn() {
	for i := int64(0); i < configs.Opts.Rounds; i++ {
		x := impl.random.Float64()
		y := impl.random.Float64()
		learn.Learn(impl.network, []float64{x, y}, []float64{math.Sin(x + y)}, configs.Opts.Speed)
		fmt.Println(fmt.Sprintf("%v / %v (%v %%)", i, configs.Opts.Rounds, i/configs.Opts.Rounds*100))
	}
	persist.ToFile(configs.Opts.Output, impl.network)
}

// Calculate starts the calculation process of the nerual network
func (impl DeepApproximatorService) Calculate() {
	x := impl.random.Float64()
	y := impl.random.Float64()
	result := impl.network.Calculate([]float64{x, y})
	idealResult := math.Sin(x + y)
	fmt.Println(fmt.Sprintf("%v should be: %v", result[0], idealResult))
}

// InitNetwork initialises the neural network
func (impl DeepApproximatorService) InitNetwork() {
	if configs.Opts.Input == "" {
		impl.network = neural.NewNetwork(2, []int{2, 10, 1})
	} else {
		impl.network = persist.FromFile(configs.Opts.Input)
	}
	source := rand.NewSource(time.Now().UnixNano())
	impl.random = rand.New(source)
}
