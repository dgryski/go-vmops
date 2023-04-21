package vmops

import (
	"bytes"
	"encoding"
	"reflect"
	"testing"
	"unsafe"
)

// Compile time check that Opcode is 8 bytes
var _ = [1]bool{}[unsafe.Sizeof(Opcode{})-8]

var _ encoding.BinaryUnmarshaler = (*VM)(unsafe.Pointer(nil))

var _ encoding.BinaryMarshaler = (*VM)(unsafe.Pointer(nil))

func TestRoundtrip(t *testing.T) {
	b, err := abcVM.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary error: %v", err)
	}

	var abc2 VM

	if err := abc2.UnmarshalBinary(b); err != nil {
		t.Errorf("UnmarshalBinary error: %v", err)
	}

	if !reflect.DeepEqual(abcVM, abc2) {
		t.Errorf("Roundtrip mismatch")
	}
}

func TestRoundtripUnsafe(t *testing.T) {
	b, _ := abcVM.MarshalBinary()
	bu, err := abcVM.MarshalBinaryUnsafe()
	if err != nil {
		t.Errorf("MarshalBinary error: %v", err)
	}

	if !bytes.Equal(b, bu) {
		t.Errorf("MarshalBinarUnsafe() mismatch\n")
	}

	var abc2 VM
	if err := abc2.UnmarshalBinaryUnsafe(bu); err != nil {
		t.Errorf("UnmarshalBinary error: %v", err)
	}

	if !reflect.DeepEqual(abcVM, abc2) {
		t.Errorf("RoundtripUnsafe mismatch")
	}
}
