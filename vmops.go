// Package vmops is an interpreter for the vmops output from libfsm.
//
// For more information, please see: https://github.com/katef/libfsm
package vmops

import "unsafe"

type Op byte

const (
	OpEOF Op = iota
	OpLT
	OpLE
	OpEQ
	OpNE
	OpGE
	OpGT
	OpALWAYS
)

type Action byte

const (
	ActionRET Action = iota
	ActionGOTO
)

type Opcode struct {
	Op     Op
	C      byte
	Action Action
	// padding byte here
	Arg int32
}

// VM is a set of opcodes encoding a DFA matching a regular expression.
type VM []Opcode

// Match runs the vm against the input string and returns result.
func (vm VM) MatchString(input string) int {
	s := unsafe.Slice(unsafe.StringData(input), len(input))
	return vm.Match(s)
}

// Match runs the vm against the input byte slice and returns result.
func (vm VM) Match(input []byte) int {
	var ip int32
	var idx = ^uint(0)
	var c byte
	ops := ([]Opcode)(vm)

	for {
		var ok bool
		switch ops[ip].Op {
		case OpEOF:
			// Are we at the end of the input?  If so, set `ok` we carry out the action.
			// Otherwise, read the next input byte.
			if idx++; idx < uint(len(input)) {
				c = input[idx]
			} else {
				ok = true
			}
		case OpLT:
			ok = c < ops[ip].C
		case OpLE:
			ok = c <= ops[ip].C
		case OpEQ:
			ok = c == ops[ip].C
		case OpNE:
			ok = c != ops[ip].C
		case OpGE:
			ok = c >= ops[ip].C
		case OpGT:
			ok = c > ops[ip].C
		case OpALWAYS:
			ok = true
		}
		if ok {
			if ops[ip].Action == ActionRET {
				return int(ops[ip].Arg)
			}
			ip = ops[ip].Arg
			continue
		}
		ip++
	}
}
