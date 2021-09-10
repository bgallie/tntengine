// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"bytes"
	"fmt"
	"math/big"
)

// Define constants needed for TNT2
const (
	BitsPerByte             = 8
	CypherBlockSize         = 256 // bits
	CypherBlockBytes        = CypherBlockSize / BitsPerByte
	MaximumRotorSize        = 8192
	NumberPermutationCycles = 4
	RotorSizeBytes          = MaximumRotorSize / BitsPerByte
)

var (
	// RotorSizes is an array of possible rotor sizes.  It consists of prime
	// numbers less than 8160 to allow for 256 bit splce at the end of the rotor.
	// The rotor sizes selected from this list will maximizes the number of
	// unique states the rotors can take.
	RotorSizes = [...]int{
		7523, 7529, 7537, 7541, 7547, 7549, 7559, 7561, 7573, 7577,
		7583, 7589, 7591, 7603, 7607, 7621, 7639, 7643, 7649, 7669,
		7673, 7681, 7687, 7691, 7699, 7703, 7717, 7723, 7727, 7741,
		7753, 7757, 7759, 7789, 7793, 7817, 7823, 7829, 7841, 7853,
		7867, 7873, 7877, 7879, 7883, 7901, 7907, 7919, 7927, 7933}

	// CycleSizes is an array of cycles to use when cycling the permutation table.
	// There are 4 cycles in each entry and they meet the following criteria:
	//      1.  The sum of the cycles is equal to 256.
	//      2.  The cycles are relatively prime to each other. (This maximizes
	//          the number of unique states the permutation can be in for the
	//          given cycles).
	CycleSizes = [...][NumberPermutationCycles]int{
		{61, 63, 65, 67}, // Number of unique states: 16,736,265 [401,670,360]
		{53, 65, 67, 71}, // Number of unique states: 16,387,685 [393,304,440]
		{55, 57, 71, 73}, // Number of unique states: 16,248,705 [389,968,920]
		{53, 61, 63, 79}, // Number of unique states: 16,090,641 [386,175,384]
		{43, 57, 73, 83}, // Number of unique states: 14,850,609 [356,414,616]
		{49, 51, 73, 83}, // Number of unique states: 15,141,441 [363,394,584]
		{47, 53, 73, 83}, // Number of unique states: 15,092,969 [362,231,256]
		{47, 53, 71, 85}} // Number of unique states: 15,033,185 [360,796,440]

	// CyclePermutations is an array of possible orderings that a particular
	// set of four (4) cycle sizes can take.  This is used to increase the number
	// of bitperms that can be generated from the randp table, increasing the
	// complexity that the cryptoanalysis faces.
	CyclePermutations = [...][NumberPermutationCycles]int{
		{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 2, 1}, {0, 3, 1, 2},
		{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 2, 0}, {1, 3, 0, 2},
		{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 1, 0}, {2, 3, 0, 1},
		{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 1, 0}, {3, 2, 0, 1}}

	// Create the proforma rotors and permutator used to create the actual rotors and permutator to use.
	Rotor1 = NewRotor(1783, 863, 1033, []byte{
		184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,
		154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,
		59, 49, 123, 208, 179, 248, 61, 216, 55, 59, 235, 57, 67, 172, 233, 232,
		87, 236, 189, 170, 196, 124, 216, 109, 4, 106, 207, 150, 166, 164, 99, 57,
		131, 27, 1, 236, 168, 78, 122, 81, 165, 26, 32, 56, 129, 105, 35, 26,
		247, 208, 56, 235, 91, 183, 67, 150, 112, 103, 173, 197, 69, 13, 115, 14,
		129, 206, 74, 46, 119, 208, 95, 67, 119, 7, 191, 210, 128, 117, 140, 245,
		41, 168, 63, 203, 53, 241, 221, 28, 158, 40, 89, 76, 126, 58, 33, 40,
		78, 130, 93, 116, 206, 66, 4, 10, 109, 86, 150, 53, 200, 34, 26, 37,
		232, 185, 214, 47, 131, 18, 241, 210, 18, 81, 107, 161, 97, 65, 238, 250,
		81, 133, 54, 158, 54, 10, 254, 135, 110, 162, 175, 250, 117, 66, 232, 66,
		50, 102, 70, 76, 185, 249, 57, 59, 247, 195, 101, 8, 157, 235, 24, 94,
		204, 74, 100, 196, 93, 24, 179, 27, 118, 168, 29, 10, 38, 204, 210, 123,
		111, 247, 225, 171, 60, 166, 239, 124, 43, 180, 223, 240, 66, 2, 68, 220,
		12, 95, 253, 145, 133, 55, 237, 183, 0, 150, 157, 68, 6, 92, 11, 77,
		241, 50, 172, 211, 182, 22, 174, 9, 82, 194, 116, 145, 66, 69, 111})

	Rotor2 = NewRotor(1753, 1494, 1039, []byte{
		100, 120, 105, 253, 78, 6, 70, 91, 136, 33, 73, 16, 15, 13, 174, 206,
		97, 207, 186, 14, 141, 185, 228, 85, 161, 253, 190, 198, 234, 193, 63, 20,
		63, 229, 90, 58, 254, 193, 63, 69, 156, 75, 113, 145, 167, 124, 26, 38,
		94, 117, 42, 25, 81, 251, 172, 67, 175, 138, 159, 85, 66, 180, 187, 101,
		204, 45, 222, 90, 143, 217, 32, 9, 109, 71, 24, 223, 43, 196, 181, 175,
		67, 118, 69, 154, 201, 178, 228, 137, 216, 184, 102, 29, 148, 77, 27, 139,
		90, 20, 115, 102, 91, 37, 244, 44, 9, 254, 144, 216, 214, 201, 70, 160,
		127, 154, 161, 160, 125, 210, 16, 141, 151, 211, 117, 153, 153, 75, 141, 252,
		109, 76, 251, 215, 116, 31, 224, 156, 56, 112, 40, 36, 180, 156, 214, 190,
		122, 206, 11, 172, 52, 68, 167, 87, 53, 234, 125, 167, 21, 100, 193, 166,
		26, 9, 237, 249, 101, 142, 141, 49, 210, 254, 139, 72, 88, 148, 223, 216,
		251, 70, 63, 0, 182, 75, 137, 218, 178, 155, 101, 102, 195, 226, 193, 26,
		9, 12, 147, 186, 248, 43, 5, 117, 133, 78, 14, 201, 165, 155, 206, 57,
		120, 35, 117, 215, 16, 129, 104, 133, 173, 50, 38, 200, 240, 210, 250, 157,
		12, 140, 182, 16, 67, 146, 32, 30, 26, 92, 157, 195, 158, 117, 29, 26,
		115, 201, 171, 66, 251, 125, 141, 213, 131, 127, 40, 102})

	Permutator1 = NewPermutator([]int{43, 57, 73, 83}, []byte{
		207, 252, 142, 205, 239, 35, 230, 62, 69, 94, 166, 89, 184, 81, 144, 120,
		27, 167, 39, 224, 75, 243, 87, 99, 47, 105, 163, 123, 129, 225, 2, 242,
		65, 43, 12, 113, 30, 102, 240, 78, 137, 109, 112, 210, 214, 118, 106, 22,
		232, 181, 164, 255, 70, 198, 160, 44, 231, 20, 228, 53, 85, 238, 178, 133,
		95, 194, 245, 234, 13, 147, 134, 25, 244, 91, 176, 38, 46, 1, 217, 249,
		250, 52, 182, 73, 206, 140, 216, 145, 60, 218, 213, 8, 151, 101, 156, 5,
		241, 67, 49, 42, 212, 180, 92, 21, 16, 130, 128, 126, 98, 199, 162, 188,
		117, 191, 66, 84, 57, 208, 158, 247, 41, 131, 227, 155, 61, 165, 253, 51,
		119, 103, 179, 93, 122, 83, 183, 116, 79, 222, 50, 59, 80, 110, 186, 141,
		90, 152, 127, 107, 54, 71, 185, 161, 169, 34, 148, 146, 157, 138, 24, 237,
		76, 196, 192, 251, 189, 201, 219, 86, 68, 37, 33, 82, 11, 170, 246, 72,
		229, 28, 32, 132, 23, 197, 108, 236, 220, 17, 150, 190, 171, 96, 26, 204,
		209, 31, 211, 4, 14, 136, 195, 45, 172, 111, 154, 36, 149, 226, 202, 187,
		193, 223, 139, 175, 124, 9, 3, 58, 125, 88, 15, 6, 121, 235, 221, 200,
		114, 254, 135, 168, 7, 29, 159, 48, 40, 115, 143, 203, 215, 77, 18, 55,
		56, 177, 100, 0, 173, 104, 248, 97, 74, 63, 233, 19, 64, 174, 153, 10})

	Rotor3 = NewRotor(1721, 1250, 660, []byte{
		25, 134, 2, 219, 108, 110, 170, 11, 12, 129, 29, 172, 198, 2, 14, 255,
		158, 7, 103, 114, 63, 69, 173, 156, 249, 147, 235, 203, 90, 200, 233, 73,
		38, 137, 10, 93, 176, 253, 64, 85, 46, 136, 21, 220, 37, 109, 149, 169,
		165, 153, 37, 42, 63, 35, 65, 196, 237, 215, 100, 226, 151, 53, 172, 215,
		240, 111, 136, 4, 47, 134, 80, 10, 165, 192, 212, 158, 48, 116, 89, 211,
		76, 120, 62, 226, 174, 97, 105, 33, 118, 245, 247, 162, 179, 90, 207, 178,
		69, 114, 201, 206, 93, 130, 79, 199, 223, 120, 233, 66, 86, 178, 59, 104,
		16, 217, 189, 78, 8, 249, 139, 156, 141, 222, 143, 8, 155, 96, 216, 156,
		210, 214, 108, 1, 80, 147, 10, 50, 53, 32, 78, 176, 6, 183, 11, 251,
		130, 192, 204, 184, 131, 159, 142, 127, 170, 183, 238, 60, 6, 47, 77, 30,
		125, 91, 170, 213, 209, 57, 250, 143, 252, 174, 54, 177, 55, 216, 220, 17,
		194, 54, 199, 66, 201, 194, 117, 226, 223, 146, 194, 177, 11, 93, 66, 182,
		46, 122, 253, 161, 204, 40, 167, 40, 92, 37, 134, 155, 0, 231, 21, 105,
		73, 171, 159, 246, 182, 91, 87, 50, 12, 5, 182, 217, 220, 84, 23, 24,
		2, 59, 88, 141, 5, 28, 254, 61, 15, 206, 228, 126, 138, 90, 57, 243,
		39, 215, 151, 181, 144, 211, 147, 198})

	Rotor4 = NewRotor(1741, 1009, 1513, []byte{
		59, 155, 29, 153, 190, 106, 54, 89, 63, 156, 123, 112, 152, 24, 237, 200,
		85, 31, 249, 221, 7, 186, 76, 48, 229, 63, 232, 43, 60, 224, 108, 113,
		71, 154, 254, 136, 83, 102, 6, 108, 108, 138, 65, 104, 190, 98, 197, 120,
		244, 159, 191, 154, 224, 194, 37, 255, 51, 135, 123, 162, 17, 170, 199, 216,
		247, 94, 186, 218, 204, 48, 242, 65, 203, 30, 22, 226, 242, 57, 40, 32,
		22, 231, 138, 222, 125, 10, 125, 108, 24, 59, 221, 99, 156, 96, 214, 129,
		20, 227, 252, 198, 205, 71, 208, 99, 94, 247, 115, 76, 198, 106, 70, 134,
		143, 223, 158, 226, 204, 99, 210, 71, 139, 87, 33, 236, 30, 244, 49, 223,
		228, 215, 142, 236, 68, 74, 166, 97, 216, 67, 14, 41, 128, 40, 55, 70,
		235, 130, 50, 118, 198, 96, 87, 26, 134, 122, 174, 119, 237, 6, 239, 91,
		84, 144, 211, 239, 252, 172, 143, 151, 5, 249, 200, 38, 149, 31, 224, 68,
		100, 250, 25, 173, 38, 74, 133, 18, 244, 7, 138, 0, 85, 143, 137, 140,
		38, 95, 191, 129, 109, 227, 224, 28, 66, 39, 80, 45, 49, 78, 63, 245,
		4, 42, 118, 84, 72, 204, 145, 70, 139, 113, 103, 179, 35, 211, 87, 205,
		38, 235, 135, 115, 15, 14, 19, 163, 29, 185, 234, 35, 191, 251, 64, 151,
		9, 166, 252, 7, 125, 133, 7, 156, 45, 14})

	Permutator2 = NewPermutator([]int{49, 51, 73, 83}, []byte{
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

	Rotor5 = NewRotor(1723, 1293, 1046, []byte{
		59, 137, 3, 62, 80, 176, 170, 169, 12, 135, 154, 73, 218, 169, 34, 130,
		71, 240, 156, 66, 122, 214, 138, 174, 35, 15, 210, 20, 0, 17, 47, 172,
		227, 243, 160, 166, 101, 87, 0, 83, 16, 204, 69, 56, 249, 1, 107, 129,
		30, 236, 248, 46, 59, 16, 136, 240, 7, 68, 175, 181, 102, 24, 221, 34,
		206, 73, 37, 100, 74, 5, 82, 49, 42, 77, 33, 219, 30, 140, 122, 201,
		173, 86, 171, 7, 139, 239, 119, 224, 83, 33, 167, 38, 38, 252, 238, 109,
		173, 151, 153, 182, 170, 199, 109, 174, 85, 177, 165, 37, 171, 94, 247, 29,
		178, 32, 54, 252, 180, 240, 170, 188, 119, 168, 101, 220, 147, 32, 153, 5,
		15, 239, 180, 141, 232, 143, 14, 49, 98, 69, 224, 22, 134, 220, 139, 165,
		26, 189, 188, 120, 113, 196, 95, 124, 238, 91, 217, 213, 114, 32, 177, 200,
		216, 95, 142, 54, 252, 162, 46, 35, 191, 106, 48, 42, 71, 37, 16, 157,
		79, 66, 33, 12, 120, 31, 247, 54, 48, 189, 177, 142, 183, 152, 122, 252,
		139, 150, 164, 251, 77, 9, 128, 220, 145, 27, 85, 162, 42, 154, 151, 87,
		176, 158, 233, 135, 198, 224, 14, 216, 73, 28, 240, 129, 130, 85, 77, 101,
		56, 212, 76, 210, 78, 21, 17, 60, 130, 231, 20, 210, 179, 86, 116, 29,
		121, 144, 166, 0, 136, 120, 97, 5})

	Rotor6 = NewRotor(1733, 1313, 1414, []byte{
		141, 233, 47, 225, 230, 220, 229, 226, 34, 136, 160, 200, 162, 159, 148, 163,
		157, 133, 38, 86, 25, 23, 18, 48, 5, 98, 112, 20, 37, 159, 82, 163,
		209, 135, 40, 197, 152, 8, 255, 234, 149, 22, 158, 19, 235, 186, 173, 247,
		109, 77, 243, 223, 143, 165, 33, 110, 122, 181, 130, 242, 116, 132, 205, 43,
		4, 81, 85, 99, 152, 109, 9, 180, 190, 100, 204, 226, 97, 214, 214, 200,
		169, 61, 53, 107, 128, 231, 15, 42, 162, 156, 119, 166, 223, 143, 234, 16,
		220, 234, 132, 0, 200, 20, 164, 12, 216, 165, 86, 49, 149, 83, 200, 208,
		151, 80, 65, 60, 102, 69, 55, 248, 199, 233, 6, 239, 204, 212, 244, 89,
		255, 240, 54, 232, 189, 143, 233, 51, 44, 167, 97, 2, 71, 233, 154, 155,
		213, 203, 55, 110, 48, 187, 130, 84, 87, 71, 158, 91, 42, 21, 229, 161,
		2, 176, 152, 186, 16, 99, 185, 200, 245, 89, 186, 173, 54, 78, 101, 242,
		169, 224, 83, 242, 78, 39, 93, 123, 86, 196, 13, 82, 104, 92, 139, 230,
		35, 84, 182, 162, 19, 119, 20, 62, 214, 197, 134, 75, 57, 52, 91, 37,
		225, 167, 86, 81, 159, 46, 98, 38, 166, 49, 253, 37, 220, 156, 187, 92,
		92, 4, 17, 20, 89, 244, 147, 114, 180, 179, 208, 196, 42, 227, 66, 2,
		166, 64, 12, 142, 162, 228, 83, 106, 244})

	// BigZero - the big int value for zero.
	BigZero = big.NewInt(0)
	// BigOne - the big int value for one.
	BigOne = big.NewInt(1)
)

// CypherBlock is the data processed by the crypters (rotors and permutators).
// It consistes of the length in bytes to process and the (32 bytes of) data to
// process.
type CypherBlock struct {
	Length      int8
	CypherBlock [CypherBlockBytes]byte
}

// Marshall converts a CypherBlock into a slice of bytes
func (cblk *CypherBlock) Marshall() []byte {
	b := make([]byte, 0, 0)
	b = append(b, byte(cblk.Length))
	b = append(b, cblk.CypherBlock[:]...)
	return b
}

// Unmarshall converts a slice of bytes (created by Marshall) into a CypherBlock
func (cblk *CypherBlock) Unmarshall(b []byte) *CypherBlock {
	blk := new(CypherBlock)
	blk.Length = int8(b[0])
	_ = copy(blk.CypherBlock[:], b[1:])
	return blk
}

// String formats a string representing the permutator (as Go source code).
func (cblk *CypherBlock) String() string {
	var output bytes.Buffer
	output.WriteString(fmt.Sprint("CypherBlock: "))
	output.WriteString(fmt.Sprintf("\t     Length: %d\n", cblk.Length))
	output.WriteString(fmt.Sprintf("\tCypherBlock:\t% X\n", cblk.CypherBlock[0:16]))
	output.WriteString(fmt.Sprintf("\t\t\t% X", cblk.CypherBlock[16:]))
	return output.String()
}

// Crypter interface
type Crypter interface {
	SetIndex(*big.Int)
	Index() *big.Int
	ApplyF(*[CypherBlockBytes]byte) *[CypherBlockBytes]byte
	ApplyG(*[CypherBlockBytes]byte) *[CypherBlockBytes]byte
}

// Counter is a cryptor that does not encrypt/decrypt any data but counts the
// number of blocks that were encrypted.
type Counter struct {
	index *big.Int
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
func (cntr *Counter) ApplyF(blk *[CypherBlockBytes]byte) *[CypherBlockBytes]byte {
	cntr.index.Add(cntr.index, BigOne)
	return blk
}

// ApplyG - this function does nothing during decryption.
func (cntr *Counter) ApplyG(blk *[CypherBlockBytes]byte) *[CypherBlockBytes]byte {
	return blk
}

// SubBlock - subtracts (not XOR) the key from the data to be decrypted
func SubBlock(blk, key *[CypherBlockBytes]byte) *[CypherBlockBytes]byte {
	var p int

	for idx, val := range *blk {
		p = p + int(val) - int(key[idx])
		blk[idx] = byte(p & 0xFF)
		p = p >> BitsPerByte
	}

	return blk
}

// AddBlock - adds (not XOR) the data to be encrypted with the key.
func AddBlock(blk, key *[CypherBlockBytes]byte) *[CypherBlockBytes]byte {
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
func EncryptMachine(ecm Crypter, left chan CypherBlock) chan CypherBlock {
	if ecm == nil {
		panic("ecm is nil")
	}
	right := make(chan CypherBlock)
	go func(ecm Crypter, left chan CypherBlock, right chan CypherBlock) {
		defer close(right)
		for {
			inp := <-left
			if inp.Length <= 0 {
				right <- inp
				break
			}
			inp.CypherBlock = *ecm.ApplyF(&inp.CypherBlock)
			right <- inp
		}
	}(ecm, left, right)

	return right
}

// DecryptMachine - set up a rotor, permutator, or counter to decrypt a block
// read from the left (input channel) and send it out on the right (output channel)
func DecryptMachine(ecm Crypter, left chan CypherBlock) chan CypherBlock {
	right := make(chan CypherBlock)
	go func(ecm Crypter, left chan CypherBlock, right chan CypherBlock) {
		defer close(right)
		for {
			inp := <-left
			if inp.Length <= 0 {
				right <- inp
				break
			}

			inp.CypherBlock = *ecm.ApplyG(&inp.CypherBlock)
			right <- inp
		}
	}(ecm, left, right)

	return right
}

// CreateEncryptMachine -
func createEncryptMachine(ecms ...Crypter) (left chan CypherBlock, right chan CypherBlock) {
	if ecms != nil {
		idx := 0
		left = make(chan CypherBlock)
		right = EncryptMachine(ecms[idx], left)

		for idx++; idx < len(ecms); idx++ {
			right = EncryptMachine(ecms[idx], right)
		}
	} else {
		panic("you must give at least one encryption device!")
	}

	return
}

// CreateDecryptMachine -
func createDecryptMachine(ecms ...Crypter) (left chan CypherBlock, right chan CypherBlock) {
	if ecms != nil {
		idx := len(ecms) - 1
		left = make(chan CypherBlock)
		right = DecryptMachine(ecms[idx], left)

		for idx--; idx >= 0; idx-- {
			right = DecryptMachine(ecms[idx], right)
		}
	} else {
		panic("you must give at least one decryption device!")
	}

	return
}
