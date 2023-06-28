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
			want: Permutator1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := reflect.ValueOf(Permutator1).Interface().(*Permutator)
			got := new(Permutator).New(p.Cycle.Length, p.Randp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermutator() = %v, want %v", got, tt.want)
			}
			if got == tt.want {
				t.Error("NewPermtator():  The new permutator must not equal Permutator1")
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
			p := new(Permutator).New(Permutator1.Cycle.Length, Permutator1.Randp)
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
			p := new(Permutator).New(Permutator1.Cycle.Length, Permutator1.Randp)
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
			p := Permutator1
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
			p := new(Permutator).New(Permutator1.Cycle.Length, Permutator1.Randp)
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
			p := new(Permutator).New(Permutator1.Cycle.Length, Permutator1.Randp)
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
				197, 74, 21, 72, 187, 171, 37, 0, 165, 148, 78, 30, 59, 205, 56, 29,
				251, 135, 34, 90, 210, 200, 3, 216, 1, 123, 185, 188, 94, 105, 248, 20,
				16, 31, 8, 196, 52, 93, 160, 104, 157, 133, 231, 207, 39, 204, 142, 144,
				242, 211, 38, 6, 154, 102, 71, 149, 213, 51, 189, 118, 233, 15, 80, 195,
				26, 240, 124, 139, 81, 100, 32, 40, 55, 190, 217, 147, 249, 112, 145, 44,
				45, 76, 246, 13, 117, 84, 206, 10, 228, 199, 41, 103, 134, 113, 97, 58,
				255, 9, 227, 212, 229, 54, 73, 27, 178, 155, 253, 28, 119, 64, 151, 63,
				225, 35, 220, 60, 203, 237, 159, 221, 172, 239, 23, 89, 2, 146, 14, 219,
				164, 177, 87, 153, 85, 66, 129, 108, 19, 218, 57, 234, 173, 107, 140, 47,
				215, 243, 179, 53, 250, 131, 77, 198, 125, 192, 241, 88, 209, 141, 222, 176,
				7, 91, 170, 48, 150, 158, 202, 43, 122, 86, 161, 70, 167, 75, 168, 191,
				183, 68, 98, 42, 96, 12, 244, 230, 143, 33, 181, 208, 109, 106, 49, 120,
				116, 65, 82, 235, 69, 245, 166, 67, 201, 238, 186, 36, 193, 111, 11, 62,
				126, 162, 254, 163, 83, 127, 223, 132, 252, 79, 110, 137, 61, 5, 180, 232,
				25, 114, 92, 194, 99, 247, 156, 95, 152, 175, 184, 130, 22, 182, 236, 174,
				138, 24, 121, 18, 101, 128, 50, 46, 214, 115, 226, 169, 136, 224, 17, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator).New(Permutator1.Cycle.Length, append([]byte(nil), Permutator1.Randp...))
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
			if got := Permutator1.String(); got != tt.want {
				t.Errorf("Permutator.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
