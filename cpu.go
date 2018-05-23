package chip8

import "time"

type CPU struct {

	//registers
	V [16]byte

	//address register
	I uint16

	//program counter
	PC uint16

	//stack
	Stack [16]uint16

	//stack pointer
	SP byte

	//delay timer
	DT byte

	//sound timer
	ST byte

	//Memory
	Memory [4096]byte

	Clock <-chan time.Time
}

func (c *CPU) Run() {
	for {
		select {
		case <-c.Clock:
			c.Execute()
		}
	}
}

func (c *CPU) Execute() {

	if c.DT > 0 {
		c.DT -= 1
	}
	if c.ST > 0 {
		c.ST -= 1
	}

	opCode := uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1])

	switch opCode & 0xF000 {
	case 0x0000:
		switch opCode {
		case 0x00E0:
			//clear screen
		case 0x00EE:
			//return from subrouting
		}
	case 0x1000:
		//jump to address NNN (0x1NNN)
	case 0x2000:
		//call subroutine NNN (0x2NNN)
	case 0x3000:
		// skip next instruction if V[X]==NN (0x3XNN)
	case 0x4000:
		// skip next instruction if V[X]!=NN (0x3XNN)
	case 0x5000:
		// skip next instruction if V[X]==V[Y] (0x3XY0)
	case 0x6000:
		//set V[X] to NN (0x6XNN)
	case 0x7000:
		//set V[X] to V[X}+NN (0x7XNN)
	case 0x8000:
		switch opCode & 0x000F {
		case 0x0001:
			//set V[X] to V[X] OR V[Y] (0x8XY1)
		case 0x0002:
			//set V[X] to V[X] AND V[Y] (0x8XY2)
		case 0x0003:
			//set V[X] to V[X] XOR V[Y] (0x8XY3)
		case 0x0004:
			//set V[X] to V[X] + V[Y] (0x8XY4), set V[F] to 1 if carry
		case 0x0005:
			//set V[X] to V[X] - V[Y] (0x8XY5), set V[F] to 1 if no borrow
		case 0x0006:
			//set V[X] to V[Y] >> 1 (0x8XY6), set V[F] to V[Y] LSB before shift
		case 0x0007:
			//set V[X] to V[Y] - V[Y] (0x8XY7)
		case 0x0006:
			//set V[X] (and V[Y]) to V[Y] << 1 (0x8XY6)
		}
	}
}
