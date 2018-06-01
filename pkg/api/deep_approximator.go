package api

// DeepApproximator describes methods to interact with the api
type DeepApproximator interface {

	// Learn starts the learning process of the neural network
	Learn()

	// Calculate starts the calculation process of the nerual network
	Calculate()
}
