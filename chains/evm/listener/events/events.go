// The Licensed Work is (c) 2023 Sygma
// SPDX-License-Identifier: LGPL-3.0-only

package events

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	StateRootSubmittedSig = "StateRootSubmitted(uint8,uint256,bytes32)"
	DepositSig            = "Deposit(uint8,uint8,bytes32,uint64,address,bytes)"
)

type StateRootSubmitted struct {
	// ID of chain from which the state root is from
	SourceDomainID uint8
	// Finalized beacon slot belonging to the state root
	Slot *big.Int
	// Execution state root
	StateRoot [32]byte
}

type Deposit struct {
	// ID of chain deposit will be bridged to
	DestinationDomainID uint8
	// SecurityModel that defines the destination verifiers
	SecurityModel uint8
	// ResourceID used to find address of handler to be used for deposit
	ResourceID [32]byte
	// Nonce of deposit
	DepositNonce uint64
	// Address of sender (msg.sender: user)
	SenderAddress common.Address
	// Deposit data
	Data []byte
}
