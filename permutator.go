// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.
package tntengine

// Define a permutator used in tntengine

import (
	"bytes"
	"fmt"
)

const (
	NumberPermutationCycles int = 1
)

// Cycle describes a cycle for the permutator so it can adjust the permutation
// table used to permutate the block.  TNT currently uses a single cycle to
// rearrange Randp into bitPerm
type Cycle struct {
	Start   int // The starting point (into randp) for this cycle.
	Length  int // The length of the cycle.
	Current int // The point in the cycle [0 .. cycle.length-1] to start
}

// Permutator is a type that defines a permutation cryptor in TNT.
type Permutator struct {
	CurrentState  int                   // Current number of cycles for this permutator.
	MaximalStates int                   // Maximum number of cycles this permutator can have before repeating.
	Cycle         Cycle                 // Cycles ordered by the current permutation.
	Randp         []byte                // Values 0 - 255 in a random order.
	bitPerm       [CipherBlockSize]byte // Permutation table created from Randp.
}

// New creates a permutator and initializes it
func (p *Permutator) New(cycleSize int, randp []byte) *Permutator {
	p.Randp = append([]byte(nil), randp...)
	p.Cycle.Length = cycleSize
	p.Cycle.Current = 0
	p.Cycle.Start = 0
	// New permutators starts with a current state of 0.
	p.CurrentState = 0
	// Calculate the maximum number of states the permutator can take.
	p.MaximalStates = p.Cycle.Length
	p.cycle()
	return p
}

// Update will update the given (proForma) permutator in place using
// (psudo-)random data generated by the TNT encrytption engine Rand
// object.
func (p *Permutator) Update(random *Rand) {
	// Create a table of byte values [0...255] in a random order
	for i, val := range random.Perm(CipherBlockSize) {
		p.Randp[i] = byte(val)
	}
	// The original TNT program used a single cycle of 256
	p.Cycle.Length = 256
	p.Cycle.Start = 0
	p.Cycle.Current = 0
	// updated permutators start with a current state of 0
	p.CurrentState = 0
	// Calculate the maximum number of states the permutator can take.
	p.MaximalStates = p.Cycle.Length
	// Cycle the permutation table to it's starting state
	p.cycle()
}

// Cycle bitPerm to it's next state.
func (p *Permutator) nextState() {
	p.Cycle.Current = (p.Cycle.Current + 1) % p.Cycle.Length
	p.CurrentState = (p.CurrentState + 1) % p.MaximalStates
	p.cycle()
}

// cycle will create a new bitPerm from Randp based on the current cycle.
// Note that tntengine only uses a single cycle of length 256.
func (p *Permutator) cycle() {
	cycle := p.Randp[:] // there is only 1 cycle so use all of p.Randp
	sIdx := p.Cycle.Current
	length := p.Cycle.Length
	for _, val := range cycle {
		p.bitPerm[val] = p.Randp[cycle[sIdx]]
		sIdx = (sIdx + 1) % length
	}
}

// SetIndex - set the Permutator to the state it would be in after encoding 'idx - 1' blocks
// of data.
func (p *Permutator) SetIndex(idx *Counter) {
	p.CurrentState = int(idx.Mod(uint64(p.MaximalStates)))
	p.Cycle.Current = p.CurrentState
	p.cycle()
}

// Index returns the current index of the cryptor.  For permeutators, this
// returns nil.
func (p *Permutator) Index() *Counter {
	return nil
}

// ApplyF performs forward permutation on the 32 byte block of data.
// Note: if the length of the incoming block is less than CipherBlockBytes
// in length, then the permutation is not applied.  This allows files whose
// length is not a multiple of 32 bytes to be correctly enrypted/decrypted.
func (p *Permutator) ApplyF(blk CipherBlock) CipherBlock {
	if len(blk) == CipherBlockBytes {
		ress := make([]byte, CipherBlockBytes)
		for i, v := range p.bitPerm {
			if GetBit(blk, uint(i)) {
				SetBit(ress, uint(v))
			}
		}
		p.nextState()
		blk = ress
	}
	return blk
}

// ApplyG performs the reverse permutation on the 32 byte block of data.
// Note: if the length of the incoming block is less than CipherBlockBytes
// in length, then the permutation is not applied.  This allows files whose
// length is not a multiple of 32 bytes to be correctly enrypted/decrypted.
func (p *Permutator) ApplyG(blk CipherBlock) CipherBlock {
	if len(blk) == CipherBlockBytes {
		ress := make([]byte, CipherBlockBytes)
		for i, v := range p.bitPerm {
			if GetBit(blk, uint(v)) {
				SetBit(ress, uint(i))
			}
		}
		p.nextState()
		blk = ress
	}
	return blk
}

// String formats a string representing the permutator (as Go source code).
func (p *Permutator) String() string {
	var output bytes.Buffer
	output.WriteString("new(Permutator).New(")
	output.WriteString(fmt.Sprintf("%d, []byte{\n", p.Cycle.Length))
	for i := 0; i < CipherBlockSize; i += 16 {
		output.WriteString("\t")
		if i != (CipherBlockSize - 16) {
			for _, k := range p.Randp[i : i+15] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d,", p.Randp[i+15]))
		} else {
			for _, k := range p.Randp[i : i+15] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d})", p.Randp[i+15]))
		}
		output.WriteString("\n")
	}
	return output.String()
}
