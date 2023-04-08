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

type errorType struct {
	errMsg string
}

// can these be interfaces https://go.dev/tour/methods/9
var Props propsType
var grid gridType

var envVars = map[string]map[string]string{
	"blockSize": map[string]string{
		"name":  "GODOKU_BLOCK_SIZE",
		"type":  "int",
		"props": "BlockSize",
	},
	"difficulty": map[string]string{
		"name":  "GODOKU_DIFFICULTY",
		"type":  "int",
		"props": "Difficulty",
	},
	"portNumber": map[string]string{
		"name":  "GODOKU_PORT",
		"type":  "string",
		"props": "PortNumber",
	},
}

/*
	var envVars = []map[string]string{
		map[string]string{
			"name":  "GODOKU_BLOCK_SIZE",
			"type":  "int",
			"props": "BlockSize",
		},
		map[string]string{
			"name":  "GODOKU_DIFFICULTY",
			"type":  "int",
			"props": "Difficulty",
		},
		map[string]string{
			"name":  "GODOKU_PORT",
			"type":  "string",
			"props": "PortNumber",
		},
	}
*/
var errors = map[string]string{
	"missingEnv": ": not found in the environment",
}

var gdkError errorType
