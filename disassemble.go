package chip8

import "fmt"

func DisassembleProgram(program []byte) string {
	pc := 0
	code := ""
	for {
		opCode := uint16(program[pc])<<8 | uint16(program[pc+1])

		code += fmt.Sprintf("[%#000x] %s\n", pc+0x200, disassemble(opCode))
		if pc > len(program)-3 {
			return code
		}
		pc += 2
	}
}

func disassemble(opCode uint16) string {

	switch opCode & 0xF000 {
	case 0x0000:
		switch opCode {
		case 0x00E0:
			return "CLS"
		case 0x00EE:
			return "RTN"
		default:
			return "RCA"
		}
	case 0x1000:
		//jump to address NNN (0x1NNN)
		addr := opCode & 0x0FFF
		return fmt.Sprintf("JMP %#x", addr)
	case 0x2000:
		//call subroutine NNN (0x2NNN)
		addr := opCode & 0x0FFF
		return fmt.Sprintf("SBR %#x", addr)
	case 0x3000:
		// skip next instruction if V[X]==NN (0x3XNN)
		addr := opCode & 0x00FF
		x := (opCode & 0x0F00) >> 8
		return fmt.Sprintf("JEQ v%d,%#x", x, addr)
	case 0x4000:
		// skip next instruction if V[X]!=NN (0x4XNN)
		addr := opCode & 0x00FF
		x := (opCode & 0x0F00) >> 8
		return fmt.Sprintf("JNE v%d,%#x", x, addr)
	case 0x5000:
		// skip next instruction if V[X]==V[Y] (0x3XY0)
		x := (opCode & 0x0F00) >> 8
		y := (opCode & 0x00F0) >> 4
		return fmt.Sprintf("JEQ v%d,v%d", x, y)
	case 0x6000:
		//set V[X] to NN (0x6XNN)addr := opCode & 0x0FFF
		addr := opCode & 0x00FF
		x := (opCode & 0x0F00) >> 8
		return fmt.Sprintf("SET v%d,%#x", x, addr)
	case 0x7000:
		//set V[X] to V[X]+NN (0x7XNN)
		addr := opCode & 0x00FF
		x := (opCode & 0x0F00) >> 8
		return fmt.Sprintf("ADD v%d,%#x", x, addr)
	case 0x8000:
		switch opCode & 0x000F {
		case 0x0000:
			//set V[X] to V[Y] (0x8XY0)
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("SET v%d,v%d", x, y)
		case 0x0001:
			//set V[X] to V[X] OR V[Y] (0x8XY1)
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("OR v%d,v%d", x, y)
		case 0x0002:
			//set V[X] to V[X] AND V[Y] (0x8XY2)
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("AND v%d,v%d", x, y)
		case 0x0003:
			//set V[X] to V[X] XOR V[Y] (0x8XY3)
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("XOR v%d,v%d", x, y)
		case 0x0004:
			//set V[X] to V[X] + V[Y] (0x8XY4), set V[F] to 1 if carry, otherwise 0
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("ADD v%d,v%d", x, y)
		case 0x0005:
			//set V[X] to V[X] - V[Y] (0x8XY5), set V[F] to 1 if no borrow, otherwise 0
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("SUB v%d,v%d", x, y)
		case 0x0006:
			//set V[X] to V[Y] >> 1 (0x8XY6), set V[F] to V[Y] LSB before shift
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("BSR v%d,v%d", x, y)
		case 0x0007:
			//set V[X] to V[Y] - V[X] (0x8XY7), set V[F] to 1 if no borrow, otherwise 0
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("SUB v%d,v%d", y, x)
		case 0x000E:
			//set V[X] (and V[Y]) to V[Y] << 1 (0x8XYE), set V[F] to V[Y] MSB before shift
			x := (opCode & 0x0F00) >> 8
			y := (opCode & 0x00F0) >> 4
			return fmt.Sprintf("BSL v%d,v%d", x, y)
		}
	case 0x9000:
		//skip next instruction if V[X]!=V[Y] (0x9XY0)
		x := (opCode & 0x0F00) >> 8
		y := (opCode & 0x00F0) >> 4
		return fmt.Sprintf("JNE v%d,v%d", x, y)
	case 0xA000:
		//set I to NNN (0xANNN)
		imm := opCode & 0x0FFF
		return fmt.Sprintf("ADR %#x", imm)
	case 0xB000:
		//jump to addr V[0]+NNN: set PC to V[0] +NNN (0xBNNN)
		imm := opCode & 0x0FFF
		return fmt.Sprintf("JMA %#x", imm)
	case 0xC000:
		//set V[x] to R+NN where R = random number between 0 and 255(0xCXNN)
		x := (opCode & 0x0F00) >> 8
		imm := opCode & 0x00FF
		return fmt.Sprintf("RND v%d,%#x", x, imm)
	case 0xD000:
		//draw sprite at position (V[X],V[Y]) with width 8, heigh N.
		//sprite bits located at Memory[I] in rows of 8 (0xDXYN)
		//V[F] is set to 1 if pixels are flipped from 1 to 0, otherwise 0
		x := (opCode & 0x0F00) >> 8
		y := (opCode & 0x00F0) >> 4
		return fmt.Sprintf("DRW v%d,v%d", x, y)

	case 0xE000:
		switch opCode & 0x00FF {
		case 0x009E:
			//skip next instruction if key pressed == v[X] (0xEX9E)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("JKP v%d", x)
		case 0x00A1:
			//skip next instruction if key pressed != v[X] (0xEXA1)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("JKN v%d", x)
		}
	case 0xF000:
		switch opCode & 0x00FF {
		case 0x0007:
			//set V[X] to DT (0xFX07)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("SET v%d,DT", x)
		case 0x000A:
			//wait for and set v[X] to key pressed (0xFX0A)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("WKP v%d", x)
		case 0x0015:
			//set DT to V[X] (0xFX15)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("SET DT,v%d", x)
		case 0x0018:
			//set ST to V[X] (0xFX18)

			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("SET ST,v%d", x)
		case 0x001E:
			//set I = I + V[X] (0xFX1E)
		case 0x0029:
			//set I to sprite address for glyph of V[X] (4x5 px font) (0xFX29)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("Fnt v%d", x)
		case 0x0033:
			//get BCD representation of V[X]
			//Memory[I+0] = Decimal MSB Digit (3)
			//Memory[I+1] = Decimal Middle Digit (2)
			//Memory[I+3] = Decimal LSB Digit (3)
			//(0xFX33)
			x := (opCode & 0x0F00) >> 8
			return fmt.Sprintf("BCD v%d", x)
		case 0x0055:
			//Store V[0] to V[X] (inclusive) at memory location I, increasing I per register
			//(0xFX55)
			return "DMP"
		case 0x0065:
			//Set V[0] to V[x] (inclusive) to values from location I, increasing I per register
			//(0xFX65)
			return "LOD"
		}
	}
	return fmt.Sprintf("!!! %#x", opCode)
}
