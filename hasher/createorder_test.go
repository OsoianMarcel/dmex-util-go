package hasher_test

import (
	"math/big"
	"testing"

	"github.com/OsoianMarcel/dmex-util-go/hasher"
	"github.com/ethereum/go-ethereum/common"
)

func Test_CreateOrder(t *testing.T) {
	createOrder := hasher.CreateOrder{
		ContractAddr:        common.HexToAddress("0xBE8f518B6515b16BCCb141496e2677F0e41faDDA"),
		FuturesContractHash: common.FromHex("0x8968e895fc5ca0159b20577972b0499a895a3cb8d467632213fbc170ed9857ba"),
		UserAddr:            common.HexToAddress("0x0A24BdE9D7618f9B745036FE0A7b57c7451ED724"),
		Amount:              big.NewInt(363077879),
		Price:               big.NewInt(12175000000),
		Side:                true,
		Nonce:               big.NewInt(4874),
		Leverage:            big.NewInt(1),
	}
	hash := createOrder.Keccak256Hash()

	expectedHash := "0x860aabb66e2b17865e150c49a920cb66572b44921ddca90611133767a653bd7e"

	if hash.Hex() != expectedHash {
		t.Error("create order hash failed")
	}
}
