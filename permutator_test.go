// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewPermutator(t *testing.T) {
	type args struct {
		cycleSizes []int
		randp      []byte
	}
	tests := []struct {
		name string
		args args
		want *Permutator
	}{
		{
			name: "tnp1",
			args: args{
				cycleSizes: []int{256},
				randp: []byte{
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
					130, 59, 205, 242, 184, 164, 131, 12, 2, 119, 96, 171, 53, 68, 8, 145},
			},
			want: Permutator1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new(Permutator).New(tt.args.cycleSizes, tt.args.randp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermutator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_Update(t *testing.T) {
	type args struct {
		random *Rand
	}
	tests := []struct {
		name string
		args args
		want *Permutator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator)
			p.Update(tt.args.random)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Updated Permutator() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPermutator_nextState(t *testing.T) {
	type fields struct {
		CurrentState  int
		MaximalStates int
		Cycles        []Cycle
		Randp         []byte
		bitPerm       [CipherBlockSize]byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Permutator{
				CurrentState:  tt.fields.CurrentState,
				MaximalStates: tt.fields.MaximalStates,
				Cycles:        tt.fields.Cycles,
				Randp:         tt.fields.Randp,
				bitPerm:       tt.fields.bitPerm,
			}
			p.nextState()
		})
	}
}

func TestPermutator_cycle(t *testing.T) {
	type fields struct {
		CurrentState  int
		MaximalStates int
		Cycles        []Cycle
		Randp         []byte
		bitPerm       [CipherBlockSize]byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Permutator{
				CurrentState:  tt.fields.CurrentState,
				MaximalStates: tt.fields.MaximalStates,
				Cycles:        tt.fields.Cycles,
				Randp:         tt.fields.Randp,
				bitPerm:       tt.fields.bitPerm,
			}
			p.cycle()
		})
	}
}

func TestPermutator_SetIndex(t *testing.T) {
	type fields struct {
		CurrentState  int
		MaximalStates int
		Cycles        []Cycle
		Randp         []byte
		bitPerm       [CipherBlockSize]byte
	}
	type args struct {
		idx *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Permutator{
				CurrentState:  tt.fields.CurrentState,
				MaximalStates: tt.fields.MaximalStates,
				Cycles:        tt.fields.Cycles,
				Randp:         tt.fields.Randp,
				bitPerm:       tt.fields.bitPerm,
			}
			p.SetIndex(tt.args.idx)
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
		blk *[CipherBlockBytes]byte
	}
	tests := []struct {
		name string
		args args
		want *[CipherBlockBytes]byte
	}{
		{
			name: "tpaf1",
			args: args{
				&[CipherBlockBytes]byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
					17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				},
			},
			want: &[CipherBlockBytes]byte{
				16, 66, 66, 102, 144, 89, 68, 25, 50, 40, 147, 34, 232, 163,
				1, 16, 69, 35, 144, 64, 2, 2, 16, 175, 98, 54, 32, 113, 10, 44, 5, 35,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Permutator1
			if got := p.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_ApplyG(t *testing.T) {
	type fields struct {
		CurrentState  int
		MaximalStates int
		Cycles        []Cycle
		Randp         []byte
		bitPerm       [CipherBlockSize]byte
	}
	type args struct {
		blk *[CipherBlockBytes]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[CipherBlockBytes]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Permutator{
				CurrentState:  tt.fields.CurrentState,
				MaximalStates: tt.fields.MaximalStates,
				Cycles:        tt.fields.Cycles,
				Randp:         tt.fields.Randp,
				bitPerm:       tt.fields.bitPerm,
			}
			if got := p.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.ApplyG() = %v, want %v", got, tt.want)
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
			want: "new(Permutator).New([]int{256}, []byte{\n" +
				"	248, 250, 32, 91, 122, 166, 115, 61, 178, 111, 37, 35, 82, 167, 157, 66, \n" +
				"	22, 65, 47, 1, 195, 182, 190, 73, 19, 218, 237, 76, 140, 155, 18, 11, \n" +
				"	30, 207, 105, 49, 230, 83, 10, 251, 52, 136, 99, 212, 108, 154, 113, 41, \n" +
				"	185, 44, 102, 226, 135, 165, 94, 27, 6, 177, 162, 161, 209, 200, 33, 23, \n" +
				"	197, 120, 71, 249, 125, 244, 217, 38, 0, 128, 95, 80, 214, 254, 163, 203, \n" +
				"	180, 137, 100, 235, 16, 58, 78, 173, 3, 118, 148, 191, 15, 7, 149, 219, \n" +
				"	39, 129, 75, 158, 224, 92, 147, 144, 236, 60, 29, 9, 252, 51, 139, 97, \n" +
				"	43, 87, 193, 222, 85, 223, 127, 153, 192, 13, 143, 70, 151, 123, 211, 72, \n" +
				"	93, 194, 229, 42, 17, 146, 196, 107, 215, 112, 231, 21, 124, 86, 132, 238, \n" +
				"	26, 189, 98, 172, 201, 175, 188, 88, 114, 5, 25, 64, 103, 246, 45, 57, \n" +
				"	109, 63, 81, 62, 204, 106, 179, 199, 116, 141, 186, 121, 84, 210, 79, 156, \n" +
				"	216, 14, 253, 233, 46, 55, 138, 34, 74, 20, 245, 89, 198, 133, 239, 142, \n" +
				"	234, 24, 176, 213, 169, 241, 90, 232, 28, 240, 183, 227, 56, 247, 160, 152, \n" +
				"	202, 4, 159, 104, 187, 31, 174, 48, 168, 67, 40, 50, 134, 228, 181, 170, \n" +
				"	225, 126, 54, 36, 220, 208, 150, 117, 255, 221, 101, 69, 77, 110, 243, 206, \n" +
				"	130, 59, 205, 242, 184, 164, 131, 12, 2, 119, 96, 171, 53, 68, 8, 145})\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Permutator1
			if got := p.String(); got != tt.want {
				t.Errorf("Permutator.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
