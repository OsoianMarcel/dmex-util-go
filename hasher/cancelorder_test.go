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
		OrderHash:    common.FromHex("0xe63b0f878c1da07c0f21ddcc0d04f7f489b5dcec16307260c596db597e3263d1"),
		UserAddr:     common.HexToAddress("0x0A24BdE9D7618f9B745036FE0A7b57c7451ED724"),
		Nonce:        big.NewInt(10086),
	}
	hash := cancelOrder.Keccak256Hash()

	expectedHash := "0xb710a6de64a93f8e7389dbbc73724a8b252b66a4cf45bdacae449da9568c66e0"

	if hash.Hex() != expectedHash {
		t.Error("cancel order hash failed")
	}
}
