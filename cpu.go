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
		default:
			// ignore RCA1802 functions
		}
	case 0x1000:
		//jump to address NNN (0x1NNN)
	case 0x2000:
		//call subroutine NNN (0x2NNN)
	case 0x3000:
		// skip next instruction if V[X]==NN (0x3XNN)
	case 0x4000:
		// skip next instruction if V[X]!=NN (0x4XNN)
	case 0x5000:
		// skip next instruction if V[X]==V[Y] (0x5XY0)
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
			//set V[X] to V[X] + V[Y] (0x8XY4), set V[F] to 1 if carry, otherwise 0
		case 0x0005:
			//set V[X] to V[X] - V[Y] (0x8XY5), set V[F] to 1 if no borrow, otherwise 0
		case 0x0006:
			//set V[X] to V[Y] >> 1 (0x8XY6), set V[F] to V[Y] LSB before shift
		case 0x0007:
			//set V[X] to V[Y] - V[Y] (0x8XY7), set V[F] to 1 if no borrow, otherwise 0
		case 0x000E:
			//set V[X] (and V[Y]) to V[Y] << 1 (0x8XYE), set V[F] to V[Y] MSB before shift
		}
	case 0x9000:
		//skip next instruction if V[X]!=V[Y] (0x9XY0)
	case 0xA000:
		//set I to NNN (0xANNN)
	case 0xB000:
		//jump to addr V[0]+NNN: set PC to V[0] +NNN (0xBNNN)
	case 0xC000:
		//set V[x] to R+NN where R = random number between 0 and 255(0xCXNN)
	case 0xD000:
		//draw sprite at position (V[X],V[Y]) with width 8, heigh N.
		//sprite bits located at Memory[I] in rows of 8 (0xDXYN)
		//V[F] is set to 1 if pixels are flipped from 1 to 0, otherwise 0
	case 0xE000:
		switch opCode & 0x00FF {
		case 0x009E:
		//skip next instruction if key pressed == v[X] (0xEX9E)
		case 0x00A1:
			//skip next instruction if key pressed != v[X] (0xEXA1)
		}
	case 0xF000:
		switch opCode & 0x00FF {
		case 0x0007:
			//set V[X] to DT (0xFX07)
		case 0x000A:
			//set v[X] to key pressed (0xFX0A)
		case 0x0015:
			//set DT to V[X] (0xFX15)
		case 0x0018:
			//set ST to V[X] (0xFX18)
		case 0x001E:
			//set I = I + V[X] (0xFX1E)
		case 0x0029:
			//set I to sprite address for glyph of V[X] (4x5 px font) (0xFX29)
		case 0x0033:
			//get BCD representation of V[X]
			//Memory[I+0] = Decimal MSB Digit (3)
			//Memory[I+1] = Decimal Middle Digit (2)
			//Memory[I+3] = Decimal LSB Digit (3)
			//(0xFX33)
		case 0x0055:
			//Store V[0] to V[X] (inclusive) at memory location I, increasing I per register
			//(0xFX55)
		case 0x0065:
			//Set V[0] to V[x] (inclusive) to values from location I, increasing I per register
			//(0xFX65)
		}
	}
}
