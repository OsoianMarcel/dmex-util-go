package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
}

func New(privateKeyHash string) (*Wallet, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHash)
	if err != nil {
		return nil, err
	}

	return &Wallet{privateKey}, nil
}

func (w *Wallet) GetAddress() (common.Address, error) {
	publicKey := w.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func (w *Wallet) Sign(digestHash []byte) (Signature, error) {
	return crypto.Sign(digestHash, w.privateKey)
}

func (w *Wallet) SignWithPrefix(digestHash []byte) (Signature, error) {
	prefixStr := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(digestHash))

	prefixedHash := crypto.Keccak256(
		[]byte(prefixStr),
		digestHash,
	)

	sign, err := w.Sign(prefixedHash)

	return sign, err
}

func (w *Wallet) SignHasher(hasher Keccak256Hasher) (Signature, error) {
	hash := hasher.Keccak256Hash()

	return w.Sign(hash.Bytes())
}

func (w *Wallet) SignHasherWithPrefix(hasher Keccak256Hasher) (Signature, error) {
	hash := hasher.Keccak256Hash()

	return w.SignWithPrefix(hash.Bytes())
}
