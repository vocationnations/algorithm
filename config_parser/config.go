package config_parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	Config struct {
		Candidate CandidateConfig `yaml:"candidate"`
		Employer  EmployerConfig  `yaml:"employer"`
		Worth     MatchingWorth   `yaml:"worth"`
	}

	MatchingWorth struct {
		Skills  float64 `yaml:"skills"`
		Culture float64 `yaml:"culture"`
	}

	Skill struct {
		Name  string `yaml:"name"`
		Value float64    `yaml:"value"`
	}

	Domain struct {
		Hierarchy float64 `yaml:"hierarchy"`
		Market    float64 `yaml:"market"`
		Adhocracy float64 `yaml:"adhocracy"`
		Clan      float64 `yaml:"clan"`
	}

	CandidateConfig struct {
		Skills  []Skill  `yaml:"skills"`
		Culture []Domain `yaml:"culture"`
	}

	EmployerConfig struct {
		Skills  []Skill  `yaml:"skills"`
		Culture []Domain `yaml:"culture"`
	}
)

// ParseConfig takes in the YAML file and returns the config struct
func ParseConfig(yamlFilePath string) (*Config, error) {
	config := &Config{}

	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Cannot read file %s, err: %v", yamlFilePath, err)
	}

	// try to unmarshall the YAML file.
	if err = yaml.Unmarshal(yamlFile,&config); err != nil {
		return nil, fmt.Errorf("ERROR: Cannot unmarshall the YAML file, err %v",err)
	}

	// validate the YAML file
	if err = validateConfig(config); err != nil {
		return nil, fmt.Errorf("ERROR: The YAML file is not valid, err %v", err)
	}

	return config, nil
}

func validateConfig(c *Config) error {
	// TODO: work on validating the YAML file
	return nil
}