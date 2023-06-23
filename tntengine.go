// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

// Define the tntengine type and it's methods

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/bgallie/jc1"
	"golang.org/x/crypto/sha3"
)

var (
	proFormaRotors = []*Rotor{
		// Define the proforma rotors used to create the actual rotors to use.
		new(Rotor).New(1789, 1065, 1499, []byte{
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
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		new(Rotor).New(1787, 1624, 249, []byte{
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
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		new(Rotor).New(1783, 1056, 1256, []byte{
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
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		new(Rotor).New(1777, 210, 241, []byte{
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
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		new(Rotor).New(1759, 955, 1559, []byte{
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
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		new(Rotor).New(1753, 370, 362, []byte{
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
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
	}
	// Define the proforma permutators used to create the actual permutators to use.
	proFormPermutators = []*Permutator{
		new(Permutator).New(256, []byte{
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
			130, 59, 205, 242, 184, 164, 131, 12, 2, 119, 96, 171, 53, 68, 8, 145}),
	}
	Rotor1               = proFormaRotors[0]
	Rotor2               = proFormaRotors[1]
	Rotor3               = proFormaRotors[2]
	Rotor4               = proFormaRotors[3]
	Rotor5               = proFormaRotors[4]
	Rotor6               = proFormaRotors[5]
	Permutator1          = proFormPermutators[0]
	counter     *Counter = new(Counter)
	jc1Key      *jc1.UberJc1
)

// TntEngine type defines the encryption/decryption machine (rotors and
// permutators).
type TntEngine struct {
	engineType    string // "E)ncrypt" or "D)ecrypt"
	engine        []Crypter
	left, right   chan CipherBlock
	cntrKey       string
	maximalStates *big.Int
}

// Left is a getter that returns the input channel for the TntEngine.
func (e *TntEngine) Left() chan CipherBlock {
	return e.left
}

// Right is a getter that returns the output channel for the TntEngine.
func (e *TntEngine) Right() chan CipherBlock {
	return e.right
}

// CounterKey is a getter that returns the SHAKE256 hash for the secret key.
// This is used to set/retrieve that next block to use in encrypting data.
func (e *TntEngine) CounterKey() string {
	return e.cntrKey
}

// Index is a getter that returns the block number of the next block to be
// encrypted.
func (e *TntEngine) Index() (cntr *big.Int) {
	if len(e.engine) != 0 {
		machine := e.engine[len(e.engine)-1]
		switch machine.(type) {
		default:
			cntr = BigZero
		case *Counter:
			cntr = machine.Index()
		}
	}

	return
}

// SetIndex is a setter function that sets the rotors and permutators so the
// the TntEngine will be ready start encrypting/decrypting at the correct block.
func (e *TntEngine) SetIndex(iCnt *big.Int) {
	for _, machine := range e.engine {
		machine.SetIndex(new(big.Int).Set(iCnt))
	}
}

// SetEngineType is a setter function that sets the engineType [D)ecrypt or E)crypt]
// of the TntEngine.
func (e *TntEngine) SetEngineType(engineType string) {
	switch string(strings.TrimSpace(engineType)[0]) {
	case "d", "D":
		e.engineType = "D"
	case "e", "E":
		e.engineType = "E"
	default:
		log.Fatalf("Missing or incorrect TntEngine engineType: [%s]", engineType)
	}
}

// Engine is a getter function that returns a slice containing the rotors and
// permutators for the TntEngine.
func (e *TntEngine) Engine() []Crypter {
	return e.engine
}

// EngineType is a getter function that returns the engine type of the TntMachine.
func (e *TntEngine) EngineType() string {
	return e.engineType
}

// MaximalStates is a getter function that returns maximum number of states that the
// engine can be in before repeating.
func (e *TntEngine) MaximalStates() *big.Int {
	return e.maximalStates
}

// Init will initialize the TntEngine generating new Rotors and Permutators using
// the proForma rotors and permutators in complex way, updating the rotors and
// permutators in place.
func (e *TntEngine) Init(secret []byte) {
	jc1Key = new(jc1.UberJc1).New(secret)
	// Create an ecryption machine based on the proForma rotors and permutators.
	e.engine = *createProFormaMachine()
	e.left, e.right = createEncryptMachine(e.engine...)
	// Get a SHA-3 hash of the encryption key.  This is used as a key to store
	// the count of blocks already encrypted to use as a starting point for the
	// encryption of the next message.
	k := make([]byte, 1024)
	blk := make(CipherBlock, CipherBlockBytes)
	h := blk[:]
	d := sha3.NewShake256()
	d.Write(jc1Key.XORKeyStream(k))
	d.Read(h)
	// Encrypt the hash starting at block 1234567890 (no good reason for this number)
	// to make it specific to the proForma machine used.
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	e.SetIndex(iCnt)
	e.left <- blk
	blk = <-e.right
	e.cntrKey = hex.EncodeToString(blk[:])
	e.SetIndex(BigZero)
	// Create a random number function [func(max int) int] that uses psudo-
	// random data generated the proforma encryption machine.
	random := new(Rand).New(e)
	// Update the rotors and permutators in a very non-linear fashion.
	e.maximalStates = new(big.Int).Set(BigOne)
	for _, machine := range e.engine {
		switch v := machine.(type) {
		default:
			fmt.Fprintf(os.Stderr, "Unknown machine: %v\n", v)
		case *Rotor:
			machine.Update(random)
			e.maximalStates = e.maximalStates.Mul(e.maximalStates, big.NewInt(int64(machine.(*Rotor).Size)))
		case *Permutator:
			machine.Update(random)
			e.maximalStates = e.maximalStates.Mul(e.maximalStates, big.NewInt(int64(machine.(*Permutator).MaximalStates)))
		case *Counter:
			machine.SetIndex(BigZero)
		}
	}
	// Now that we have created the new rotors and permutators from the proform
	// machine, populate the TntEngine with them.
	newMachine := make([]Crypter, 9)
	cnt := copy(newMachine, e.engine)
	counter.SetIndex(BigZero)
	newMachine[cnt] = counter
	e.engine = newMachine
}

// BuildCiperMachine will create a "machine" to encrypt or decrypt data sent to the
// left channel and outputed on the right channel for the TntEngine.  The engineType
// determines weither a encrypt machine or a decrypt machine will be created.
func (e *TntEngine) BuildCipherMachine() {
	switch e.engineType {
	case "D":
		e.left, e.right = createDecryptMachine(e.engine...)
	case "E":
		e.left, e.right = createEncryptMachine(e.engine...)
	default:
		log.Fatalf("Missing or incorrect TntEngine engineType: [%s]", e.engineType)

	}
}

// CloseCipherMachine will close down the cipher machine by exiting the go function
// that performs the encryption/decryption using the individual rotors/permutators.
// This is done by passing the CipherMachine a CypherBlock with a length of zero (0).
func (e *TntEngine) CloseCipherMachine() {
	blk := new(CipherBlock)
	e.Left() <- *blk
	<-e.Right()
}

// createProFormaMachine initializes the proForma machine used to create the
// TNT encryption machine.  If the machineFileName is not empty then the
// proForma machine is loaded from that file, else the hardcoded rotors and
// permutators are used to initialize the proForma machine.
func createProFormaMachine() *[]Crypter {
	newMachine := make([]Crypter, 8)
	// Create the proforma encryption machine.  The layout of the machine is:
	// 		rotor, rotor, permutator, rotor, rotor, permutator, rotor, rotor
	// Note: The two permutators are identical which simulates the original 'C' code where the permutator
	//		 is used twice before it is cycled to it's next state.
	// Create the ProFormaMachine by making a copy of the hardcoded proforma rotors and permutators.
	// This resolves an issue running tests where TntEngine.Init() is called multiple times which
	// caused a failure on the second call.
	newMachine[0] = new(Rotor).New(Rotor1.Size, Rotor1.Start, Rotor1.Step, append([]byte(nil), Rotor1.Rotor...))
	newMachine[1] = new(Rotor).New(Rotor2.Size, Rotor2.Start, Rotor2.Step, append([]byte(nil), Rotor2.Rotor...))
	newMachine[2] = new(Permutator).New(Permutator1.Cycle.Length, append([]byte(nil), Permutator1.Randp...))
	newMachine[3] = new(Rotor).New(Rotor3.Size, Rotor3.Start, Rotor3.Step, append([]byte(nil), Rotor3.Rotor...))
	newMachine[4] = new(Rotor).New(Rotor4.Size, Rotor4.Start, Rotor4.Step, append([]byte(nil), Rotor4.Rotor...))
	newMachine[5] = new(Permutator).New(Permutator1.Cycle.Length, append([]byte(nil), Permutator1.Randp...))
	newMachine[6] = new(Rotor).New(Rotor5.Size, Rotor5.Start, Rotor5.Step, append([]byte(nil), Rotor5.Rotor...))
	newMachine[7] = new(Rotor).New(Rotor6.Size, Rotor6.Start, Rotor6.Step, append([]byte(nil), Rotor6.Rotor...))

	return &newMachine
}
