package hasher_test

import (
	"math/big"
	"testing"

	"github.com/OsoianMarcel/dmx-sign/hasher"
	"github.com/ethereum/go-ethereum/common"
)

func Test_CreateOrder(t *testing.T) {
	createOrder := hasher.CreateOrder{
		ContractAddr:        common.HexToAddress("0xBE8f518B6515b16BCCb141496e2677F0e41faDDA"),
		FuturesContractHash: common.FromHex("0x4606c86aaf7a0bfc52f04010fc09543c4c021cb56459e4f9fafb845cb6fae730"),
		UserAddr:            common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Amount:              big.NewInt(1),
		Price:               big.NewInt(1),
		Side:                true,
		Nonce:               big.NewInt(1),
		Leverage:            big.NewInt(1),
	}
	hash := createOrder.Keccak256Hash()

	expectedHash := "0x44191d2a03fb3f724aeb71269af555eaba7a9f978f38a11949e5c839f4e869ec"

	if hash.Hex() != expectedHash {
		t.Error("create order hash failed")
	}
}
