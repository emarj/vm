package main

import (
	"ronche.se/vm"
)

func main() {

	m := vm.New()

	prog := []vm.Word{
		0xc3030A00, // MOV $3,ACC
		0xc3010100, // MOV $1,R1
		0xC50A0000, // PRINT ACC
		0xA10A0100, // SUB ACC,R1 // ACC = ACC -1
		0xE2000002, // JNZ 0x02
		//0xE0000006, // JUMP $
		0xFF000000, // HALT
	}

	/*0xc3060000, // MOV $3,R0
	0xc3020100, // MOV $1,R1
	0x3E000000, // PRINT
	0xA1000100, // SUB R0,R1 // ACC = R0 -1 // R0 = R0 -1
	0xc40A0000, // MOV ACC,R0
	0xE2000002, // JNZ 0x02
	//0xE0000006, // JUMP $
	0xFF000000, // HALT*/

	m.LoadProgram(prog)

	m.Run()
}
