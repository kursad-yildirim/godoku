package gdk

var valueArray []int
var indexArray []int

type propsType struct {
	BlockSize      int    `json:"blockSize"`
	GridSize       int    `json:"gridSize"`
	Difficulty     int    `json:"difficulty"`
	GridPopulation int    `json:"gridPopulation"`
	PortNumber     string `json:"portNumber"`
}

type cellType struct {
	value   int
	display bool
}

type gridType struct {
	members [][]cellType
}

// can these be interfaces https://go.dev/tour/methods/9
var Props propsType
var grid gridType

var envVars = map[string]string{
	"blockSize":  "GODOKU_BLOCK_SIZE",
	"difficulty": "GODOKU_DIFFICULTY",
	"portNumber": "GODOKU_PORT",
}

var errors = map[string]string{
	"missingEnv": ": not found in the environment",
}
