// The Licensed Work is (c) 2023 Sygma
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sygmaprotocol/sygma-inclusion-prover/config"
)

type Spec string

const (
	MainnetSpec Spec = "mainnet"
	GnosisSpec  Spec = "gnosis"
)

type EVMConfig struct {
	config.BaseNetworkConfig
	BeaconEndpoint        string `split_words:"true"`
	ArchiveBeaconEndpoint string `split_words:"true"`
	Router                string
	Executor              string
	Hashi                 string
	Yaho                  string
	StartBlock            uint64   `split_words:"true"`
	StateRootAddresses    []string `split_words:"true"`
	SlotIndex             uint8    `required:"true" split_words:"true"`
	MaxGasPrice           int64    `default:"500000000000" split_words:"true"`
	GasMultiplier         float64  `default:"1" split_words:"true"`
	GasIncreasePercentage int64    `default:"15" split_words:"true"`
	BlockConfirmations    int64    `default:"1" split_words:"true"`
	BlockInterval         int64    `default:"5" split_words:"true"`
	BlockRetryInterval    uint64   `default:"5" split_words:"true"`
	FreshStart            bool     `default:"false" split_words:"true"`
	Latest                bool     `default:"false" split_words:"true"`
	GenericResources      []string `default:"0000000000000000000000000000000000000000000000000000000000000500" split_words:"true"`
	Spec                  Spec     `default:"mainnet"`
}

// LoadEVMConfig loads EVM config from the environment and validates the fields
func LoadEVMConfig(domainID uint8) (*EVMConfig, error) {
	var c EVMConfig
	err := envconfig.Process(fmt.Sprintf("%s_DOMAINS_%d", config.PREFIX, domainID), &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
