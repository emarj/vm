package vm

import (
	"errors"
	"fmt"
)

type Memory struct {
	mem []Word
}

func NewMemory(size int) *Memory {

	if size%4 != 0 {
		size += 4 - size%4
	}

	return &Memory{
		mem: make([]Word, size),
	}
}

func (m *Memory) addrOK(addr uint) bool {
	return int(addr) >= len(m.mem)
}

func (m *Memory) Set(data []Word) error {
	if len(data) > len(m.mem) {
		return fmt.Errorf("data size exceeds memory size")
	}
	nb := copy(m.mem, data)
	Log.Printf("loaded %d words (%d bytes) in memory\n", nb, nb*4)

	return nil
}

func (m *Memory) Dump() []Word {
	var snapshot []Word
	copy(snapshot, m.mem)
	return snapshot
}

func (m *Memory) LoadW(addr uint) (Word, error) {
	if m.addrOK(addr) {
		return 0, errors.New("load: addr is out of range")
	}
	return m.mem[addr], nil
}

func (m *Memory) StoreW(addr uint, v Word) error {
	if m.addrOK(addr) {
		return errors.New("store: addr is out of range")
	}

	m.mem[addr] = v
	return nil
}
