package chip8_test

import (
	"testing"

	"github.com/alisdairrankine/chip8"
)

func TestSET(t *testing.T) {

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
}

func TestMaths(t *testing.T) {
	cpu := chip8.NewCPU(nil)
	op := uint16(0x7008)
	cpu.ExecuteOp(op)
	cpu.ExecuteOp(op)
	if cpu.V[0] != 16 {
		t.Fail()
	}

	op = uint16(0x8011)
	cpu.V[0] = 0x55
	cpu.V[1] = 0x02
	cpu.ExecuteOp(op)
	if cpu.V[0] != 87 {
		t.Fail()
	}

	op = uint16(0x8012)
	cpu.V[0] = 0x55
	cpu.V[1] = 0xFE
	cpu.ExecuteOp(op)
	if cpu.V[0] != 84 {
		t.Fail()
	}

	op = uint16(0x8013)
	cpu.V[0] = 0x55
	cpu.V[1] = 0xFF
	cpu.ExecuteOp(op)
	if cpu.V[0] != 0xAA {
		t.Fail()
	}

	op = uint16(0x8014)
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
		t.Log("0x8014 failed")
		t.Fail()
	}

	op = uint16(0x8015)
	cpu.V[0] = 0xFF
	cpu.V[1] = 0x02
	cpu.ExecuteOp(op)
	if cpu.V[0] != 0xFD || (cpu.V[0xF]&0x01) != 0x01 {
		t.Log("0x8015(3) failed")
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

func TestGraphics(t *testing.T) {
	t.Skip()
}
