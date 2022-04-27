package vm

import (
	"fmt"
	"log"
	"os"
)

var Log *log.Logger = log.New(os.Stdout, "", 0)

const MEM_SIZE int = 1024

const (
	OP_NOP  byte = iota
	OP_HALT byte = 0xFF
)

const (
	OP_ADD byte = 0xA0 + iota
	OP_SUB
)

const (
	OP_STORED byte = 0xC0 + iota
	OP_STOREI
	OP_LOADI
	OP_MOVD
	OP_MOVI
	OP_PRINT
)

const (
	OP_JUMP byte = 0xE0 + iota
	OP_JZ
	OP_JNZ
)

type Word uint32

type Machine struct {
	stopped bool
	mem     *Memory
	alu     *ALU

	//Registers
	ip                 uint
	currentInstruction Word

	reg0 byte
	reg1 byte
}

func New() *Machine {
	return &Machine{
		mem: NewMemory(MEM_SIZE),
		alu: &ALU{},
	}
}

func (m *Machine) LoadProgram(prog []Word) error {
	return m.mem.Set(prog)
}

func (m *Machine) halt() {
	Log.Printf("execution halted at %x\n", m.ip)
	m.stopped = true
}

func (m *Machine) start() {
	Log.Printf("execution started at %x\n", m.ip)
	m.stopped = false
}

func (m *Machine) Run() {
	m.start()

	var err error

	for {
		if m.stopped {
			return
		}

		err = m.fetch()
		if err != nil {
			Log.Println(err)
			break
		}
		err = m.exec()
		if err != nil {
			Log.Println(err)
			break
		}

		m.ip++

	}

	m.halt()
}

func (m *Machine) fetch() error {
	return m.load(m.ip, &m.currentInstruction)
}

func (m *Machine) load(addr uint, reg *Word) error {
	v, err := m.mem.LoadW(addr)
	if err != nil {
		return err
	}

	*reg = v
	return nil
}

func (m *Machine) exec() error {
	//0x{OpCode}{b2}{b3}{b4}
	opCode := byte(m.currentInstruction >> 24)
	b2 := byte(m.currentInstruction >> 16)
	b3 := byte(m.currentInstruction >> 8)
	b4 := byte(m.currentInstruction)

	switch opCode {
	default:
		return fmt.Errorf("unknown instruction '%x'", m.currentInstruction)
	case OP_NOP:
	//do literally nothing
	case OP_HALT: //halt
		m.halt()

	case OP_PRINT:
		r, err := m.findReg(b2)
		if err != nil {
			return err
		}
		fmt.Printf("0x%x = %d\n", *r, *r)
	case OP_ADD:
		r1, err := m.findReg(b2)
		if err != nil {
			return err
		}
		r2, err := m.findReg(b3)
		if err != nil {
			return err
		}
		m.alu.Add(*r1, *r2)
	case OP_SUB:
		r1, err := m.findReg(b2)
		if err != nil {
			return err
		}
		r2, err := m.findReg(b3)
		if err != nil {
			return err
		}
		m.alu.Sub(*r1, *r2)
	case OP_JUMP:
		m.jump(Word(b4))
	case OP_JZ:
		if m.alu.Flags.ZF {
			m.jump(Word(b4))
		}
	case OP_JNZ:
		if !m.alu.Flags.ZF {
			m.jump(Word(b4))
		}
	case OP_MOVD:
		r, err := m.findReg(b3)
		if err != nil {
			return err
		}
		*r = b2
	case OP_MOVI:
		r1, err := m.findReg(b2)
		if err != nil {
			return err
		}
		r2, err := m.findReg(b3)
		if err != nil {
			return err
		}
		*r2 = *r1
	case OP_STORED:
		m.mem.StoreW(uint(b3), Word(b2))
	case OP_STOREI:
		r, err := m.findReg(b2)
		if err != nil {
			return err
		}
		m.mem.StoreW(uint(b3), Word(*r))
	case OP_LOADI:
		r, err := m.findReg(b2)
		if err != nil {
			return err
		}
		data, err := m.mem.LoadW(uint(b3))
		if err != nil {
			return err
		}
		*r = byte(data)
	}

	return nil
}

func (m *Machine) jump(addr Word) {
	m.ip = uint(addr) - 1
}

func (m *Machine) findReg(regCode byte) (*byte, error) {
	switch regCode {
	case 0xA:
		return &m.alu.Acc, nil
	case 0:
		return &m.reg0, nil
	case 1:
		return &m.reg1, nil
	}

	return nil, fmt.Errorf("unknown reg '%x'", regCode)
}
