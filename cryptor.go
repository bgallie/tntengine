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
	BitsPerByte             int = 8
	CipherBlockSize         int = 256 // bits
	CipherBlockBytes        int = CipherBlockSize / BitsPerByte
	NumberPermutationCycles int = 1
)

var (
	// Define the proforma rotors and permutator used to create the actual rotors and permutators to use.
	Rotor1 = new(Rotor).New(1789, 1065, 1499, []byte{
		63, 180, 255, 162, 59, 142, 61, 13, 187, 226, 49, 134, 163, 38, 44, 14,
		255, 73, 155, 237, 208, 42, 217, 227, 194, 245, 229, 169, 96, 163, 33, 145,
		222, 156, 57, 87, 220, 186, 118, 131, 89, 103, 27, 145, 153, 207, 16, 55,
		248, 183, 83, 65, 15, 253, 147, 136, 217, 189, 124, 150, 193, 113, 87, 127,
		101, 202, 87, 3, 80, 160, 132, 129, 1, 134, 154, 36, 194, 3, 186, 148,
		241, 226, 134, 255, 59, 78, 202, 236, 166, 151, 184, 209, 115, 21, 177, 17,
		106, 189, 209, 128, 13, 224, 94, 163, 47, 117, 151, 3, 9, 88, 20, 74,
		188, 243, 174, 130, 193, 247, 161, 74, 119, 95, 40, 111, 215, 174, 84, 170,
		234, 27, 241, 147, 210, 26, 139, 92, 231, 118, 227, 206, 0, 186, 161, 82,
		149, 59, 93, 134, 84, 108, 116, 191, 127, 153, 92, 59, 80, 53, 10, 112,
		127, 228, 183, 134, 214, 74, 150, 134, 145, 60, 22, 217, 213, 195, 251, 240,
		232, 1, 193, 235, 142, 191, 153, 123, 46, 86, 198, 123, 33, 34, 148, 104,
		18, 96, 34, 17, 139, 199, 225, 84, 245, 102, 137, 167, 240, 84, 152, 144,
		171, 21, 67, 253, 113, 97, 156, 145, 55, 87, 247, 45, 54, 48, 157, 247,
		135, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	Rotor2 = new(Rotor).New(1787, 1624, 249, []byte{
		129, 226, 91, 144, 187, 249, 232, 34, 223, 100, 147, 100, 255, 141, 49, 97,
		81, 113, 220, 221, 74, 179, 165, 175, 248, 59, 108, 143, 97, 124, 233, 244,
		138, 234, 170, 166, 225, 88, 239, 179, 208, 180, 183, 222, 224, 190, 83, 1,
		158, 148, 103, 103, 168, 240, 236, 179, 251, 162, 120, 0, 187, 83, 231, 248,
		123, 74, 245, 204, 35, 16, 31, 43, 74, 34, 166, 112, 218, 30, 34, 108,
		5, 146, 77, 149, 246, 27, 50, 114, 252, 208, 196, 154, 46, 249, 50, 116,
		125, 218, 160, 82, 35, 175, 109, 84, 61, 240, 132, 79, 62, 92, 28, 180,
		210, 142, 82, 100, 246, 193, 23, 136, 47, 76, 170, 60, 225, 110, 145, 102,
		81, 230, 50, 160, 36, 84, 183, 113, 146, 182, 99, 167, 154, 12, 148, 148,
		196, 68, 78, 48, 109, 243, 27, 197, 212, 246, 141, 252, 91, 38, 160, 250,
		13, 130, 165, 184, 49, 203, 237, 13, 20, 172, 72, 233, 86, 248, 131, 35,
		255, 52, 149, 111, 5, 10, 188, 109, 197, 143, 227, 164, 55, 167, 12, 236,
		222, 190, 91, 45, 172, 23, 177, 214, 191, 168, 196, 144, 90, 244, 7, 91,
		54, 2, 135, 15, 204, 152, 10, 221, 4, 124, 53, 189, 112, 69, 181, 8,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	Rotor3 = new(Rotor).New(1783, 1056, 1256, []byte{
		85, 83, 244, 107, 246, 228, 12, 122, 179, 102, 163, 91, 221, 248, 59, 190,
		210, 118, 254, 185, 204, 251, 52, 148, 234, 218, 90, 143, 6, 219, 84, 21,
		188, 183, 152, 204, 251, 185, 80, 116, 152, 116, 37, 182, 8, 238, 129, 4,
		123, 170, 169, 14, 43, 180, 155, 57, 178, 238, 231, 95, 143, 197, 65, 220,
		152, 150, 223, 171, 153, 203, 63, 88, 163, 93, 229, 73, 165, 91, 162, 196,
		57, 72, 30, 9, 108, 216, 191, 60, 32, 214, 229, 183, 252, 121, 143, 47,
		73, 27, 219, 201, 107, 135, 45, 251, 165, 146, 253, 238, 126, 25, 203, 166,
		20, 151, 159, 139, 75, 214, 5, 142, 63, 48, 227, 24, 219, 221, 119, 86,
		191, 93, 138, 68, 224, 6, 192, 168, 106, 48, 161, 161, 161, 244, 61, 133,
		125, 21, 174, 235, 141, 134, 194, 58, 12, 196, 123, 8, 63, 71, 6, 80,
		243, 61, 20, 8, 41, 55, 155, 139, 153, 179, 122, 35, 21, 7, 157, 88,
		147, 164, 25, 167, 202, 57, 234, 142, 39, 169, 197, 205, 119, 173, 46, 146,
		140, 180, 35, 233, 40, 176, 180, 18, 76, 177, 105, 179, 30, 117, 146, 6,
		123, 146, 74, 136, 109, 83, 79, 23, 0, 217, 108, 66, 4, 232, 135, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	Rotor4 = new(Rotor).New(1777, 210, 241, []byte{
		193, 157, 83, 55, 201, 219, 89, 192, 90, 41, 34, 176, 0, 155, 13, 31,
		107, 9, 140, 207, 10, 42, 253, 149, 155, 130, 148, 165, 125, 211, 212, 203,
		225, 240, 255, 90, 33, 55, 196, 146, 99, 225, 49, 56, 231, 234, 99, 139,
		30, 218, 205, 200, 59, 197, 176, 28, 205, 117, 180, 205, 183, 213, 220, 186,
		46, 168, 228, 53, 54, 212, 110, 108, 112, 115, 46, 206, 226, 142, 34, 81,
		93, 122, 92, 113, 143, 89, 97, 167, 251, 179, 135, 139, 236, 200, 148, 29,
		116, 235, 18, 250, 98, 90, 126, 142, 246, 18, 81, 32, 17, 118, 241, 246,
		63, 255, 152, 125, 66, 45, 105, 162, 252, 15, 50, 122, 194, 159, 174, 216,
		79, 233, 98, 241, 119, 146, 161, 163, 125, 147, 119, 157, 193, 118, 189, 154,
		153, 183, 92, 142, 35, 85, 68, 108, 129, 84, 222, 55, 250, 15, 107, 171,
		248, 81, 207, 120, 226, 124, 84, 48, 56, 14, 31, 51, 204, 233, 255, 251,
		34, 86, 139, 146, 87, 0, 140, 54, 133, 7, 232, 14, 7, 207, 220, 38,
		63, 102, 6, 99, 44, 200, 9, 98, 104, 132, 122, 20, 235, 246, 81, 171,
		92, 173, 117, 228, 93, 41, 58, 60, 117, 73, 202, 213, 245, 102, 130, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	Rotor5 = new(Rotor).New(1759, 955, 1559, []byte{
		126, 94, 133, 59, 14, 15, 200, 24, 20, 186, 46, 74, 114, 124, 228, 123,
		68, 6, 0, 216, 31, 129, 43, 10, 39, 51, 61, 142, 52, 156, 11, 205,
		237, 20, 253, 239, 28, 56, 68, 105, 98, 164, 203, 65, 137, 25, 241, 207,
		46, 87, 90, 19, 35, 248, 12, 133, 200, 182, 11, 225, 157, 151, 205, 14,
		255, 16, 44, 78, 80, 77, 187, 186, 52, 169, 206, 165, 193, 210, 23, 191,
		1, 48, 223, 149, 241, 154, 136, 103, 160, 84, 21, 143, 119, 174, 18, 219,
		56, 162, 243, 237, 154, 54, 27, 102, 169, 129, 205, 245, 152, 96, 14, 122,
		173, 8, 187, 244, 48, 133, 210, 151, 96, 230, 103, 214, 244, 103, 91, 94,
		139, 118, 155, 172, 53, 125, 164, 233, 161, 75, 114, 77, 201, 67, 0, 202,
		91, 128, 115, 36, 7, 27, 235, 35, 94, 200, 234, 52, 222, 163, 101, 38,
		67, 45, 195, 36, 64, 218, 194, 67, 110, 144, 104, 95, 239, 27, 85, 236,
		144, 132, 181, 104, 223, 210, 11, 153, 209, 134, 18, 87, 29, 76, 206, 11,
		46, 81, 195, 190, 107, 129, 248, 22, 123, 29, 127, 111, 253, 13, 43, 244,
		1, 178, 204, 34, 207, 248, 237, 125, 97, 84, 171, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	Rotor6 = new(Rotor).New(1753, 370, 362, []byte{
		86, 98, 39, 86, 179, 206, 169, 158, 65, 148, 145, 37, 171, 206, 137, 80,
		96, 3, 185, 25, 178, 190, 93, 134, 87, 121, 9, 220, 250, 188, 132, 246,
		95, 62, 247, 88, 238, 47, 163, 19, 236, 121, 181, 231, 70, 185, 131, 160,
		185, 177, 141, 124, 246, 180, 177, 96, 105, 236, 99, 181, 226, 206, 207, 156,
		228, 137, 236, 223, 48, 104, 128, 227, 177, 231, 217, 94, 29, 14, 82, 41,
		19, 17, 16, 158, 44, 99, 36, 142, 255, 116, 43, 23, 162, 22, 179, 70,
		19, 24, 189, 189, 111, 51, 141, 13, 125, 73, 195, 113, 27, 75, 6, 82,
		24, 188, 208, 16, 186, 232, 137, 53, 19, 107, 23, 25, 155, 255, 192, 188,
		92, 81, 79, 10, 68, 1, 228, 66, 232, 166, 163, 139, 18, 149, 129, 125,
		238, 139, 47, 174, 254, 137, 202, 214, 239, 131, 212, 165, 62, 13, 135, 21,
		193, 222, 10, 145, 122, 199, 120, 117, 219, 200, 250, 125, 66, 196, 95, 193,
		217, 125, 136, 105, 1, 52, 233, 27, 118, 50, 184, 74, 187, 149, 144, 208,
		68, 67, 241, 19, 119, 240, 84, 172, 163, 235, 120, 10, 227, 52, 4, 82,
		167, 248, 169, 252, 209, 201, 18, 16, 98, 245, 152, 173, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	Permutator1 = new(Permutator).New(256, []byte{
		248, 250, 32, 91, 122, 166, 115, 61, 178, 111, 37, 35, 82, 167, 157, 66,
		22, 65, 47, 1, 195, 182, 190, 73, 19, 218, 237, 76, 140, 155, 18, 11,
		30, 207, 105, 49, 230, 83, 10, 251, 52, 136, 99, 212, 108, 154, 113, 41,
		185, 44, 102, 226, 135, 165, 94, 27, 6, 177, 162, 161, 209, 200, 33, 23,
		197, 120, 71, 249, 125, 244, 217, 38, 0, 128, 95, 80, 214, 254, 163, 203,
		180, 137, 100, 235, 16, 58, 78, 173, 3, 118, 148, 191, 15, 7, 149, 219,
		39, 129, 75, 158, 224, 92, 147, 144, 236, 60, 29, 9, 252, 51, 139, 97,
		43, 87, 193, 222, 85, 223, 127, 153, 192, 13, 143, 70, 151, 123, 211, 72,
		93, 194, 229, 42, 17, 146, 196, 107, 215, 112, 231, 21, 124, 86, 132, 238,
		26, 189, 98, 172, 201, 175, 188, 88, 114, 5, 25, 64, 103, 246, 45, 57,
		109, 63, 81, 62, 204, 106, 179, 199, 116, 141, 186, 121, 84, 210, 79, 156,
		216, 14, 253, 233, 46, 55, 138, 34, 74, 20, 245, 89, 198, 133, 239, 142,
		234, 24, 176, 213, 169, 241, 90, 232, 28, 240, 183, 227, 56, 247, 160, 152,
		202, 4, 159, 104, 187, 31, 174, 48, 168, 67, 40, 50, 134, 228, 181, 170,
		225, 126, 54, 36, 220, 208, 150, 117, 255, 221, 101, 69, 77, 110, 243, 206,
		130, 59, 205, 242, 184, 164, 131, 12, 2, 119, 96, 171, 53, 68, 8, 145})

	// BigZero - the big int value for zero.
	BigZero = big.NewInt(0)
	// BigOne - the big int value for one.
	BigOne = big.NewInt(1)
)

// CipherBlock is the data processed by the crypters (rotors and permutators).
// It consistes of the length in bytes to process and the (32 bytes of) data to
// process.
type CipherBlock struct {
	Length      int8
	CipherBlock [CipherBlockBytes]byte
}

// String formats a string representing the CipherBlock.
func (cblk *CipherBlock) String() string {
	var output bytes.Buffer
	output.WriteString("CipherBlock: ")
	output.WriteString(fmt.Sprintf("\t     Length: %d\n", cblk.Length))
	output.WriteString(fmt.Sprintf("\tCipherBlock:\t% X\n", cblk.CipherBlock[0:16]))
	output.WriteString(fmt.Sprintf("\t\t\t% X", cblk.CipherBlock[16:]))
	return output.String()
}

// Crypter interface
type Crypter interface {
	Update(*Rand)                                           // function to update the rotor/permutator
	SetIndex(*big.Int)                                      // setter for the index value
	Index() *big.Int                                        // getter for the index value
	ApplyF(*[CipherBlockBytes]byte) *[CipherBlockBytes]byte // encryption function
	ApplyG(*[CipherBlockBytes]byte) *[CipherBlockBytes]byte // decryption function
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
func (cntr *Counter) ApplyF(blk *[CipherBlockBytes]byte) *[CipherBlockBytes]byte {
	cntr.index.Add(cntr.index, BigOne)
	return blk
}

// ApplyG - this function does nothing for a Counter during decryption.
func (cntr *Counter) ApplyG(blk *[CipherBlockBytes]byte) *[CipherBlockBytes]byte {
	return blk
}

// SubBlock -  subtracts (not XOR) the key from the data to be decrypted
func SubBlock(blk, key *[CipherBlockBytes]byte) *[CipherBlockBytes]byte {
	var p int
	for idx, val := range *blk {
		p = p + int(val) - int(key[idx])
		blk[idx] = byte(p & 0xFF)
		p = p >> BitsPerByte
	}
	return blk
}

// AddBlock - adds (not XOR) the data to be encrypted with the key.
func AddBlock(blk, key *[CipherBlockBytes]byte) *[CipherBlockBytes]byte {
	var p int
	for i, v := range *blk {
		p += int(v) + int(key[i])
		blk[i] = byte(p & 0xFF)
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
			if inp.Length <= 0 {
				right <- inp
				ecm = nil
				break
			}
			inp.CipherBlock = *ecm.ApplyF(&inp.CipherBlock)
			right <- inp
		}
	}(ecm, left, right)
	return right
}

// DecryptMachine - set up a rotor, permutator, or counter to decrypt a block
// read from the left (input channel) and send it out on the right (output channel)
func DecryptMachine(ecm Crypter, left chan CipherBlock) chan CipherBlock {
	right := make(chan CipherBlock)
	go func(ecm Crypter, left chan CipherBlock, right chan CipherBlock) {
		defer close(right)
		for {
			inp := <-left
			if inp.Length <= 0 {
				right <- inp
				ecm = nil
				break
			}
			inp.CipherBlock = *ecm.ApplyG(&inp.CipherBlock)
			right <- inp
		}
	}(ecm, left, right)
	return right
}

// CreateEncryptMachine - Chain the encryption machines together, using channels to pass
//
//	the data to be encrypted to the individual encryption machines.
//	The data is entered in the 'left' channel and the encrypted data
//	is read from the 'right' channel.
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
//
//	channels to pass the data to be decrypted to the individual
//	decryption machines.  The encrypted data is entered in the 'left'
//	channel and the plaintext data is read from the 'right' channel.
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
