// main package is responsible for runing the VocationNations match-making algorithm
package main

import (
	"flag"
	"fmt"
	"github.com/vocationnations/algorithm/algorithm"
	"github.com/vocationnations/algorithm/config_parser"
)

const (
	// unsetYAML is the default value for config YAML file if it is not provided
	unsetYAML = "dummyYAML.yml"
)

// runs the main program which calculates the percent match for given parameters.
func main() {
	// candidate config
	yamlParam := "yaml"
	yamlValue := unsetYAML
	yamlUsage := "configuration file for the candidate"
	yaml := flag.String(yamlParam, yamlValue, yamlUsage)

	// parsing the flags has to be done after setting up all the flags
	flag.Parse()

	// if either of YAML files are not provided
	if *yaml == unsetYAML {
		panic("ERROR: YAML file is required.")
	}

	// parse the config
	cfg, err := config_parser.ParseConfig(*yaml)
	if err != nil {
		panic(fmt.Sprintf("ERROR: YAML config file was incorrect, err %v", err))
	}

	// run the algorithm
	res, err := algorithm.Run(cfg)
	if err != nil {
		panic(fmt.Sprintf("ERROR: The algorithm failed to run, erro %v", err))
	}

	// print the results
	fmt.Println(algorithm.Print(res))

}
