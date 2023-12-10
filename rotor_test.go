// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewRotor(t *testing.T) {
	type args struct {
		size  int
		start int
		step  int
		rotor []byte
	}
	tests := []struct {
		name string
		args args
		want *Rotor
	}{
		{
			name: "tnr",
			args: args{
				size:  1789,
				start: 1065,
				step:  1499,
				rotor: []byte{
					63, 180, 255, 162, 59, 142, 61, 13, 187, 226, 49, 134, 163, 38, 44, 14,
					255, 73, 155, 237, 208, 42, 217, 227, 194, 245, 229, 169, 96, 163, 33, 145,
					222, 156, 57, 87, 220, 186, 118, 131, 89, 103, 27, 145, 153, 207, 16, 55,
					248, 183, 83, 65, 15, 253, 147, 136, 217, 189, 124, 150, 193, 113, 87, 127,
					101, 202, 87, 3, 80, 160, 132, 129, 1, 134, 154, 36, 194, 3, 186, 148,
					241, 226, 134, 255, 59, 78, 202, 236, 166, 151, 184, 209, 115, 21, 177, 17,
					106, 189, 209, 128, 13, 224, 94, 163, 47, 117, 151, 3, 9, 88, 20, 74,
					188, 243, 174, 130, 193, 247, 161, 74, 119, 95, 40, 111, 215, 174, 84, 170,
					234, 27, 241, 147, 210, 26, 139, 92, 231, 118, 227, 206, 0, 186, 161, 82,
					149, 59, 93, 134, 84, 108, 116, 191, 127, 153, 92, 59, 80, 53, 10, 112,
					127, 228, 183, 134, 214, 74, 150, 134, 145, 60, 22, 217, 213, 195, 251, 240,
					232, 1, 193, 235, 142, 191, 153, 123, 46, 86, 198, 123, 33, 34, 148, 104,
					18, 96, 34, 17, 139, 199, 225, 84, 245, 102, 137, 167, 240, 84, 152, 144,
					171, 21, 67, 253, 113, 97, 156, 145, 55, 87, 247, 45, 54, 48, 157, 247,
					135, 246, 95, 116, 199, 177, 167, 97, 87, 60, 198, 112, 212, 132, 197, 225,
					63, 105, 179, 29, 90, 37, 123, 92, 184, 190, 60, 21, 108, 52, 36, 18},
			},
			want: proFormaRotors[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new(Rotor).New(tt.args.size, tt.args.start, tt.args.step, tt.args.rotor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRotor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_sliceRotor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "trsr1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(Rotor).New(1789, 1065, 1499, []byte{
				63, 180, 255, 162, 59, 142, 61, 13, 187, 226, 49, 134, 163, 38, 44, 14,
				255, 73, 155, 237, 208, 42, 217, 227, 194, 245, 229, 169, 96, 163, 33, 145,
				222, 156, 57, 87, 220, 186, 118, 131, 89, 103, 27, 145, 153, 207, 16, 55,
				248, 183, 83, 65, 15, 253, 147, 136, 217, 189, 124, 150, 193, 113, 87, 127,
				101, 202, 87, 3, 80, 160, 132, 129, 1, 134, 154, 36, 194, 3, 186, 148,
				241, 226, 134, 255, 59, 78, 202, 236, 166, 151, 184, 209, 115, 21, 177, 17,
				106, 189, 209, 128, 13, 224, 94, 163, 47, 117, 151, 3, 9, 88, 20, 74,
				188, 243, 174, 130, 193, 247, 161, 74, 119, 95, 40, 111, 215, 174, 84, 170,
				234, 27, 241, 147, 210, 26, 139, 92, 231, 118, 227, 206, 0, 186, 161, 82,
				149, 59, 93, 134, 84, 108, 116, 191, 127, 153, 92, 59, 80, 53, 10, 112,
				127, 228, 183, 134, 214, 74, 150, 134, 145, 60, 22, 217, 213, 195, 251, 240,
				232, 1, 193, 235, 142, 191, 153, 123, 46, 86, 198, 123, 33, 34, 148, 104,
				18, 96, 34, 17, 139, 199, 225, 84, 245, 102, 137, 167, 240, 84, 152, 144,
				171, 21, 67, 253, 113, 97, 156, 145, 55, 87, 247, 45, 54, 48, 157, 247,
				135, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
			r.sliceRotor()
			if !reflect.DeepEqual(r, proFormaRotors[0]) {
				t.Errorf("Sliced Rotor() = %v, want %v", r, proFormaRotors[0])
			}
		})
	}
}

func TestRotor_Update(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	rnd := new(Rand).New(tntMachine)
	defer tntMachine.CloseCipherMachine()
	tests := []struct {
		name string
		want *Rotor
	}{
		{
			name: "tur1",
			want: new(Rotor).New(1789, 423, 790, []byte{
				101, 242, 26, 36, 171, 137, 1, 90, 242, 165, 237, 246, 213, 49, 188, 126,
				74, 149, 64, 93, 2, 27, 219, 184, 81, 72, 33, 163, 33, 41, 77, 86,
				36, 111, 248, 236, 148, 11, 154, 215, 229, 175, 238, 148, 231, 178, 189, 39,
				106, 199, 242, 141, 87, 60, 171, 172, 225, 38, 184, 224, 225, 118, 157, 80,
				91, 140, 172, 90, 23, 215, 73, 41, 49, 132, 255, 250, 140, 22, 138, 252,
				143, 34, 38, 244, 249, 154, 125, 16, 180, 155, 74, 238, 173, 105, 133, 229,
				21, 18, 131, 103, 127, 83, 36, 0, 107, 65, 178, 38, 1, 97, 27, 193,
				21, 60, 207, 211, 249, 204, 69, 165, 242, 218, 170, 179, 198, 18, 95, 26,
				82, 109, 233, 212, 252, 16, 69, 128, 226, 216, 124, 243, 21, 178, 166, 182,
				48, 201, 245, 41, 153, 53, 89, 207, 230, 234, 129, 82, 59, 175, 90, 252,
				115, 107, 199, 114, 25, 236, 249, 110, 97, 180, 177, 25, 255, 86, 224, 234,
				248, 16, 50, 177, 49, 40, 248, 47, 166, 238, 64, 6, 51, 125, 206, 76,
				18, 97, 248, 171, 46, 109, 19, 224, 115, 165, 179, 7, 198, 172, 137, 28,
				49, 50, 87, 36, 81, 20, 229, 89, 213, 118, 220, 12, 98, 119, 87, 183,
				76, 94, 131, 100, 53, 49, 64, 75, 190, 180, 221, 190, 58, 134, 215, 79,
				169, 18, 168, 75, 96, 99, 27, 55, 10, 41, 100, 52, 36, 165, 201, 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(Rotor).New(proFormaRotors[0].Size, proFormaRotors[0].Start, proFormaRotors[0].Size, append([]byte(nil), proFormaRotors[0].Rotor...))
			r.Update(rnd)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Updated Rotor() = %v, want %v", r, tt.want)
			}
		})
	}
}

func TestRotor_SetIndex(t *testing.T) {
	type args struct {
		idx *big.Int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trsi1",
			args: args{
				idx: big.NewInt(10000),
			},
			want: 1034,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := proFormaRotors[0]
			r.SetIndex(tt.args.idx)
			if r.Current != tt.want {
				t.Errorf("r.Current = %v, want %v", r.Current, tt.want)
			}
		})
	}
}

func TestRotor_Index(t *testing.T) {
	tests := []struct {
		name string
		want *big.Int
	}{
		{
			name: "trsi1",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := proFormaRotors[0]
			if got := r.Index(); got != tt.want {
				t.Errorf("Rotor.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_ApplyF(t *testing.T) {
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
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
			want: []byte{
				63, 180, 255, 162, 59, 142, 61, 13, 187, 226, 49, 134, 163, 38, 44, 14,
				255, 73, 155, 237, 208, 42, 217, 227, 194, 245, 229, 169, 96, 163, 33, 145,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := proFormaRotors[0]
			r.Current = 0
			if got := r.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_ApplyG(t *testing.T) {
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
					63, 180, 255, 162, 59, 142, 61, 13, 187, 226, 49, 134, 163, 38, 44, 14,
					255, 73, 155, 237, 208, 42, 217, 227, 194, 245, 229, 169, 96, 163, 33, 145,
				},
			},
			want: []byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := proFormaRotors[0]
			r.Current = 0
			if got := r.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.ApplyG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_String(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "trs1",
			want: "new(Rotor).New(1789, 1065, 1499, []byte{\n" +
				"	63, 180, 255, 162, 59, 142, 61, 13, 187, 226, 49, 134, 163, 38, 44, 14,\n" +
				"	255, 73, 155, 237, 208, 42, 217, 227, 194, 245, 229, 169, 96, 163, 33, 145,\n" +
				"	222, 156, 57, 87, 220, 186, 118, 131, 89, 103, 27, 145, 153, 207, 16, 55,\n" +
				"	248, 183, 83, 65, 15, 253, 147, 136, 217, 189, 124, 150, 193, 113, 87, 127,\n" +
				"	101, 202, 87, 3, 80, 160, 132, 129, 1, 134, 154, 36, 194, 3, 186, 148,\n" +
				"	241, 226, 134, 255, 59, 78, 202, 236, 166, 151, 184, 209, 115, 21, 177, 17,\n" +
				"	106, 189, 209, 128, 13, 224, 94, 163, 47, 117, 151, 3, 9, 88, 20, 74,\n" +
				"	188, 243, 174, 130, 193, 247, 161, 74, 119, 95, 40, 111, 215, 174, 84, 170,\n" +
				"	234, 27, 241, 147, 210, 26, 139, 92, 231, 118, 227, 206, 0, 186, 161, 82,\n" +
				"	149, 59, 93, 134, 84, 108, 116, 191, 127, 153, 92, 59, 80, 53, 10, 112,\n" +
				"	127, 228, 183, 134, 214, 74, 150, 134, 145, 60, 22, 217, 213, 195, 251, 240,\n" +
				"	232, 1, 193, 235, 142, 191, 153, 123, 46, 86, 198, 123, 33, 34, 148, 104,\n" +
				"	18, 96, 34, 17, 139, 199, 225, 84, 245, 102, 137, 167, 240, 84, 152, 144,\n" +
				"	171, 21, 67, 253, 113, 97, 156, 145, 55, 87, 247, 45, 54, 48, 157, 247,\n" +
				"	135, 246, 95, 116, 199, 177, 167, 97, 87, 60, 198, 112, 212, 132, 197, 225,\n" +
				"	63, 105, 179, 29, 90, 37, 123, 92, 184, 190, 60, 21, 108, 52, 36, 18})\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := proFormaRotors[0]
			if got := r.String(); got != tt.want {
				t.Errorf("Rotor.String() = %v, want = %v", got, tt.want)
			}
		})
	}
}
