package main

import "testing"
import "fmt"
import "bytes"
// import "encoding/hex"

// A very basic script 4c01054c010587 compares 2 numbers

// PUSHDATA(4c) size(01 byte) data(05)   | 5 |
// PUSHDATA(4c) size(01 byte) data(05)   | 5 |
// OPEQUAL(87) pop 2 items from stack and compare. Put 1 on stack if equal else put 0 on stack
func TestScriptEqual(t *testing.T) {
	
	// Should have 0x01 on stack
	sigScript := script{data: []byte{OP_PUSHDATA1, 1, 5, OP_PUSHDATA1, 1, 5, OP_EQUAL}}
	fmt.Printf("%x\n", sigScript)
	result := executeScript(&sigScript)
	if bytes.Compare(result.pop(), []byte{1}) != 0 {
		t.Errorf("Basic comparison script failed")
	}

	// Same script as above in hex format
	sigScript = script{}
	sigScript.writeHex("4c01054c010587")
	result = executeScript(&sigScript)
	if bytes.Compare(result.pop(), []byte{1}) != 0 {
		t.Errorf("Basic comparison script failed")
	}

	// Should have 0x00 on stack
	sigScript = script{data: []byte{OP_PUSHDATA1, 1, 5, OP_PUSHDATA1, 1, 6, OP_EQUAL}}
	result = executeScript(&sigScript)
	if bytes.Compare(result.pop(), []byte{0}) != 0 {
		t.Errorf("Basic comparison script failed")
	}
}

// Test that public key 0250863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352
// when ripemd160(sha256(pubkey)) = f54a5851e9372b87810a8e60cdd2e7cfd80b6e31
func TestHashOp(t *testing.T) {
	sigScript := script{}
	sigScript.writeByte(OP_PUSHDATA1)
	sigScript.writeByte(33)
	sigScript.writeHex("0250863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352")
	sigScript.writeByte(OP_DUP)
	sigScript.writeByte(OP_HASH160)
	sigScript.writeByte(OP_PUSHDATA1)
	sigScript.writeByte(20)
	sigScript.writeHex("f54a5851e9372b87810a8e60cdd2e7cfd80b6e31")
	sigScript.writeByte(OP_EQUAL_VERIFY)
	fmt.Printf("%x\n", sigScript)

	result := executeScript(&sigScript);
	fmt.Printf("%x\n", result)
}