package wallet

import "github.com/ethereum/go-ethereum/common"

type Keccak256Hasher interface {
	Keccak256Hash() common.Hash
}
