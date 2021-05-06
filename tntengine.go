// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

// Package tntengine - define TntEngine type and it's methods
package tntengine

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/bgallie/jc1"
	"golang.org/x/crypto/sha3"
)

var (
	counter         *Counter = new(Counter)
	proFormaMachine []Crypter
	rotorSizes      []int
	rotorSizesIndex int
	cycleSizes      []int
	cycleSizesIndex int
	proFormaFile    os.File
	random          *Rand
)

// TntEngine type defines the encryption/decryption machine (rotors and
// permutators).
type TntEngine struct {
	engineType  string // "E)ncrypt" or "D)ecrypt"
	engine      []Crypter
	left, right chan CypherBlock
	jc1Key      *jc1.UberJc1
	cntrKey     string
}

// Left is a getter that returns the input channel for the TntEngine.
func (e *TntEngine) Left() chan CypherBlock {
	return e.left
}

// Right is a getter that returns the output channel for the TntEngine.
func (e *TntEngine) Right() chan CypherBlock {
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
		machine.SetIndex(iCnt)
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

// Init will initialize the TntEngine generating new Rotors and Permutators using
// the proForma rotors and permutators in complex way, updating the rotors and
// permutators in place.
func (e *TntEngine) Init(secret []byte, proFormaFileName string) {
	e.jc1Key = jc1.NewUberJc1(secret)
	// Create an ecryption machine based on the proForma rotors and permutators.
	e.engine = *createProFormaMachine(proFormaFileName)
	e.left, e.right = createEncryptMachine(e.engine...)
	// Create a random number function [func(max int) int] that uses psudo-
	// random data generated the proforma encryption machine.
	random := NewRand(e)
	// Get a SHA-3 hash of the encryption key.  This is used as a key to store
	// the count of blocks already encrypted to use as a starting point for the
	// encryption of the next message.
	k := make([]byte, 1024)
	h := make([]byte, 32)
	d := sha3.NewShake256()
	d.Write(e.jc1Key.XORKeyStream(k))
	d.Read(h)
	e.cntrKey = hex.EncodeToString(h)
	// Create a permutaion of the rotor indices to allow picking the rotors in
	// a random order based on the key.
	rotorSizes = random.Perm(len(RotorSizes))
	// Create a permutaion of cycle sizes indices to allow picking the cycle
	// sizes in a random order based on the key.
	cycleSizes = random.Perm(len(CycleSizes))
	// Update the rotors and permutators in a very non-linear fashion.
	for pfCnt, machine := range e.engine {
		switch v := machine.(type) {
		default:
			fmt.Fprintf(os.Stderr, "Unknown machine: %v\n", v)
		case *Rotor:
			updateRotor(machine.(*Rotor), random)
		case *Permutator:
			p := new(Permutator)
			updatePermutator(p, random)
			e.engine[pfCnt] = p
		case *Counter:
			machine.(*Counter).SetIndex(BigZero)
		}
	}
	// Now that we have created the new rotors and permutators from the proform
	// machine, populate the TntEngine with them.
	newMachine := make([]Crypter, 9, 9)
	machineOrder := random.Perm(len(e.engine))
	for idx, val := range machineOrder {
		newMachine[idx] = e.engine[val]
	}
	counter.SetIndex(BigZero)
	newMachine[len(newMachine)-1] = counter
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

// createProFormaMachine initializes the proForma machine used to create the
// TNT2 encryption machine.  If the machineFileName is not empty then the
// proForma machine is loaded from that file, else the hardcoded rotors and
// permutators are used to initialize the proForma machine.
func createProFormaMachine(machineFileName string) *[]Crypter {
	var newMachine []Crypter
	if len(machineFileName) == 0 {
		// log.Println("Using built in proforma rotors and permutators")
		// Create the proforma encryption machine.  The layout of the machine is:
		// 		rotor, rotor, permutator, rotor, rotor, permutator, rotor, rotor
		newMachine = []Crypter{
			Rotor1, Rotor2, Permutator1,
			Rotor3, Rotor4, Permutator2,
			Rotor5, Rotor6}
	} else {
		// log.Printf("Using proforma rotors and permutators from %s\n", machineFileName)
		in, err := os.Open(machineFileName)
		checkFatal(err)
		jDecoder := json.NewDecoder(in)
		// Create the proforma encryption machine from the given proforma machine file.
		// The layout of the machine is:
		// 		rotor, rotor, permutator, rotor, rotor, permutator, rotor, rotor
		var rotor1, rotor2, rotor3, rotor4, rotor5, rotor6 *Rotor
		var permutator1, permutator2 *Permutator
		newMachine = []Crypter{rotor1, rotor2, permutator1, rotor3, rotor4, permutator2, rotor5, rotor6}

		for cnt, machine := range newMachine {
			switch v := machine.(type) {
			default:
				fmt.Fprintf(os.Stderr, "Unknown machine: %v\n", v)
			case *Rotor:
				r := new(Rotor)
				err = jDecoder.Decode(&r)
				checkFatal(err)
				newMachine[cnt] = r
			case *Permutator:
				p := new(Permutator)
				err = jDecoder.Decode(&p)
				checkFatal(err)
				newMachine[cnt] = p
			}
		}
	}

	return &newMachine
}

// updateRotor will update the given (proforma) rotor in place using (psudo-
// random) data generated by the TNT2 encrytption algorithm using the proForma
// rotors and permutators.
func updateRotor(r *Rotor, random *Rand) {
	// Get size, start and step of the new rotor
	rotorSize := RotorSizes[rotorSizes[rotorSizesIndex]]
	rotorSizesIndex = (rotorSizesIndex + 1) % len(RotorSizes)
	start := random.Intn(rotorSize)
	step := random.Intn(rotorSize)

	// blkCnt is the total number of bytes needed to hold rotorSize bits + a slice of 256 bits
	blkCnt := (((rotorSize + CypherBlockSize + 7) / 8) + 31) / 32
	// blkBytes is the number of bytes rotor r needs to increase to hold the new rotor.
	blkBytes := (blkCnt * 32) - len(r.Rotor)
	// Adjust the size of r.Rotor to match the new rotor size.
	adjRotor := make([]byte, blkBytes)
	r.Rotor = append(r.Rotor, adjRotor...)
	// Fill the rotor with random data using TNT2 encryption to generate the
	// random data by encrypting the next 32 bytes of data from the uberJC1
	// algorithm until the next rotor is filled.
	random.Read(r.Rotor)

	// update the rotor with the new size, start, and step and slice the first
	// 256 bits of the rotor to the end of the rotor.
	r.Update(rotorSize, start, step)
}

// updatePermutator will update the given (proForma) permutator in place using
// (psudo-random) data generated by the TNT2 encrytption algorithm using the
// proForma rotors and permutators.
func updatePermutator(p *Permutator, random *Rand) {
	var randp [CypherBlockSize]byte
	// Create a table of byte values [0...255] in a random order
	for i, val := range random.Perm(CypherBlockSize) {
		randp[i] = byte(val)
	}
	// Chose a CycleSizes and randomize order of the values
	length := len(CycleSizes[cycleSizesIndex])
	cycles := make([]int, length, length)
	randi := random.Perm(length)
	for idx, val := range randi {
		cycles[idx] = CycleSizes[cycleSizes[cycleSizesIndex]][val]
	}
	p.Update(cycles, randp[:])
	cycleSizesIndex = (cycleSizesIndex + 1) % len(CycleSizes)
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
