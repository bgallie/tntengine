// Package tntEngine - define TntEngine type and it's methods
package tntEngine

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/bgallie/jc1"
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

type TntEngine struct {
	engineType  string // "Encrypt" or "Decrypt"
	engine      []Crypter
	left, right chan CypherBlock
	jc1Key      *jc1.UberJc1
	cntrKey     string
}

func (e *TntEngine) Left() chan CypherBlock {
	return e.left
}

func (e *TntEngine) Right() chan CypherBlock {
	return e.right
}

func (e *TntEngine) Key() *jc1.UberJc1 {
	return e.jc1Key
}

func (e *TntEngine) CounterKey() string {
	return e.cntrKey
}

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

func (e *TntEngine) SetIndex(iCnt *big.Int) {
	for _, machine := range e.engine {
		machine.SetIndex(iCnt)
	}
}

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

func (e *TntEngine) Engine() []Crypter {
	return e.engine
}

func (e *TntEngine) EngineType() string {
	return e.engineType
}

func (e *TntEngine) Init(secret []byte, proFormaFileName string) {
	e.jc1Key = jc1.NewUberJc1(secret)
	// Create an ecryption machine based on the proForma rotors and permutators.
	e.engine = *createProFormaMachine(proFormaFileName)
	e.left, e.right = CreateEncryptMachine(e.engine...)
	// Create a random number function [func(max int) int] that uses psudo-
	// random data generated the proforma encryption machine.
	random := NewRand(e)
	// Get a 'checksum' of the encryption key.  This is used as a key to store
	// the count of blocks already encrypted to use as a starting point for the
	// next encryption.
	var blk CypherBlock
	var cksum [CypherBlockBytes]byte
	blk.Length = CypherBlockBytes
	_ = copy(blk.CypherBlock[:], e.jc1Key.XORKeyStream(cksum[:]))
	e.left <- blk
	blk = <-e.right
	e.cntrKey = hex.EncodeToString(blk.CypherBlock[:])
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
	// e.SetEngineType("E")
	// e.left, e.right = CreateEncryptMachine(e.engine...)
	// random = NewRand(e)
}

func (e *TntEngine) BuildCipherMachine() {
	switch e.engineType {
	case "D":
		e.left, e.right = CreateDecryptMachine(e.engine...)
	case "E":
		e.left, e.right = CreateEncryptMachine(e.engine...)
	default:
		log.Fatalf("Missing or incorrect TntEngine engineType: [%s]", e.engineType)

	}
}

/*
	createProFormaMachine initializes the pro-forma machine used to create the
	TNT2 encryption machine.  If the machineFileName is not empty then the
	pro-forma machine is loaded from that file, else the hardcoded rotors and
	permutators are used to initialize the pro-formaa machine.
*/
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

/*
	updateRotor will update the given (proforma) rotor in place using (psudo-
	random) data generated by the TNT2 encrytption algorithm using the pro-forma
	rotors and permutators.
*/
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

/*
	updatePermutator will update the given (proforma) permutator in place using
	(psudo-random) data generated by the TNT2 encrytption algorithm using the
	proforma rotors and permutators.
*/
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
