package vm

import "math/bits"

type ALU struct {
	Acc   byte
	Flags struct {
		ZF bool
		PF bool
		CF bool
		BF bool
	}
}

func (alu *ALU) computeFlags() {
	alu.Flags.ZF = (alu.Acc == 0)
	alu.Flags.PF = (bits.OnesCount8(alu.Acc) % 2) == 0
}

func (alu *ALU) Add(x, y byte) {
	sum16 := uint16(x) + uint16(y)
	alu.Acc = byte(sum16)
	alu.Flags.CF = (sum16 >> 8) > 0
	alu.computeFlags()
}

func (alu *ALU) Sub(x, y byte) {
	sum16 := uint16(x) - uint16(y)
	alu.Acc = byte(sum16)
	alu.Flags.CF = (sum16 >> 8) > 0
	alu.computeFlags()
}
