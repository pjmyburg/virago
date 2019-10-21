package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

const (
	// Region is the default AWS region to use in all API responses
	Region = "us-west-2"
	// AccountID is the default AWS account ID to use in all API responses
	AccountID = "000000000000"
)

// Config represents a config yaml file loaded into memory
type Config struct {
	Queues []Queue
	Topics []Topic
}

// Queue represents configuration for a single SQS queue
type Queue struct {
	Name string
}

// Topic represents configuration for a single SNS topic
type Topic struct {
	Name string
}

// LoadYaml loads the specified config file into memory
func LoadYaml(fileName string) (*Config, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = yaml.Unmarshal(file, conf)

	return conf, nil
}
