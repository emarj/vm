package vm

import "testing"

func TestParseOpCode(t *testing.T) {
	tests := []struct {
		input  string
		output byte
	}{
		{input: "ADD", output: 0xA0},
		{input: " ADD ", output: 0xA0},
		{input: "\t AdD ", output: 0xA0},
		{input: "\n ADD ", output: 0xA0},
		{input: "aDd", output: 0xA0},
	}

	var got InstructionDesc
	var err error
	for _, test := range tests {
		got, err = parseOpCode(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if got.OpCode != test.output {
			t.Fatalf("got %x expected %x", got, test.output)
		}
	}
}
func TestAssembler(t *testing.T) {
	tests := []struct {
		input  string
		output Word
	}{{input: "ADD $4 $5", output: 0xA0040500}}

	var got Word
	var err error
	for _, test := range tests {
		got, err = AssembleInstruction(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.output {
			t.Fatalf("got %x expected %x", got, test.output)
		}
	}
}
