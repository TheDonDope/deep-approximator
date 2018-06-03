package configs

// Opts are the program options, configurable by command line argument
var Opts struct {
	Learn bool `short:"l" long:"learn" description:"Initiate or continue a learning process. If an input file is given via the -i/--input-file argument the network will continue to learn from a previous trained model. Else a new learning process will be started."`

	Calculate bool `short:"c" long:"calculate" description:"Execute the trained neural network with the given input parameters. You need to either provide an input file via the -i/--input-file argument or use the -l/--learn parameter to first start a learning process and afterwards do a calculation"`

	Input string `short:"i" long:"input-file" description:"The name of the pretrained model (JSON file) to use."`

	Output string `short:"o" long:"output-file" description:"The name of the JSON file to persist the trained model to." default:"results.json"`

	HiddenLayers int `short:"h" long:"hidden-layers" description:"The number of hidden layers to use. (default: 1)" default:"1"`

	Nodes int `short:"n" long:"nodes" description:"The number of nodes per hidden layer. (default: 10)" default:"10"`

	Rounds int `short:"r" long:"rounds" description:"The number of learning iterations (default: 1.000.000)" default:"1000000"`

	Speed float64 `short:"s" long:"speed" description:"The factor to be used to alter the weights of all nodes. (default: 0.1)" default:"0.1"`

	Expression string `short:"e" long:"expression" description:"The expression learned by the neural network. (default: x+y)" default:"0.25*x*x + 0.5*y*y"`
}
