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
				cycleSizes: []int{43, 57, 73, 83},
				randp: []byte{
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
					56, 177, 100, 0, 173, 104, 248, 97, 74, 63, 233, 19, 64, 174, 153, 10},
			},
			want: Permutator1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPermutator(tt.args.cycleSizes, tt.args.randp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermutator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_Update(t *testing.T) {
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
				cycleSizes: []int{43, 57, 73, 83},
				randp: []byte{
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
					56, 177, 100, 0, 173, 104, 248, 97, 74, 63, 233, 19, 64, 174, 153, 10},
			},
			want: Permutator1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(Permutator)
			p.Update(tt.args.cycleSizes, tt.args.randp)
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
		bitPerm       [CypherBlockSize]byte
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
		bitPerm       [CypherBlockSize]byte
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
		bitPerm       [CypherBlockSize]byte
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
		blk *[CypherBlockBytes]byte
	}
	tests := []struct {
		name string
		args args
		want *[CypherBlockBytes]byte
	}{
		{
			name: "tpaf1",
			args: args{
				&[CypherBlockBytes]byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
					17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				},
			},
			want: &[CypherBlockBytes]byte{
				209, 217, 128, 24, 115, 4, 114, 33, 6, 18, 17, 204, 16, 160, 173, 86,
				133, 128, 48, 33, 152, 233, 34, 224, 3, 136, 162, 192, 32, 5, 2, 4,
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
		bitPerm       [CypherBlockSize]byte
	}
	type args struct {
		blk *[CypherBlockBytes]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[CypherBlockBytes]byte
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
			want: "permutator.New([]int{43, 57, 73, 83}, []byte{\n" +
				"	207, 252, 142, 205, 239, 35, 230, 62, 69, 94, 166, 89, 184, 81, 144, 120, \n" +
				"	27, 167, 39, 224, 75, 243, 87, 99, 47, 105, 163, 123, 129, 225, 2, 242, \n" +
				"	65, 43, 12, 113, 30, 102, 240, 78, 137, 109, 112, 210, 214, 118, 106, 22, \n" +
				"	232, 181, 164, 255, 70, 198, 160, 44, 231, 20, 228, 53, 85, 238, 178, 133, \n" +
				"	95, 194, 245, 234, 13, 147, 134, 25, 244, 91, 176, 38, 46, 1, 217, 249, \n" +
				"	250, 52, 182, 73, 206, 140, 216, 145, 60, 218, 213, 8, 151, 101, 156, 5, \n" +
				"	241, 67, 49, 42, 212, 180, 92, 21, 16, 130, 128, 126, 98, 199, 162, 188, \n" +
				"	117, 191, 66, 84, 57, 208, 158, 247, 41, 131, 227, 155, 61, 165, 253, 51, \n" +
				"	119, 103, 179, 93, 122, 83, 183, 116, 79, 222, 50, 59, 80, 110, 186, 141, \n" +
				"	90, 152, 127, 107, 54, 71, 185, 161, 169, 34, 148, 146, 157, 138, 24, 237, \n" +
				"	76, 196, 192, 251, 189, 201, 219, 86, 68, 37, 33, 82, 11, 170, 246, 72, \n" +
				"	229, 28, 32, 132, 23, 197, 108, 236, 220, 17, 150, 190, 171, 96, 26, 204, \n" +
				"	209, 31, 211, 4, 14, 136, 195, 45, 172, 111, 154, 36, 149, 226, 202, 187, \n" +
				"	193, 223, 139, 175, 124, 9, 3, 58, 125, 88, 15, 6, 121, 235, 221, 200, \n" +
				"	114, 254, 135, 168, 7, 29, 159, 48, 40, 115, 143, 203, 215, 77, 18, 55, \n" +
				"	56, 177, 100, 0, 173, 104, 248, 97, 74, 63, 233, 19, 64, 174, 153, 10})\n",
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
