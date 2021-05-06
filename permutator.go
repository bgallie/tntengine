// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"bytes"
	"fmt"
	"math/big"
	"sync"
)

// Cycle describes a cycle for the permutator so it can adjust the permutation
// table used to permutate the block.  TNT2 currently uses 4 cycles to rearrange
// Randp into bitPerm
type Cycle struct {
	Start   int // The starting point (into randp) for this cycle.
	Length  int // The length of the cycle.
	Current int // The point in the cycle [0 .. cycle.length-1] to start
}

// Permutator is a type that defines a permutation cryptor in TNT2.
type Permutator struct {
	CurrentState  int                   // Current number of cycles for this permutator.
	MaximalStates int                   // Maximum number of cycles this permutator can have before repeating.
	Cycles        []Cycle               // Cycles ordered by the current permutation.
	Randp         []byte                // Values 0 - 255 in a random order.
	bitPerm       [CypherBlockSize]byte // Permutation table created from randp.
}

// New creates a permutator and initializes it
func NewPermutator(cycleSizes []int, randp []byte) *Permutator {
	var p Permutator
	p.Randp = randp
	p.Cycles = make([]Cycle, len(cycleSizes))

	for i := range cycleSizes {
		p.Cycles[i].Length = cycleSizes[i]
		p.Cycles[i].Current = 0
		// Adjust the start to reflect the lenght of the previous cycles
		if i == 0 { // no previous cycle so start at 0
			p.Cycles[i].Start = 0
		} else {
			p.Cycles[i].Start = p.Cycles[i-1].Start + p.Cycles[i-1].Length
		}
	}

	p.CurrentState = 0
	p.MaximalStates = p.Cycles[0].Length

	for i := 1; i < len(p.Cycles); i++ {
		p.MaximalStates *= p.Cycles[i].Length
	}

	p.cycle()
	return &p
}

// Update the permutator with a new initialCycles and Randp
func (p *Permutator) Update(cycleSizes []int, randp []byte) {
	p.Randp = randp
	p.Cycles = make([]Cycle, len(cycleSizes))

	for i := range cycleSizes {
		p.Cycles[i].Length = cycleSizes[i]
		p.Cycles[i].Current = 0
		// Adjust the start to reflect the lenght of the previous cycles
		if i == 0 { // no previous cycle so start at 0
			p.Cycles[i].Start = 0
		} else {
			p.Cycles[i].Start = p.Cycles[i-1].Start + p.Cycles[i-1].Length
		}
	}

	p.CurrentState = 0
	p.MaximalStates = p.Cycles[0].Length

	for i := 1; i < len(p.Cycles); i++ {
		p.MaximalStates *= p.Cycles[i].Length
	}

	p.cycle()
}

// cycle bitPerm to it's next state.
func (p *Permutator) nextState() {
	for _, val := range p.Cycles {
		val.Current = (val.Current + 1) % val.Length
	}

	p.CurrentState = (p.CurrentState + 1) % p.MaximalStates
	p.cycle()
}

// cycle will create a new bitPerm from Randp based on the current cycle.
func (p *Permutator) cycle() {
	var wg sync.WaitGroup

	for _, val := range p.Cycles {
		wg.Add(1)

		go func(cycle []byte, sIdx int, length int) {
			defer wg.Done()

			for _, val := range cycle {
				p.bitPerm[val] = p.Randp[cycle[sIdx]]
				sIdx = (sIdx + 1) % length
			}
		}(p.Randp[val.Start:val.Start+val.Length], val.Current, val.Length)
	}

	wg.Wait()
}

// SetIndex - set the Permutator to the state it would be in after encoding 'idx - 1' blocks
// of data.
func (p *Permutator) SetIndex(idx *big.Int) {
	q := new(big.Int)
	r := new(big.Int)
	q, r = q.DivMod(idx, big.NewInt(int64(p.MaximalStates)), r)
	p.CurrentState = int(r.Int64())

	for _, val := range p.Cycles {
		val.Current = p.CurrentState % val.Length
	}

	p.cycle()
}

// Index returns the current index of the cryptor.  For permeutators, this
// returns nil.
func (p *Permutator) Index() *big.Int {
	return nil
}

// ApplyF performs forward permutation on the 32 byte block of data.
func (p *Permutator) ApplyF(blk *[CypherBlockBytes]byte) *[CypherBlockBytes]byte {
	var res [CypherBlockBytes]byte
	blks := blk[:]
	ress := res[:]

	for i, v := range p.bitPerm {
		if GetBit(blks, uint(i)) {
			SetBit(ress, uint(v))
		}
	}

	p.nextState()
	*blk = res
	return blk
}

// ApplyG performs the reverse permutation on the 32 byte block of data.
func (p *Permutator) ApplyG(blk *[CypherBlockBytes]byte) *[CypherBlockBytes]byte {
	var res [CypherBlockBytes]byte
	blks := blk[:]
	ress := res[:]

	for i, v := range p.bitPerm {
		if GetBit(blks, uint(v)) {
			SetBit(ress, uint(i))
		}
	}

	p.nextState()
	*blk = res
	return blk
}

// String formats a string representing the permutator (as Go source code).
func (p *Permutator) String() string {
	var output bytes.Buffer
	output.WriteString(fmt.Sprint("permutator.New([]int{"))

	for _, v := range p.Cycles[0 : NumberPermutationCycles-1] {
		output.WriteString(fmt.Sprintf("%d, ", v.Length))
	}

	output.WriteString(fmt.Sprintf("%d}, []byte{\n", p.Cycles[NumberPermutationCycles-1].Length))

	for i := 0; i < CypherBlockSize; i += 16 {
		output.WriteString("\t")

		if i != (CypherBlockSize - 16) {
			for _, k := range p.Randp[i : i+16] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
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
