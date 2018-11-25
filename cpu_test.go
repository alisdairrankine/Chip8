package chip8_test

import (
	"testing"

	"github.com/alisdairrankine/chip8"
)

func TestOpCode0NNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode00E0(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode00EE(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode1NNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode2NNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode3XNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode4XNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode5XY0(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode6XNN(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x6058)
	cpu.ExecuteOp(op)
	if cpu.V[0] != 88 {
		t.Fail()
	}
}
func TestOpCode7XNN(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x7008)
	cpu.ExecuteOp(op)
	cpu.ExecuteOp(op)
	if cpu.V[0] != 16 {
		t.Fail()
	}
}
func TestOpCode8XY0(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x6058)
	cpu.ExecuteOp(op)
	if cpu.V[0] != 88 {
		t.Fail()
	}
	op = uint16(0x8100)
	cpu.ExecuteOp(op)
	if cpu.V[1] != 88 {
		t.Fail()
	}
	t.Skip()
}
func TestOpCode8XY1(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x8011)
	cpu.V[0] = 0x55
	cpu.V[1] = 0x02
	cpu.ExecuteOp(op)
	if cpu.V[0] != 87 {
		t.Fail()
	}

}
func TestOpCode8XY2(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x8012)
	cpu.V[0] = 0x55
	cpu.V[1] = 0xFE
	cpu.ExecuteOp(op)
	if cpu.V[0] != 84 {
		t.Fail()
	}

}
func TestOpCode8XY3(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x8013)
	cpu.V[0] = 0x55
	cpu.V[1] = 0xFF
	cpu.ExecuteOp(op)
	if cpu.V[0] != 0xAA {
		t.Fail()
	}
}
func TestOpCode8XY4(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x8014)
	cpu.V[0] = 0xFE
	cpu.V[1] = 0x01
	cpu.ExecuteOp(op)
	if cpu.V[0] != 0xFF || (cpu.V[0xF]&0x01) != 0 {
		t.Errorf("V[0] expected: %#x, actual: %#x", 0x00, cpu.V[0])
		t.Errorf("v[F] expected: %#x, actual: %#x", 0x00, cpu.V[0xF])
		t.Fail()
	}

	op = uint16(0x8014)
	cpu.V[0] = 0xFF
	cpu.V[1] = 0x02
	cpu.ExecuteOp(op)
	if cpu.V[0] != 0x01 || (cpu.V[0xF]&0x01) != 1 {
		t.Fail()
	}
}
func TestOpCode8XY5(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x8015)
	cpu.V[0] = 0xFF
	cpu.V[1] = 0x02
	cpu.ExecuteOp(op)
	if cpu.V[0] != 0xFD || (cpu.V[0xF]&0x01) != 0x01 {
		t.Fail()
	}

	// op = uint16(0x8015)
	// cpu.V[0] = 0x01
	// cpu.V[1] = 0x02
	// cpu.ExecuteOp(op)
	// if cpu.V[0] != 0x00 || (cpu.V[0xF]&0x01) != 0x00 {
	// 	t.Log("0x8015(2) failed")
	// 	t.Fail()
	// }
}
func TestOpCode8XY6(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode8XY7(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode8XYE(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCode9XY0(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeANNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeDXYN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeBNNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeCXNN(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeEX9E(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeEXA1(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX07(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX0A(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX15(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX18(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX1E(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX29(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX33(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX55(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
func TestOpCodeFX65(t *testing.T) {
	//cpu := chip8.NewCPU(nil)
	t.Skip()
}
