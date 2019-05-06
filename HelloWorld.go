package tools

import (
        "fmt"
        "github.com/goburrow/modbus"
)

//TT_lrc is a structure that implements a simple Longitudinal Redundancy Check.
//It allows bytes to be added one at a time(AddByte) or it can process an entire array(AddBytes).
//The Reset function can be used to reset the sum and start a new one.
type TT_lrc struct {
	xorSum byte
}

//Reset resets the current sum so that it can be used to calculate a new one.
func (lrc *TT_lrc) Reset() *TT_lrc {
	lrc.xorSum = 0
	return lrc
}

//AddByte adds a new byte to the LRC, it returns the TT_lrc so that multiple calls can be linked.
//	Example
//		var someLrc TT_lrc
//		someLrc.AddByte(0x31).AddByte(0x30).AddByte(0x31)   // Adds 2 bytes
//		fmt.printf("LRC:%d", someLrc.LRC())                 // Prints:LRC:30
func (lrc *TT_lrc) AddByte(aByte byte) *TT_lrc {
	lrc.xorSum = lrc.xorSum ^ aByte

	return lrc
}

//AddBytes adds a all the bytes from a slice to the LRC, it returns the LRC sum so that the calculation 
//can getting the LRC can be done in one step.
//	Example
//		var someLrc TT_lrc
//		fmt.printf("LRC:%d", someLrc.AddBytes({0x30,0x31}))		// Adds 2 bytes and Prints:LRC:1
func (lrc *TT_lrc) AddBytes(someBytes []byte) byte {
	var aByte byte
	for _, aByte = range someBytes {
		lrc.xorSum = lrc.xorSum ^ aByte
	}

	return lrc.xorSum
}

//LRC returns the LRC of all the bytes added
//	Example
//		var someLrc TT_lrc
//		someLrc.AddByte(0x30).AddByte(0x31)		// Adds 2 bytes
//		fmt.printf("LRC:%d", someLrc.LRC())		// Prints:LRC:1
func (lrc *TT_lrc) LRC() byte {	
	return lrc.xorSum
}

