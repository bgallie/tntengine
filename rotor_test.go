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
			want: new(Rotor).New(1789, 1679, 1025, []byte{
				161, 48, 75, 216, 12, 30, 103, 205, 176, 240, 35, 55, 247, 159, 136, 37,
				70, 238, 138, 224, 211, 134, 89, 95, 212, 157, 47, 180, 245, 128, 10, 107,
				178, 137, 133, 228, 14, 19, 53, 135, 155, 246, 134, 231, 116, 106, 89, 75,
				121, 212, 240, 251, 194, 74, 26, 1, 127, 16, 201, 120, 69, 100, 13, 27,
				244, 133, 46, 254, 210, 37, 13, 31, 82, 101, 45, 40, 29, 139, 247, 54,
				46, 0, 83, 176, 193, 229, 145, 15, 166, 10, 235, 220, 27, 166, 199, 228,
				123, 227, 6, 18, 5, 141, 121, 9, 141, 168, 177, 173, 211, 175, 213, 71,
				138, 138, 12, 131, 203, 3, 19, 221, 245, 99, 69, 4, 211, 49, 47, 244,
				215, 196, 167, 133, 95, 227, 254, 114, 54, 223, 233, 51, 45, 48, 114, 171,
				208, 154, 232, 199, 26, 217, 19, 21, 70, 49, 124, 38, 125, 227, 231, 144,
				61, 231, 239, 91, 78, 160, 16, 232, 53, 118, 4, 170, 234, 70, 99, 8,
				235, 239, 227, 84, 47, 132, 214, 21, 108, 142, 153, 98, 228, 34, 179, 185,
				218, 125, 80, 88, 230, 35, 80, 109, 154, 157, 132, 246, 221, 172, 98, 182,
				171, 113, 135, 214, 94, 53, 117, 62, 195, 21, 110, 82, 237, 71, 28, 60,
				20, 102, 9, 155, 193, 227, 172, 25, 22, 126, 228, 230, 254, 19, 177, 196,
				200, 93, 17, 124, 218, 48, 235, 139, 186, 243, 133, 182, 30, 80, 97, 13}),
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
