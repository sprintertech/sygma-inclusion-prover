// The Licensed Work is (c) 2023 Sygma
// SPDX-License-Identifier: LGPL-3.0-only

package config

import "github.com/kelseyhightower/envconfig"

const PREFIX = "INCLUSION_PROVER"

type Config struct {
	Observability *Observability   `env_config:"observability"`
	Store         *Store           `env_config:"store"`
	Domains       map[uint8]string `required:"true"`
	ChainIDS      map[uint8]uint64 `required:"true"`
}

type Observability struct {
	LogLevel   string `default:"debug" split_words:"true"`
	LogFile    string `default:"out.log" split_words:"true"`
	HealthPort uint16 `default:"9001" split_words:"true"`
}

type Store struct {
	Path string `default:"./lvldbdata"`
}

// LoadConfig loads config from the environment and validates the fields
func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process(PREFIX, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
