package chip8

import "os"

func LoadProgram(file string) []byte {
	mem := []byte{}

	f, _ := os.Open(file)
	f.Read(mem)
	return mem
}
