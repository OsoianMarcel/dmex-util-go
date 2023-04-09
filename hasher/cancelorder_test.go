package hasher_test

import (
	"math/big"
	"testing"

	"github.com/OsoianMarcel/dmex-util-go/hasher"
	"github.com/ethereum/go-ethereum/common"
)

func Test_CancelOrder(t *testing.T) {
	cancelOrder := hasher.CancelOrder{
		ContractAddr: common.HexToAddress("0xBE8f518B6515b16BCCb141496e2677F0e41faDDA"),
		OrderHash:    common.FromHex("0x4606c86aaf7a0bfc52f04010fc09543c4c021cb56459e4f9fafb845cb6fae730"),
		UserAddr:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Nonce:        big.NewInt(1),
	}
	hash := cancelOrder.Keccak256Hash()

	expectedHash := "0x2e0aee7f64a547734a84b4e28f66361818302d9ab22e74761431c815e5a887c5"

	if hash.Hex() != expectedHash {
		t.Error("cancel order hash failed")
	}
}
