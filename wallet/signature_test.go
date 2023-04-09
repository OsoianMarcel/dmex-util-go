package wallet_test

import (
	"testing"

	"github.com/OsoianMarcel/dmx-sign/wallet"
	"github.com/ethereum/go-ethereum/common"
)

func Test_Signature_Unmarshal(t *testing.T) {
	signHash := "0x4b0918370621a8e5ba3291c31db3c9295ed73d1813f922240aaefac7eaa5a8a021e12ce69a3126e201392fcb388971a22756b025e0311a65e41b7daab681d48301"

	var signature wallet.Signature = common.FromHex(signHash)

	r, s, v := signature.Unmarshal()

	if r.Hex() != "0x4b0918370621a8e5ba3291c31db3c9295ed73d1813f922240aaefac7eaa5a8a0" {
		t.Error("wrong r value")
		return
	}

	if s.Hex() != "0x21e12ce69a3126e201392fcb388971a22756b025e0311a65e41b7daab681d483" {
		t.Error("wrong r value")
		return
	}

	if v != 28 {
		t.Error("wrong v value")
	}
}
