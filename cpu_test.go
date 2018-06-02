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
	op = uint16(0x7108)
	cpu.ExecuteOp(op)
	cpu.ExecuteOp(op)
	if cpu.V[1] != 16 {
		t.Fail()
	}
}
