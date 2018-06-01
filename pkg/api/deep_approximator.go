package api

// DeepApproximator describes methods to interact with the api
type DeepApproximator interface {
	Learn()

	Calculate()
}
