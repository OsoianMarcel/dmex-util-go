package hasher

import "testing"

func Test_BoolToBytes_True(t *testing.T) {
	b := boolToBytes(true)

	if len(b) != 1 {
		t.Error("wrong number of bytes")
	}

	if b[0] != 1 {
		t.Error("bad conversion")
	}
}

func Test_BoolToBytes_False(t *testing.T) {
	b := boolToBytes(false)

	if len(b) != 1 {
		t.Error("wrong number of bytes")
	}

	if b[0] != 0 {
		t.Error("bad conversion")
	}
}
