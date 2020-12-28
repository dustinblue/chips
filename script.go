package main

import "bytes"
import "fmt"
import "crypto/sha256"
import "golang.org/x/crypto/ripemd160"
import "encoding/hex"

// Implement a bitcoin like script processing stack
// https://en.bitcoin.it/wiki/Script

const OP_PUSHDATA1 = 0x4c
const OP_EQUAL = 0x87
const OP_EQUAL_VERIFY = 0x88
const OP_HASH160 = 0xa9
const OP_DUP = 0x76
const OP_VERIFY = 0x69

// Represents our stack items of variable length bytes
type stack struct {
	data [][]byte
}

// Push item onto the stack
func (s *stack) push(bytes []byte) {
	s.data = append([][]byte{bytes}, s.data...)
}
  
// Pop item off the stack
func (s *stack) pop() []byte {
	pop := s.data[0]
	s.data = s.data[1:]
	return pop
}

// Represents our coin script
type script struct {
	data []byte
}

// Read (pop) 1 byte from our script.
func (s *script) read1() int {
	pop := s.data[0]
	s.data = s.data[1:]
	return int(pop)
}

// Read (pop) count bytes from our script and return them
func (s *script) read(count int) []byte {
	pop := s.data[:count]
	s.data = s.data[count:]
	return pop
}

func (s *script) writeHex(data string) {
	bytes, _ := hex.DecodeString(data)
	s.data = append(s.data, bytes...)
}

func (s *script) write(bytes []byte) {
	s.data = append(s.data, bytes...)
}

func (s *script) writeByte(data int) {
	s.data = append(s.data, byte(data))
}

func op_pushdata1(Script *script,  Stack *stack) {
	// First byte is length of data
	len := Script.read1()
	data := Script.read(len)
	Stack.push(data)
}

// compare 2 items from stack
func op_equal(Script *script,  Stack *stack) {
	comp1 := Stack.pop()
	comp2 := Stack.pop()
	result := bytes.Compare(comp1, comp2)
	if (result == 0) {
		Stack.push([]byte{1})
	} else {
		Stack.push([]byte{0})
	}
}

func op_equal_verify(Script *script, Stack *stack) {
	op_equal(Script, Stack)
	op_verify(Script, Stack)
}

func op_verify(Script *script, Stack *stack) {
	top := Stack.pop()
	result := bytes.Compare(top, []byte{1})
	// Script is invalid empty script and leave 0 on stack
	if (result != 0) {
		Script.data = []byte{}
		Stack.data = [][]byte{[]byte{0}}
	}
}

// ripemd160(sha256(data)) put result on stack
func op_hash160(Script *script,  Stack *stack) {
	data := Stack.pop()
	hash256 := sha256.Sum256(data)
	h := ripemd160.New()
	h.Write(hash256[:])
	hash160 := h.Sum(nil)
	Stack.push(hash160)
}

// ripemd160(sha256(data)) put result on stack
func op_dup(Script *script,  Stack *stack) {
	data := Stack.pop()
	Stack.push(data)
	Stack.push(data)
}

func executeScript(Script *script) *stack {
	
	stack := stack{}

	for (len(Script.data) > 0) {
		fmt.Printf("%x\nn", stack)

		// pop 1 btye for op code
		opcode := Script.read1()

		switch opcode {
			case OP_PUSHDATA1:
				op_pushdata1(Script, &stack)
			case OP_EQUAL:
				op_equal(Script, &stack)
			case OP_EQUAL_VERIFY:
				op_equal_verify(Script, &stack) 
			case OP_VERIFY:
				op_verify(Script, &stack) 
			case OP_HASH160:
				op_hash160(Script, &stack)
			case OP_DUP:
				op_dup(Script, &stack)
		}
		fmt.Printf("%x\n\n", stack)
	}

	// Return whatever is left on the stack
	return &stack
}