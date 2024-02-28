// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.
package tntengine

// Define the Crypter interface, constants, and variables used in tntengine

import (
	"bytes"
	"fmt"
	"math/bits"
)

// CipherBlock is the data processed by the crypters (rotors and permutators).
// It consistes of the length in bytes to process and the (32 bytes of) data to
// process.
type CipherBlock []byte

// Define constants needed for tntengine
const (
	BitsPerByte      int = 8
	CipherBlockSize  int = 256 // bits
	CipherBlockBytes int = CipherBlockSize / BitsPerByte
)

var (
	// BigZero - the big int value for zero.
	BigZero *Counter = &Counter{0, 0}
	// BigOne - the big int value for one.
	BigOne *Counter = &Counter{0, 1}
)

// String formats a string representing the CipherBlock.
func (cblk CipherBlock) String() string {
	var output bytes.Buffer
	blk := make([]byte, CipherBlockBytes)
	_ = copy(blk, cblk)
	output.WriteString("CipherBlock:")
	output.WriteString(fmt.Sprintf("\t     Length: %d\n", len(cblk)))
	output.WriteString(fmt.Sprintf("            \t   Capacity: %d\n", cap(cblk)))
	output.WriteString(fmt.Sprintf("            \t       Data:\t% X\n", blk[0:16]))
	output.WriteString(fmt.Sprintf("            \t\t\t% X", blk[16:]))
	return output.String()
}

// Crypter interface
type Crypter interface {
	Update(*Rand)                   // function to update the rotor/permutator
	SetIndex(*Counter)              // setter for the index value
	Index() *Counter                // getter for the index value
	ApplyF(CipherBlock) CipherBlock // encryption function
	ApplyG(CipherBlock) CipherBlock // decryption function
}

// Counter is the type used to hold the counter of blocks processed the
// the encryption engine.  The Counter is large enough to hold the maximum
// number of blocks that can be encrypted before the pattern repeats
type Counter [2]uint64

// Add - add n to the Counter.  Add() will panic if the Counter overflows
func (cntr *Counter) Add(n uint64) *Counter {
	var carry uint64 = 0
	i := len(cntr) - 1
	cntr[i], carry = bits.Add64(cntr[i], n, 0)
	for i--; i >= 0; i-- {
		cntr[i], carry = bits.Add64(cntr[i], 0, carry)
	}
	if carry != 0 {
		panic("overflow in Counter.Add()")
	}
	return cntr
}

func (dividend *Counter) Div(divisor uint64) *Counter {
	var r uint64 = 0
	var i = 0
	// Skip zeroos in the dividend
	for ; i < len(dividend); i++ {
		if dividend[i] != 0 {
			break
		}
	}
	// divide the remaining parts of the divivend
	for ; i < len(dividend); i++ {
		dividend[i], r = bits.Div64(r, dividend[i], divisor)
	}
	return dividend
}

func (dividend *Counter) DivMod(divisor uint64) (*Counter, uint64) {
	var r uint64 = 0
	var i = 0
	// Skip zeroos in the dividend
	for ; i < len(dividend); i++ {
		if dividend[i] != 0 {
			break
		}
	}
	// divide the remaining parts of the divivend
	for ; i < len(dividend); i++ {
		dividend[i], r = bits.Div64(r, dividend[i], divisor)
	}
	return dividend, r
}

// Index - increment the counter.
func (cntr *Counter) Increment() {
	var carry uint64 = 0
	i := len(cntr) - 1
	cntr[i], carry = bits.Add64(cntr[i], 1, carry)
	for i--; i >= 0; i-- {
		cntr[i], carry = bits.Add64(cntr[i], 0, carry)
	}
	if carry != 0 {
		panic("overflow in Counter.Increment()")
	}
}

// Index - retrieves the current index value
func (cntr *Counter) Index() *Counter {
	return cntr
}

// Update is a no-op for Counters
func (cntr *Counter) Update(random *Rand) {
	// Do nothing.
}

// Check to see if the Counter is zero
func (index *Counter) IsZero() bool {
	for _, val := range index {
		if val != 0 {
			return false
		}
	}
	return true
}

// Mod - calculate mod(counter, n)
func (dividend Counter) Mod(divisor uint64) uint64 {
	var r uint64 = 0
	var i = 0
	for ; i < len(dividend); i++ {
		if dividend[i] != 0 {
			break
		}
	}
	for ; i < len(dividend); i++ {
		_, r = bits.Div64(r, dividend[i], divisor)
	}
	return r
}

// Mul - Multiplies Counter by n.  It will panic if the
// results overflows cntr.index
func (multiplicand *Counter) Mul(multiplier uint64) *Counter {
	var carry uint64 = 0
	for i := len(multiplicand) - 1; i >= 0; i-- {
		var cary uint64
		if multiplicand[i] == 0 {
			cary = 0
		} else {
			cary, multiplicand[i] = bits.Mul64(multiplicand[i], multiplier)
		}
		for j := i; j >= 0 && carry != 0; j-- {
			multiplicand[j], carry = bits.Add64(multiplicand[j], carry, 0)
		}
		carry = cary
	}
	if carry != 0 {
		panic("Counter overflow in Counter.Mul()")
	}
	return multiplicand
}

// Set initializes the Counter to the given value.
func (index *Counter) SetIndex(val *Counter) {
	*index = *val
}

// SetString - sets the Counter to the numeric value of the (base10) string of digits.
// SetString will panic if the string contains non-digit characters.
func (index *Counter) SetString(val string) (*Counter, bool) {
	var good bool = true
	index.SetIndex(BigZero)
	var carry uint64
	for _, chr := range val {
		if chr < 48 || chr > 57 {
			good = false
			return index, good
		}
		index.Mul(10)
		j := len(index) - 1
		index[j], carry = bits.Add64(index[j], (uint64(chr) - 48), 0)
		for i := j - 1; i >= 0 && carry > 0; i-- {
			index[i], carry = bits.Add64(index[i], 0, carry)
		}
	}
	if carry != 0 {
		panic("Counter overflow in Counter.SetSTring()")
	}
	return index, good
}

// String - convert the Counter to a base10 string representing it's numeric value.
func (index Counter) String() string {
	var buf = make([]byte, 154)
	bIdx := len(buf) - 1
	if !index.IsZero() {
		var dividend Counter
		dividend.SetIndex(&index)
		r := uint64(0)
		for !dividend.IsZero() {
			_, r = dividend.DivMod(10)
			buf[bIdx] = byte(r + 48)
			bIdx--
		}
		return fmt.Sprint(string(buf[bIdx+1:]))
	}

	return "0"
}

// Sub - subtract n from Counter.  Sub() will panic if the Counter underflows.
func (cntr *Counter) Sub(n uint64) *Counter {
	var borrow uint64 = 0
	i := len(cntr) - 1
	cntr[i], borrow = bits.Sub64(cntr[i], n, borrow)
	for i--; i >= 0; i-- {
		cntr[i], borrow = bits.Sub64(cntr[i], 0, borrow)
	}
	if borrow != 0 {
		panic("underflow in Counter.Sub()")
	}
	return cntr
}

// ApplyF - increments the counter for each block that is encrypted.
func (cntr *Counter) ApplyF(blk CipherBlock) CipherBlock {
	cntr.Increment()
	return blk
}

// ApplyG - this function does nothing for a Counter during decryption.
func (cntr *Counter) ApplyG(blk CipherBlock) CipherBlock {
	return blk
}

// SubBlock -  subtracts (not XOR) the key from the data to be decrypted
func SubBlock(blk, key CipherBlock) CipherBlock {
	var p int
	for idx, val := range blk {
		p = p + int(val) - int(key[idx])
		blk[idx] = byte(p & 0xFF)
		p >>= BitsPerByte
	}
	return blk
}

// AddBlock - adds (not XOR) the data to be encrypted with the key.
func AddBlock(blk, key CipherBlock) CipherBlock {
	var p int
	for idx, val := range blk {
		p += int(val) + int(key[idx])
		blk[idx] = byte(p & 0xFF)
		p >>= BitsPerByte
	}
	return blk
}

// EncryptMachine - set up a rotor, permutator, or counter to encrypt a block
// read from the left (input channel) and send it out on the right (output channel)
func EncryptMachine(ecm Crypter, left chan CipherBlock) chan CipherBlock {
	if ecm == nil {
		panic("ecm is nil")
	}
	right := make(chan CipherBlock)
	go func(ecm Crypter, left chan CipherBlock, right chan CipherBlock) {
		defer close(right)
		for {
			inp := <-left
			if len(inp) <= 0 {
				right <- inp
				ecm = nil
				break
			}
			inp = ecm.ApplyF(inp)
			right <- inp
		}
	}(ecm, left, right)
	return right
}

// DecryptMachine - set up a rotor, permutator, or counter to decrypt a block
// read from the left (input channel) and send it out on the right (output channel)
func DecryptMachine(ecm Crypter, left chan CipherBlock) chan CipherBlock {
	if ecm == nil {
		panic("ecm is nil")
	}
	right := make(chan CipherBlock)
	go func(ecm Crypter, left chan CipherBlock, right chan CipherBlock) {
		defer close(right)
		for {
			inp := <-left
			if len(inp) <= 0 {
				right <- inp
				ecm = nil
				break
			}
			inp = ecm.ApplyG(inp)
			right <- inp
		}
	}(ecm, left, right)
	return right
}

// CreateEncryptMachine - Chain the encryption machines together, using channels to pass
// the data to be encrypted to the individual encryption machines.
// The data is entered in the 'left' channel and the encrypted data is read from the
// 'right' channel.
func createEncryptMachine(ecms ...Crypter) (left chan CipherBlock, right chan CipherBlock) {
	if ecms != nil {
		idx := 0
		left = make(chan CipherBlock)
		right = EncryptMachine(ecms[idx], left)
		for idx++; idx < len(ecms); idx++ {
			right = EncryptMachine(ecms[idx], right)
		}
	} else {
		panic("you must give at least one encryption device!")
	}
	return
}

// CreateDecryptMachine - Chain the decryption machines together (in reverse order), using
// channels to pass the data to be decrypted to the individual decryption machines.
// The encrypted data is entered in the 'left' channel and the plaintext data is read from
// the 'right' channel.
func createDecryptMachine(ecms ...Crypter) (left chan CipherBlock, right chan CipherBlock) {
	if ecms != nil {
		idx := len(ecms) - 1
		left = make(chan CipherBlock)
		right = DecryptMachine(ecms[idx], left)
		for idx--; idx >= 0; idx-- {
			right = DecryptMachine(ecms[idx], right)
		}
	} else {
		panic("you must give at least one decryption device!")
	}
	return
}
