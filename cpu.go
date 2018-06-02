package chip8

import (
	"crypto/rand"
	"fmt"
	"time"
)

const WordLength = 2

type CPU struct {

	//registers
	V [16]byte

	//address register
	I uint16

	//program counter
	PC uint16

	//stack
	Stack [48]uint16

	//stack pointer
	SP byte

	//delay timer
	DT byte

	//sound timer
	ST byte

	//Memory
	Memory [4096]byte

	Clock <-chan time.Time

	Finished bool
}

func NewCPU(timer <-chan time.Time) *CPU {
	return &CPU{
		Clock: timer,
		PC:    0x200,
	}
}

func (c *CPU) Run() {
	fmt.Println("Running Chip8")
	fmt.Println("Starting...")
	for {
		select {
		case <-c.Clock:
			c.Execute()
			if c.Finished {
				fmt.Println("Finished")
				return
			}
		}
	}
}

func (c *CPU) Execute() {

	if c.PC > 4083 {
		c.Finished = true
		return
	}

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
			c.PC += WordLength
		case 0x00EE:
			//return from subroutrine
			c.PC = c.PopFromStack()
			c.PC += WordLength

		default:
			// ignore RCA1802 functions
			//NOP
			c.PC += WordLength
		}
	case 0x1000:
		//jump to address NNN (0x1NNN)
		addr := opCode & 0x0FFF
		c.PC = addr
	case 0x2000:
		//call subroutine NNN (0x2NNN)
		c.PushToStack(c.PC)
		c.PC = opCode & 0x0FFF
	case 0x3000:
		// skip next instruction if V[X]==NN (0x3XNN)
		x := (opCode & 0x0F00) >> 8
		vx := c.V[x]
		imm := byte(opCode & 0x00FF)
		if imm == vx {
			c.PC += 2 * WordLength
		} else {
			c.PC += WordLength
		}
	case 0x4000:
		// skip next instruction if V[X]!=NN (0x4XNN)
		x := (opCode & 0x0F00) >> 8
		vx := c.V[x]
		imm := byte(opCode & 0x00FF)
		if imm != vx {
			c.PC += 2 * WordLength
		} else {
			c.PC += WordLength
		}
	case 0x5000:
		// skip next instruction if V[X]==V[Y] (0x5XY0)
		x := (opCode & 0x0F00) >> 8
		vx := c.V[x]
		y := (opCode & 0x00F0) >> 4
		vy := c.V[y]
		if vy == vx {
			c.PC += 2 * WordLength
		} else {
			c.PC += WordLength
		}
	case 0x6000:
		//set V[X] to NN (0x6XNN)
		x := (opCode & 0x0F00) >> 8
		imm := byte(opCode & 0x00FF)
		c.V[x] = imm
		c.PC += WordLength
	case 0x7000:
		//set V[X] to V[X}+NN (0x7XNN)
		//don't change carry bit
		x := (opCode & 0x0F00) >> 8
		imm := byte(opCode & 0x00FF)
		c.V[x] = imm + c.V[x]
		c.PC += WordLength
	case 0x8000:
		switch opCode & 0x000F {
		case 0x0000:
			//set V[X] to V[Y] (0x8XY0)
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vy
			c.PC += WordLength
		case 0x0001:
			//set V[X] to V[X] OR V[Y] (0x8XY1)
			x := (opCode & 0x0F00) >> 8
			vx := c.V[x]
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vx | vy
			c.PC += WordLength
		case 0x0002:
			//set V[X] to V[X] AND V[Y] (0x8XY2)
			x := (opCode & 0x0F00) >> 8
			vx := c.V[x]
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vx & vy
			c.PC += WordLength
		case 0x0003:
			//set V[X] to V[X] XOR V[Y] (0x8XY3)
			x := (opCode & 0x0F00) >> 8
			vx := c.V[x]
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vx ^ vy
			c.PC += WordLength
		case 0x0004:
			//set V[X] to V[X] + V[Y] (0x8XY4), set V[F] to 1 if carry, otherwise 0
			x := (opCode & 0x0F00) >> 8
			vx := c.V[x]
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vx + vy
			if vx+vy > 0xFF {
				c.V[0xF] |= 0x01
			} else {
				c.V[0xF] &= 0xFE
			}
			c.PC += WordLength
		case 0x0005:
			//set V[X] to V[X] - V[Y] (0x8XY5), set V[F] to 1 if no borrow, otherwise 0
			x := (opCode & 0x0F00) >> 8
			vx := c.V[x]
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vx - vy
			if vx < vy {
				c.V[0xF] |= 0x01
			} else {
				c.V[0xF] &= 0xFE
			}
			c.PC += WordLength
		case 0x0006:
			//set V[X] to V[Y] >> 1 (0x8XY6), set V[F] to V[Y] LSB before shift
			x := (opCode & 0x0F00) >> 8

			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[0xF] = vy & 0x01
			c.V[x] = vy >> 1

			c.PC += WordLength
		case 0x0007:
			//set V[X] to V[Y] - V[X] (0x8XY7), set V[F] to 1 if no borrow, otherwise 0
			x := (opCode & 0x0F00) >> 8
			vx := c.V[x]
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[x] = vy - vx
			if vy < vx {
				c.V[0xF] |= 0x01
			} else {
				c.V[0xF] &= 0xFE
			}
			c.PC += WordLength
		case 0x000E:
			//set V[X] (and V[Y]) to V[Y] << 1 (0x8XYE), set V[F] to V[Y] MSB before shift
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			vy := c.V[y]
			c.V[0xF] = c.V[0xF] & (((vy & 0x80) >> 7) | 0xFE)
			c.V[x] = vy << 1
			c.V[y] = vy << 1

			c.PC += WordLength
		}
	case 0x9000:
		//skip next instruction if V[X]!=V[Y] (0x9XY0)
		x := (opCode & 0x0F00) >> 8
		vx := c.V[x]
		y := (opCode & 0x00F0) >> 4
		vy := c.V[y]
		if vy != vx {
			c.PC += 2 * WordLength
		} else {
			c.PC += WordLength
		}
	case 0xA000:
		//set I to NNN (0xANNN)
		addr := opCode & 0x0FFF
		c.I = addr
		c.PC += WordLength
	case 0xB000:
		//jump to addr V[0]+NNN: set PC to V[0] +NNN (0xBNNN)
		addr := opCode & 0x0FFF
		c.PushToStack(c.PC)
		c.PC = addr + uint16(c.V[0])
	case 0xC000:
		//set V[x] to R & NN where R = random number between 0 and 255(0xCXNN)
		rnd := []byte{0xFF}
		rand.Read(rnd)
		x := (opCode & 0x0F00) >> 8
		imm := byte(opCode & 0x00FF)
		c.V[x] = imm & rnd[0]
		c.PC += WordLength
	case 0xD000:
		//draw sprite at position (V[X],V[Y]) with width 8, heigh N.
		//sprite bits located at Memory[I] in rows of 8 (0xDXYN)
		//V[F] is set to 1 if pixels are flipped from 1 to 0, otherwise 0
		//TODO
		c.PC += WordLength
	case 0xE000:
		switch opCode & 0x00FF {
		case 0x009E:
			//skip next instruction if key pressed == v[X] (0xEX9E)
			//TODO
			c.PC += WordLength
		case 0x00A1:
			//skip next instruction if key pressed != v[X] (0xEXA1)
			//TODO
			c.PC += WordLength
		}
	case 0xF000:
		switch opCode & 0x00FF {
		case 0x0007:
			//set V[X] to DT (0xFX07)
			x := (opCode & 0x0F00) >> 8
			c.V[x] = c.DT
			c.PC += WordLength

		case 0x000A:
			//set v[X] to key pressed (0xFX0A)
			//TODO
			c.PC += WordLength
		case 0x0015:
			//set DT to V[X] (0xFX15)
			x := (opCode & 0x0F00) >> 8
			c.DT = c.V[x]
			c.PC += WordLength
		case 0x0018:
			//set ST to V[X] (0xFX18)
			x := (opCode & 0x0F00) >> 8
			c.ST = c.V[x]
			c.PC += WordLength
		case 0x001E:
			//set I = I + V[X] (0xFX1E)
			x := (opCode & 0x0F00) >> 8
			c.I += uint16(c.V[x])
			c.PC += WordLength
		case 0x0029:
			//set I to sprite address for glyph of V[X] (4x5 px font) (0xFX29)
			x := (opCode & 0x0F00) >> 8
			c.I = uint16(c.V[x] * 5) //each glyph is 5 bytes
			c.PC += WordLength
		case 0x0033:
			//get BCD representation of V[X]
			//Memory[I+0] = Decimal MSB Digit (3)
			//Memory[I+1] = Decimal Middle Digit (2)
			//Memory[I+3] = Decimal LSB Digit (3)
			//(0xFX33)
			x := (opCode & 0x0F00) >> 8
			c.Memory[c.I] = c.V[x] / 100
			c.Memory[c.I+1] = (c.V[x] / 10) % 10
			c.Memory[c.I+2] = (c.V[x] % 100) % 10
			c.PC += WordLength
		case 0x0055:
			//Store V[0] to V[X] (inclusive) at memory location I, increasing I per register
			//(0xFX55)
			for i := 0; i < 15; i++ {
				c.I += uint16(i)
				c.Memory[c.I] = c.V[i]
			}
			c.PC += WordLength
		case 0x0065:
			//Set V[0] to V[x] (inclusive) to values from location I, increasing I per register
			//(0xFX65)
			for i := 0; i < 15; i++ {
				c.I += uint16(i)
				c.V[i] = c.Memory[c.I]
			}
			c.PC += WordLength
		default:
			//nop
			c.PC += WordLength
		}
	}

	fmt.Printf("\n[%s] PC: %#x SP: %#x\n", disassemble(opCode), c.PC, c.SP)
	for i, data := range c.V {
		fmt.Printf(" V%d: %#x", i, data)
	}

}

func (c *CPU) PushToStack(addr uint16) {

	if c.SP < byte(len(c.Stack)) {
		c.SP++
		c.Stack[c.SP] = addr
	} else {
		panic("Stack overflow")
	}

}

func (c *CPU) PopFromStack() uint16 {
	addr := c.Stack[c.SP]
	if c.SP > 0 {
		c.SP--
	} else {
		panic("stack overflow")
	}
	return addr
}

func (c *CPU) LoadData(addr uint16, data []byte) {

	for i, b := range data {
		if (int(addr) + i) > len(c.Memory) {
			return
		}
		c.Memory[int(addr)+i] = b
	}
}
