package hasher

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CancelOrder struct {
	ContractAddr common.Address
	OrderHash    []byte
	UserAddr     common.Address
	Nonce        *big.Int
}

func (p *CancelOrder) Keccak256Hash() common.Hash {
	return crypto.Keccak256Hash(
		p.ContractAddr.Bytes(),
		p.OrderHash,
		p.UserAddr.Bytes(),
		common.LeftPadBytes(p.Nonce.Bytes(), 32),
	)
}
