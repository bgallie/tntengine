// Package tntEngine - define TntEngine type and it's methods
package tntEngine

import (
	"math/bits"
)

// var blk CypherBlock

type Rand struct {
	tntMachine *TntEngine
	idx        int
	blk        CypherBlock
}

func NewRand(src *TntEngine) *Rand {
	var rand Rand
	rand.tntMachine = src
	rand.idx = CypherBlockBytes
	return &rand
}

func (rnd *Rand) Intn(max int) int {
	if max <= 0 {
		panic("argument to intn is <= 0")
	}

	n := max - 1
	// bitLen is the maximum bit length needed to encode a value < max.
	bitLen := bits.Len(uint(n))
	if bitLen == 0 {
		// the only valid result is 0
		return n
	}
	// k is the maximum byte length needed to encode a value < max.
	k := (bitLen + 7) / 8
	// b is the number of bits in the most significant byte of max-1.
	b := uint(bitLen % 8)
	if b == 0 {
		b = 8
	}

	bytes := make([]byte, k)

	for {
		_, _ = rnd.Read(bytes)
		// Clear bits in the first byte to increase the probability
		// that the candidate is < max.
		bytes[0] &= uint8(int(1<<b) - 1)

		// Change the data in the byte slice into an integer ('n')
		n = 0
		for _, val := range bytes {
			n = (n << 8) | int(val)
		}

		if n < max {
			return n
		}
	}
}

func (rnd *Rand) Int63n(n int64) int64 {
	return int64(rnd.Intn(int(n)))
}

func (rnd *Rand) Perm(n int) []int {
	res := make([]int, n, n)

	for i := range res {
		res[i] = i
	}

	for i := (n - 1); i > 0; i-- {
		j := rnd.Intn(i)
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func (rnd *Rand) Read(p []byte) (n int, err error) {
	err = nil
	p = p[:0]
	left := rnd.tntMachine.Left()
	right := rnd.tntMachine.Right()
	for {
		if rnd.idx >= CypherBlockBytes {
			if rnd.blk.Length == 0 {
				rnd.blk.Length = CypherBlockBytes
				_ = copy(rnd.blk.CypherBlock[:], rnd.tntMachine.jc1Key.XORKeyStream(rnd.blk.CypherBlock[:]))
			}
			left <- rnd.blk
			rnd.blk = <-right
			rnd.idx = 0
		}
		leftInBlk := len(rnd.blk.CypherBlock) - rnd.idx
		remaining := cap(p) - len(p)
		if remaining >= leftInBlk {
			p = append(p, rnd.blk.CypherBlock[rnd.idx:]...)
			rnd.idx += leftInBlk
		} else {
			p = append(p, rnd.blk.CypherBlock[0:remaining]...)
			rnd.idx += remaining
			break
		}
	}

	return len(p), err
}
