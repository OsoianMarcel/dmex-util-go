package wallet

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Signature []byte

// Decodes the signature in following values: R, S, V.
func (s Signature) Unmarshal() (common.Hash, common.Hash, uint8) {
	sR := s[:32]
	sS := s[32:64]
	sV := uint8(int(s[64])) + 27

	return common.BytesToHash(sR), common.BytesToHash(sS), sV
}

func (s Signature) Hex() string {
	return hexutil.Encode(s)
}

func (s Signature) Bytes() []byte {
	return s
}
