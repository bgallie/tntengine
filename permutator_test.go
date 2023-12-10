// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewPermutator(t *testing.T) {
	tests := []struct {
		name string
		want *Permutator
	}{
		{
			name: "tnp1",
			want: proFormPermutators[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := reflect.ValueOf(proFormPermutators[0]).Interface().(*Permutator)
			got := new(Permutator).New(p.Cycle.Length, p.Randp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermutator() = %v, want %v", got, tt.want)
			}
			if got == tt.want {
				t.Error("NewPermtator():  The new permutator must not equal proFormPermutators[0]")
			}
		})
	}
}

func TestPermutator_nextState(t *testing.T) {
	// This tests both Permutator.nextState() and Permutator.cycle()
	tests := []struct {
		name string
		want [256]byte
	}{
		{
			name: "tpns1",
			want: [256]byte{
				93, 213, 153, 127, 57, 218, 14, 175, 189, 53, 171, 18, 32, 238, 68, 61,
				162, 98, 35, 40, 164, 151, 120, 241, 216, 197, 133, 115, 130, 111, 152, 79,
				191, 73, 95, 100, 134, 49, 248, 194, 102, 20, 65, 173, 147, 177, 27, 250,
				116, 150, 196, 21, 215, 125, 230, 231, 12, 51, 163, 247, 155, 253, 56, 137,
				144, 41, 190, 52, 178, 254, 88, 119, 7, 1, 195, 45, 124, 139, 210, 103,
				174, 33, 199, 37, 159, 170, 17, 24, 193, 90, 255, 143, 172, 176, 76, 180,
				121, 212, 84, 187, 69, 244, 54, 131, 89, 44, 233, 48, 25, 23, 242, 83,
				117, 136, 166, 200, 86, 145, 201, 39, 38, 16, 179, 104, 78, 184, 94, 5,
				219, 80, 161, 82, 243, 206, 220, 106, 158, 224, 105, 129, 64, 245, 101, 217,
				77, 2, 169, 26, 142, 50, 223, 70, 183, 234, 87, 47, 168, 71, 225, 236,
				114, 4, 63, 227, 42, 149, 222, 246, 249, 59, 126, 165, 240, 91, 185, 198,
				31, 81, 97, 232, 112, 186, 239, 36, 204, 108, 13, 11, 3, 75, 128, 66,
				167, 181, 208, 138, 9, 192, 146, 85, 207, 156, 122, 46, 29, 205, 229, 60,
				188, 28, 203, 0, 252, 141, 8, 43, 157, 10, 110, 251, 202, 92, 58, 72,
				15, 211, 107, 6, 55, 99, 235, 182, 140, 113, 19, 22, 209, 214, 237, 132,
				34, 148, 74, 160, 67, 118, 154, 109, 96, 123, 30, 135, 226, 221, 62, 228},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator).New(proFormPermutators[0].Cycle.Length, proFormPermutators[0].Randp)
			p.nextState() // the first call does not cycles
			if p.bitPerm != tt.want {
				t.Errorf("p.bitPerm = %v, want %v", p.bitPerm, tt.want)
			}
		})
	}
}

func TestPermutator_SetIndex(t *testing.T) {
	type args struct {
		idx *big.Int
	}
	tests := []struct {
		name string
		args args
		want [256]byte
	}{
		{
			name: "tpsi1",
			args: args{
				idx: BigOne,
			},
			want: [256]byte{
				93, 213, 153, 127, 57, 218, 14, 175, 189, 53, 171, 18, 32, 238, 68, 61,
				162, 98, 35, 40, 164, 151, 120, 241, 216, 197, 133, 115, 130, 111, 152, 79,
				191, 73, 95, 100, 134, 49, 248, 194, 102, 20, 65, 173, 147, 177, 27, 250,
				116, 150, 196, 21, 215, 125, 230, 231, 12, 51, 163, 247, 155, 253, 56, 137,
				144, 41, 190, 52, 178, 254, 88, 119, 7, 1, 195, 45, 124, 139, 210, 103,
				174, 33, 199, 37, 159, 170, 17, 24, 193, 90, 255, 143, 172, 176, 76, 180,
				121, 212, 84, 187, 69, 244, 54, 131, 89, 44, 233, 48, 25, 23, 242, 83,
				117, 136, 166, 200, 86, 145, 201, 39, 38, 16, 179, 104, 78, 184, 94, 5,
				219, 80, 161, 82, 243, 206, 220, 106, 158, 224, 105, 129, 64, 245, 101, 217,
				77, 2, 169, 26, 142, 50, 223, 70, 183, 234, 87, 47, 168, 71, 225, 236,
				114, 4, 63, 227, 42, 149, 222, 246, 249, 59, 126, 165, 240, 91, 185, 198,
				31, 81, 97, 232, 112, 186, 239, 36, 204, 108, 13, 11, 3, 75, 128, 66,
				167, 181, 208, 138, 9, 192, 146, 85, 207, 156, 122, 46, 29, 205, 229, 60,
				188, 28, 203, 0, 252, 141, 8, 43, 157, 10, 110, 251, 202, 92, 58, 72,
				15, 211, 107, 6, 55, 99, 235, 182, 140, 113, 19, 22, 209, 214, 237, 132,
				34, 148, 74, 160, 67, 118, 154, 109, 96, 123, 30, 135, 226, 221, 62, 228},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator).New(proFormPermutators[0].Cycle.Length, proFormPermutators[0].Randp)
			p.SetIndex(tt.args.idx)
			if p.bitPerm != tt.want {
				t.Errorf("p.bitPerm = %v, want %v", p.bitPerm, tt.want)
			}
		})
	}
}

func TestPermutator_Index(t *testing.T) {
	tests := []struct {
		name string
		want *big.Int
	}{
		{
			name: "tpi1",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := proFormPermutators[0]
			if got := p.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_ApplyF(t *testing.T) {
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tpaf1",
			args: args{
				[]byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
					17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				},
			},
			want: []byte{
				16, 66, 66, 102, 144, 89, 68, 25, 50, 40, 147, 34, 232, 163,
				1, 16, 69, 35, 144, 64, 2, 2, 16, 175, 98, 54, 32, 113, 10, 44, 5, 35,
			},
		},
		{ // A CipherBlock with less than 32 bytes will not have the permutation applied to it.
			name: "tpaf2",
			args: args{
				[]byte{
					1, 2, 3, 4, 5, 6,
				},
			},
			want: []byte{
				1, 2, 3, 4, 5, 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator).New(proFormPermutators[0].Cycle.Length, proFormPermutators[0].Randp)
			if got := p.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_ApplyG(t *testing.T) {
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tpafg1",
			args: args{
				[]byte{
					16, 66, 66, 102, 144, 89, 68, 25, 50, 40, 147, 34, 232, 163,
					1, 16, 69, 35, 144, 64, 2, 2, 16, 175, 98, 54, 32, 113, 10, 44, 5, 35,
				},
			},
			want: []byte{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
				17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
			},
		},
		{ // A CipherBlock with less than 32 bytes will not have the permutation applied to it.
			name: "tpafg2",
			args: args{
				[]byte{
					16, 66, 66, 102,
				},
			},
			want: []byte{
				16, 66, 66, 102,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator).New(proFormPermutators[0].Cycle.Length, proFormPermutators[0].Randp)
			if got := p.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.ApplyG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_Update(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	rnd := new(Rand).New(tntMachine)
	defer tntMachine.CloseCipherMachine()
	tests := []struct {
		name string
		want *Permutator
	}{
		{
			name: "tpu1",
			want: new(Permutator).New(256, []byte{
				204, 239, 182, 117, 24, 144, 159, 57, 84, 116, 228, 56, 166, 210, 8, 231,
				241, 169, 174, 253, 17, 31, 85, 139, 180, 62, 42, 198, 161, 150, 10, 5,
				54, 187, 238, 66, 29, 113, 168, 194, 63, 244, 203, 68, 15, 152, 215, 142,
				75, 58, 13, 243, 255, 193, 45, 30, 3, 153, 230, 112, 134, 226, 129, 120,
				122, 190, 130, 250, 160, 121, 55, 88, 32, 4, 207, 206, 100, 99, 12, 20,
				233, 104, 185, 28, 9, 44, 235, 7, 123, 37, 19, 46, 43, 170, 76, 78,
				51, 6, 212, 147, 119, 40, 164, 136, 50, 135, 96, 200, 114, 177, 52, 145,
				110, 108, 158, 71, 218, 115, 176, 47, 59, 222, 237, 191, 102, 79, 89, 53,
				25, 173, 246, 254, 124, 128, 216, 162, 109, 82, 225, 95, 156, 236, 186, 181,
				209, 247, 221, 219, 65, 107, 245, 199, 83, 127, 103, 131, 18, 252, 133, 251,
				214, 155, 16, 125, 195, 183, 34, 143, 138, 22, 192, 132, 217, 202, 73, 23,
				220, 208, 140, 91, 80, 157, 118, 229, 172, 223, 60, 87, 141, 106, 39, 189,
				227, 197, 175, 154, 11, 148, 111, 224, 86, 77, 41, 205, 163, 249, 72, 81,
				184, 27, 2, 93, 64, 149, 74, 126, 188, 49, 248, 165, 90, 1, 137, 171,
				36, 26, 101, 178, 94, 38, 61, 67, 98, 151, 242, 70, 69, 146, 232, 196,
				48, 179, 14, 92, 234, 0, 240, 97, 213, 33, 35, 105, 21, 211, 167, 201}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator).New(proFormPermutators[0].Cycle.Length, append([]byte(nil), proFormPermutators[0].Randp...))
			p.Update(rnd)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Updated Permutator() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPermutator_String(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "tps1",
			want: "new(Permutator).New(256, []byte{\n" +
				"	248, 250, 32, 91, 122, 166, 115, 61, 178, 111, 37, 35, 82, 167, 157, 66,\n" +
				"	22, 65, 47, 1, 195, 182, 190, 73, 19, 218, 237, 76, 140, 155, 18, 11,\n" +
				"	30, 207, 105, 49, 230, 83, 10, 251, 52, 136, 99, 212, 108, 154, 113, 41,\n" +
				"	185, 44, 102, 226, 135, 165, 94, 27, 6, 177, 162, 161, 209, 200, 33, 23,\n" +
				"	197, 120, 71, 249, 125, 244, 217, 38, 0, 128, 95, 80, 214, 254, 163, 203,\n" +
				"	180, 137, 100, 235, 16, 58, 78, 173, 3, 118, 148, 191, 15, 7, 149, 219,\n" +
				"	39, 129, 75, 158, 224, 92, 147, 144, 236, 60, 29, 9, 252, 51, 139, 97,\n" +
				"	43, 87, 193, 222, 85, 223, 127, 153, 192, 13, 143, 70, 151, 123, 211, 72,\n" +
				"	93, 194, 229, 42, 17, 146, 196, 107, 215, 112, 231, 21, 124, 86, 132, 238,\n" +
				"	26, 189, 98, 172, 201, 175, 188, 88, 114, 5, 25, 64, 103, 246, 45, 57,\n" +
				"	109, 63, 81, 62, 204, 106, 179, 199, 116, 141, 186, 121, 84, 210, 79, 156,\n" +
				"	216, 14, 253, 233, 46, 55, 138, 34, 74, 20, 245, 89, 198, 133, 239, 142,\n" +
				"	234, 24, 176, 213, 169, 241, 90, 232, 28, 240, 183, 227, 56, 247, 160, 152,\n" +
				"	202, 4, 159, 104, 187, 31, 174, 48, 168, 67, 40, 50, 134, 228, 181, 170,\n" +
				"	225, 126, 54, 36, 220, 208, 150, 117, 255, 221, 101, 69, 77, 110, 243, 206,\n" +
				"	130, 59, 205, 242, 184, 164, 131, 12, 2, 119, 96, 171, 53, 68, 8, 145})\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := proFormPermutators[0].String(); got != tt.want {
				t.Errorf("Permutator.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
