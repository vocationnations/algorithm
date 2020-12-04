package config_parser

import (
	"fmt"
	"github.com/vocationnations/algorithm/config_parser/constants"
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
		Name  string  `yaml:"name"`
		Value float64 `yaml:"value"`
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
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, fmt.Errorf("ERROR: Cannot unmarshall the YAML file, err %v", err)
	}

	// validate the YAML file
	if err = validateConfig(config); err != nil {
		return nil, fmt.Errorf("ERROR: The YAML file is not valid, err %v", err)
	}

	return config, nil
}

func validateConfig(c *Config) error {

	// worth should add up to 100%
	if c.Worth.Skills+c.Worth.Culture != 100 {
		return fmt.Errorf("ERROR: The worth should add up to 100%")
	}

	// the total number of categories for employer and candidate culture is equal
	//  to TOTAL_CATEGORIES
	if len(c.Candidate.Culture) != constants.TotalCategories &&
		len(c.Employer.Culture) != constants.TotalCategories {
		return fmt.Errorf("ERROR: The total number of categories in culture for candidate and employer "+
			"should be equal to %d", constants.TotalCategories)
	}

	// TODO: check to ensure that all the skill values are within 0-100

	// TODO: Check to ensure that domain values in very category sum up to 100% for both employer and candidates.

	return nil
}
