package hasher

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CreateOrder struct {
	ContractAddr        common.Address
	FuturesContractHash []byte
	UserAddr            common.Address
	Amount              *big.Int
	Price               *big.Int
	Side                bool
	Nonce               *big.Int
	Leverage            *big.Int
}

func (p *CreateOrder) Keccak256Hash() common.Hash {
	return crypto.Keccak256Hash(
		p.ContractAddr.Bytes(),
		p.FuturesContractHash,
		p.UserAddr.Bytes(),
		common.LeftPadBytes(p.Amount.Bytes(), 32),
		common.LeftPadBytes(p.Price.Bytes(), 32),
		boolToBytes(p.Side),
		common.LeftPadBytes(p.Nonce.Bytes(), 32),
		common.LeftPadBytes(p.Leverage.Bytes(), 32),
	)
}
