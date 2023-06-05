// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

// Define the rotor used by tntengine.

import (
	"bytes"
	"fmt"
	"math/big"
)

// Rotor - the type of the TNT rotor
type Rotor struct {
	Size    int    // the size in bits for this rotor
	Start   int    // the initial starting position of the rotor
	Step    int    // the step size in bits for this rotor
	Current int    // the current position of this rotor
	Rotor   []byte // the rotor
}

// New - creates a new Rotor with the given size, start, step and rotor data.
func (r *Rotor) New(size, start, step int, rotor []byte) *Rotor {
	r.Start, r.Current = start, start
	r.Size = size
	r.Step = step
	r.Rotor = append([]byte(nil), rotor...)
	r.sliceRotor()
	return r
}

// Update - updates the given Rotor with a new start, step and (psudo)
//   - random rotor data.
func (r *Rotor) Update(random *Rand) {
	// Get start and step of the new rotor
	rotorSize := r.Size
	start := random.Intn(rotorSize)
	step := random.Intn(rotorSize-1) + 1
	// Fill the rotor with random data using tntengine Rand function to generate the
	// random data to fill the rotor.
	random.Read(r.Rotor)
	r.Step = step
	r.Start, r.Current = start, start
	r.sliceRotor()
}

// sliceRotor - appends the first 256 bits of the rotor to the end of the rotor.
func (r *Rotor) sliceRotor() {
	var i, j uint
	j = uint(r.Size)
	for i = 0; i < 256; i++ {
		if GetBit(r.Rotor, i) {
			SetBit(r.Rotor, j)
		} else {
			ClrBit(r.Rotor, j)
		}
		j++
	}
}

// SetIndex - set the current rotor position based on the given index
func (r *Rotor) SetIndex(idx *big.Int) {
	// Special case if idx == 0
	if idx.Sign() == 0 {
		r.Current = r.Start
	} else {
		p := new(big.Int)
		q := new(big.Int)
		rem := new(big.Int)
		p = p.Mul(idx, new(big.Int).SetInt64(int64(r.Step)))
		p = p.Add(p, new(big.Int).SetInt64(int64(r.Start)))
		_, rem = q.DivMod(p, new(big.Int).SetInt64(int64(r.Size)), rem)
		r.Current = int(rem.Int64())
	}
}

// Index - Rotor does not track the index.
func (r *Rotor) Index() *big.Int {
	return nil
}

// ApplyF - encrypts the given block of data.
func (r *Rotor) ApplyF(blk *[CipherBlockBytes]byte) *[CipherBlockBytes]byte {
	var res [CipherBlockBytes]byte
	ress := res[:]
	rotor := r.Rotor
	idx := r.Current

	for cnt := 0; cnt < CipherBlockSize; cnt++ {
		if GetBit(rotor, uint(idx)) {
			SetBit(ress, uint(cnt))
		}

		idx++
	}

	r.Current = (r.Current + r.Step) % r.Size
	return AddBlock(blk, &res)
}

// ApplyG - decrypts the given block of data
func (r *Rotor) ApplyG(blk *[CipherBlockBytes]byte) *[CipherBlockBytes]byte {
	var res [CipherBlockBytes]byte
	ress := res[:]
	rotor := r.Rotor[:]
	idx := r.Current

	for cnt := 0; cnt < CipherBlockSize; cnt++ {
		if GetBit(rotor, uint(idx)) {
			SetBit(ress, uint(cnt))
		}

		idx++
	}

	r.Current = (r.Current + r.Step) % r.Size
	return SubBlock(blk, &res)
}

// String - converts a Rotor to a string representation of the Rotor.
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
