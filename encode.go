package vmops

import (
	"encoding/binary"
	"errors"
	"unsafe"
)

func (vm VM) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 0, len(vm)*8)
	var u32 [4]byte
	for _, op := range vm {
		data = append(data, byte(op.Op), byte(op.C), byte(op.Action), 0)
		binary.LittleEndian.PutUint32(u32[:], uint32(op.Arg))
		data = append(data, u32[:]...)
	}

	return data, nil
}

var ErrBadLength = errors.New("vmops: bad data length")

var ErrCorrupt = errors.New("vmops: bad data")

func (vm *VM) UnmarshalBinary(data []byte) error {
	if len(data)%8 != 0 {
		return ErrBadLength
	}

	ops := make([]Opcode, 0, len(data)/8)
	maxArg := uint32(cap(ops))

	var b []byte
	for len(data) > 0 {
		b, data = data[:8], data[8:]

		if b[0] > byte(OpALWAYS) {
			return ErrCorrupt
		}
		if b[2] > byte(ActionGOTO) {
			return ErrCorrupt
		}
		// padding byte
		if b[3] != 0 {
			return ErrCorrupt
		}

		arg := binary.LittleEndian.Uint32(b[4:])
		if b[2] == byte(ActionGOTO) && arg > maxArg {
			return ErrCorrupt
		}
		ops = append(ops, Opcode{
			Op:     Op(b[0]),
			C:      b[1],
			Action: Action(b[2]),
			Arg:    int32(binary.LittleEndian.Uint32(b[4:])),
		})
	}
	*vm = VM(ops)
	return nil
}

func (vm *VM) UnmarshalBinaryUnsafe(data []byte) error {
	if len(data)%8 != 0 {
		return ErrBadLength
	}
	*vm = VM(unsafe.Slice((*Opcode)(unsafe.Pointer(&data[0])), uintptr(len(data))/unsafe.Sizeof(Opcode{})))
	return nil
}

func (vm VM) MarshalBinaryUnsafe() (data []byte, err error) {
	ops := ([]Opcode)(vm)
	return unsafe.Slice((*byte)(unsafe.Pointer(&(ops[0]))), uintptr(len(ops))*unsafe.Sizeof(Opcode{})), nil
}
