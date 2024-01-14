// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

// Define the rotor used by tntengine.

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
)

// Rotor is the type of a TNT rotor
type Rotor struct {
	Size    int    // the size in bits for this rotor
	Start   int    // the initial starting position of the rotor
	Step    int    // the step size in bits for this rotor
	Current int    // the current position of this rotor
	Rotor   []byte // the rotor
}

// New fills the (empty) rotor r with the given size, start, step and rotor data.
func (r *Rotor) New(size, start, step int, rotor []byte) *Rotor {
	r.Start, r.Current = start, start
	r.Size = size
	r.Step = step
	r.Rotor = append([]byte(nil), rotor...)
	r.sliceRotor()
	return r
}

// Update rotor r with a new start, step and (psudo)-random data.
func (r *Rotor) Update(random *Rand) {
	// Get start and step of the new rotor
	start := random.Intn(r.Size) // 0 <= start < r.Size
	// The step must be in the range 0 < step < r.Size.  If it happens to be equal
	// to 0 or r.Size, then the rotor will not step (it will always be equal to start
	// each time it steps)
	r.Step = random.Intn(r.Size-1) + 1 // 0 < step < r.Size
	// Fill the rotor with random data using tntengine Rand function to generate the
	// random data to fill the rotor.  As an optimization, set the random.idx equal to
	// CypherBlockSize.  This will ensure that the random.Read() will only call the
	// encryption function one time per block instead of up to two times per block.
	random.idx = CipherBlockBytes
	blk := make(CipherBlock, CipherBlockBytes)
	CipherBlocksToRead := (r.Size + 7) / CipherBlockSize
	j := 0
	for i := 0; i < CipherBlocksToRead; i++ {
		random.Read(blk)
		copy(r.Rotor[j:], blk[:])
		j += CipherBlockBytes
	}
	r.Start, r.Current = start, start
	r.sliceRotor() // Append the first 256 bits of the rotor to the end of the rotor
}

// sliceRotor appends the first 256 bits of the rotor to the end of the rotor.
func (r *Rotor) sliceRotor() {
	var size, sBlk, sBit, Rshift, Lshift uint
	var i int
	size = uint(r.Size)
	sBlk = size >> 3
	sBit = size & 7
	Rshift = 8 - sBit
	Lshift = sBit
	fmt.Fprintf(os.Stderr, "size: %d sBlk: %d sBit: %d Rshift: %d\n", size, sBlk, sBit, Rshift)
	if sBit == 0 {
		copy(r.Rotor[sBlk:], r.Rotor[0:CipherBlockBytes])
	} else {
		// The copy appending will be done at the byte level instead of the bit level
		// so that we only loop 32 times instead of 256 times.
		for i = 0; i < CipherBlockBytes; i++ {
			r.Rotor[sBlk] &= (0xff >> Rshift)       // Clear out the bits that will be replaced
			r.Rotor[sBlk] |= (r.Rotor[i] << Lshift) // and add in the bits from the beginning of the rotor
			sBlk++
			r.Rotor[sBlk] = (r.Rotor[i] >> Rshift) // Seed the next byte at the end with the remaining bits from the beginning byte.
		}
	}
}

// SetIndex positions the rotor to the position it would be in after
// processing idx number of blocks.
func (r *Rotor) SetIndex(idx *big.Int) {
	// Special case if idx == 0
	if idx.Sign() == 0 {
		r.Current = r.Start
	} else {
		// Calculate the new r.Current:
		// r.Current = mod(((idx * r.Step) + r.Start), r.Size) + r.Start
		p := new(big.Int)
		q := new(big.Int)
		rem := new(big.Int)
		p = p.Mul(idx, new(big.Int).SetInt64(int64(r.Step)))
		p = p.Add(p, new(big.Int).SetInt64(int64(r.Start)))
		_, rem = q.DivMod(p, new(big.Int).SetInt64(int64(r.Size)), rem)
		r.Current = int(rem.Int64())
	}
}

// Always return nil since the block count is not tracked for rotors.
func (r *Rotor) Index() *big.Int {
	return nil
}

// Get the number of bits in the CipherBlock from rotor r.
func (r *Rotor) getRotorBlock(blk CipherBlock) CipherBlock {
	// This code handles short blocks to accomadate file lenghts
	// that are not multiples of "CipherBlockBytes"
	ress := make([]byte, len(blk))
	rotor := r.Rotor
	sBit := r.Current & 7
	bIdx := r.Current >> 3
	// The copy operates at the byte level instead of the bit level
	// so we only loop 32 times instead of 256 times.
	if sBit == 0 {
		copy(ress, rotor[bIdx:])
	} else {
		sLeft := 8 - sBit
		for bCnt := 0; bCnt < len(ress); bCnt++ {
			ress[bCnt] = rotor[bIdx]>>sBit |
				(rotor[bIdx+1] << sLeft)
			bIdx++
		}
	}
	// Step the rotor to its new position.
	r.Current = (r.Current + r.Step) % r.Size
	return ress
}

// ApplyF encrypts the given block of data using the rotor r.
func (r *Rotor) ApplyF(blk CipherBlock) CipherBlock {
	// Add (not XOR) the rotor bits to the bits in the input block.
	return AddBlock(blk, r.getRotorBlock(blk))
}

// ApplyG decrypts the given block of data using the rotor r.
func (r *Rotor) ApplyG(blk CipherBlock) CipherBlock {
	// Subtract (not XOR) the rotor bits from the bits in the input block.
	return SubBlock(blk, r.getRotorBlock(blk))
}

// String converts a Rotor to a string representation of the Rotor.
func (r *Rotor) String() string {
	var output bytes.Buffer
	rotorLen := len(r.Rotor)
	output.WriteString(fmt.Sprintf("new(Rotor).New(%d, %d, %d, []byte{\n",
		r.Size, r.Start, r.Step))
	for i := 0; i < rotorLen; i += 16 {
		output.WriteString("\t")
		if i+16 < rotorLen {
			for _, k := range r.Rotor[i : i+15] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d,", r.Rotor[i+15]))
		} else {
			l := len(r.Rotor[i:])
			for _, k := range r.Rotor[i : i+l-1] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d})", r.Rotor[i+l-1]))
		}
		output.WriteString("\n")
	}

	return output.String()
}
