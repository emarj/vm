package vm

import (
	"fmt"
	"strings"
)

type ArgDesc struct {
	Size uint
}

type InstructionDesc struct {
	Name   string
	OpCode byte
	Args   []ArgDesc
}

var instrMap map[string]InstructionDesc = map[string]InstructionDesc{
	"ADD": {
		Name:   "ADD",
		OpCode: OP_ADD,
		Args: []ArgDesc{
			{
				Size: 1,
			},
			{
				Size: 1,
			},
		},
	},
}

func parseOpCode(opCodeStr string) (InstructionDesc, error) {
	opCodeStr = strings.TrimSpace(opCodeStr)
	opCodeStr = strings.ToUpper(opCodeStr)

	opCode, exists := instrMap[opCodeStr]
	if !exists {
		return InstructionDesc{}, fmt.Errorf("unknown instruction identifier '%s'", opCodeStr)
	}
	return opCode, nil

}

/*
func parseArg(argStr string) (InstructionDesc, error) {
	opCodeStr = strings.TrimSpace(opCodeStr)
	opCodeStr = strings.ToUpper(opCodeStr)

	opCode, exists := instrMap[opCodeStr]
	if !exists {
		return InstructionDesc{}, fmt.Errorf("unknown instruction identifier '%s'", opCodeStr)
	}
	return opCode, nil

}*/

func AssembleInstruction(instr string) (Word, error) {

	/*tkns := strings.Split(instr, " ")

	tkns[0]*/

	return 0, nil
}
