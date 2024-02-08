// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.
package tntengine

// Define the Crypter interface, constants, and variables used in tntengine

import (
	"bytes"
	"fmt"
	"math/big"
)

// Define constants needed for tntengine
const (
	BitsPerByte      int = 8
	CipherBlockSize  int = 256 // bits
	CipherBlockBytes int = CipherBlockSize / BitsPerByte
)

var (
	// BigZero - the big int value for zero.
	BigZero = big.NewInt(0)
	// BigOne - the big int value for one.
	BigOne = big.NewInt(1)
)

// CipherBlock is the data processed by the crypters (rotors and permutators).
// It consistes of the length in bytes to process and the (32 bytes of) data to
// process.
type CipherBlock []byte

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
	SetIndex(*big.Int)              // setter for the index value
	Index() *big.Int                // getter for the index value
	ApplyF(CipherBlock) CipherBlock // encryption function
	ApplyG(CipherBlock) CipherBlock // decryption function
}

// Counter is a crypter that does not encrypt/decrypt any data but counts the
// number of blocks that were encrypted.
type Counter struct {
	index *big.Int
}

func (cntr *Counter) Update(random *Rand) {
	// Do nothing.
}

// SetIndex - sets the initial index value
func (cntr *Counter) SetIndex(index *big.Int) {
	cntr.index = new(big.Int).Set(index)
}

// Index - retrieves the current index value
func (cntr *Counter) Index() *big.Int {
	return cntr.index
}

// ApplyF - increments the counter for each block that is encrypted.
func (cntr *Counter) ApplyF(blk CipherBlock) CipherBlock {
	cntr.index.Add(cntr.index, BigOne)
	return blk
}

// ApplyG - this function does nothing for a Counter during decryption.
func (cntr *Counter) ApplyG(blk CipherBlock) CipherBlock {
	return blk
}

func (cntr *Counter) String() string {
	return fmt.Sprint(cntr.index)
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
