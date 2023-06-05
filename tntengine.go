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
	counter *Counter = new(Counter)
	jc1Key  *jc1.UberJc1
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
	jc1Key = jc1.NewUberJc1(secret)
	// Create an ecryption machine based on the proForma rotors and permutators.
	e.engine = *createProFormaMachine()
	e.left, e.right = createEncryptMachine(e.engine...)
	// Get a SHA-3 hash of the encryption key.  This is used as a key to store
	// the count of blocks already encrypted to use as a starting point for the
	// encryption of the next message.
	k := make([]byte, 1024)
	blk := *new(CipherBlock)
	blk.Length = 32
	h := blk.CipherBlock[:]
	d := sha3.NewShake256()
	d.Write(jc1Key.XORKeyStream(k))
	d.Read(h)
	// Encrypt the hash starting at block 1234567890 (no good reason for this number)
	// to make it specific to the proForma machine used.
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	e.SetIndex(iCnt)
	e.left <- blk
	blk = <-e.right
	e.cntrKey = hex.EncodeToString(blk.CipherBlock[:])
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
	newMachine[2] = new(Permutator).New(Permutator1.Cycle.Length, Permutator1.Randp)
	newMachine[3] = new(Rotor).New(Rotor3.Size, Rotor3.Start, Rotor3.Step, append([]byte(nil), Rotor3.Rotor...))
	newMachine[4] = new(Rotor).New(Rotor4.Size, Rotor4.Start, Rotor4.Step, append([]byte(nil), Rotor4.Rotor...))
	newMachine[5] = new(Permutator).New(Permutator1.Cycle.Length, Permutator1.Randp)
	newMachine[6] = new(Rotor).New(Rotor5.Size, Rotor5.Start, Rotor5.Step, append([]byte(nil), Rotor5.Rotor...))
	newMachine[7] = new(Rotor).New(Rotor6.Size, Rotor6.Start, Rotor6.Step, append([]byte(nil), Rotor6.Rotor...))

	return &newMachine
}
