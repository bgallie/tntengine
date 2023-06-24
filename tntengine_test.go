// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestTntEngine_Left(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.BuildCipherMachine()
	tests := []struct {
		name string
		want chan CipherBlock
	}{
		{
			name: "ttel1",
			want: tntMachine.left,
		},
	}
	for _, tt := range tests {
		e := tntMachine
		t.Run(tt.name, func(t *testing.T) {
			if got := e.Left(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TntEngine.Left() = %v, want %v", got, tt.want)
			}
		})
	}
	var blk CipherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestTntEngine_Right(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.BuildCipherMachine()
	tests := []struct {
		name string
		want chan CipherBlock
	}{
		{
			name: "tter1",
			want: tntMachine.right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			if got := e.Right(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TntEngine.Right() = %v, want %v", got, tt.want)
			}
		})
	}
	var blk CipherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestTntEngine_CounterKey(t *testing.T) {
	var tntMachine TntEngine
	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "ttec1",
			key:  "SecretKey",
			want: "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tntMachine.Init([]byte(tt.key))
			if got := tntMachine.CounterKey(); got != tt.want {
				t.Errorf("TntEngine.CounterKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTntEngine_Index(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	tntMachine.SetIndex(iCnt)
	tests := []struct {
		name     string
		wantCntr *big.Int
	}{
		{
			name:     "ttei1",
			wantCntr: iCnt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			if gotCntr := e.Index(); !reflect.DeepEqual(gotCntr, tt.wantCntr) {
				t.Errorf("TntEngine.Index() = %v, want %v", gotCntr, tt.wantCntr)
			}
		})
	}
}

func TestTntEngine_SetIndex(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	type args struct {
		iCnt *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "ttesi1",
			args: args{iCnt},
			want: iCnt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			e.SetIndex(tt.args.iCnt)
			if got := e.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TntEngine.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTntEngine_SetEngineType(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	type args struct {
		engineType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "tteset1",
			args: args{"E"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			e.SetEngineType(tt.args.engineType)
			if got := e.engineType; got != tt.args.engineType {
				t.Errorf("TntEngine.SetEngineType() = %v, want %v", got, tt.args.engineType)
			}
		})
	}
}

func TestTntEngine_Engine(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetIndex(BigZero)
	cnter := new(Counter)
	cnter.SetIndex(BigZero)
	tests := []struct {
		name string
		want []Crypter
	}{
		{
			name: "tnte1",
			want: []Crypter{
				new(Rotor).New(1753, 175, 1630, []byte{
					208, 73, 91, 124, 80, 77, 179, 203, 35, 146, 21, 89, 206, 211, 112, 182,
					33, 157, 109, 103, 201, 62, 145, 122, 69, 53, 177, 81, 236, 132, 42, 45,
					165, 56, 71, 91, 34, 210, 11, 173, 33, 234, 230, 232, 155, 197, 119, 225,
					193, 233, 156, 247, 186, 234, 89, 38, 229, 230, 86, 74, 11, 111, 122, 157,
					156, 151, 227, 196, 206, 196, 37, 107, 111, 255, 202, 54, 248, 178, 142, 234,
					0, 74, 184, 20, 7, 17, 170, 246, 152, 103, 177, 185, 31, 102, 122, 34,
					7, 37, 42, 31, 202, 228, 143, 232, 255, 32, 158, 66, 190, 200, 176, 125,
					248, 73, 126, 87, 0, 124, 62, 242, 12, 90, 221, 207, 16, 246, 70, 224,
					42, 11, 66, 171, 52, 58, 152, 16, 12, 5, 138, 103, 143, 13, 43, 24,
					93, 10, 201, 178, 227, 166, 63, 166, 44, 232, 65, 12, 231, 161, 209, 51,
					126, 78, 150, 220, 145, 69, 125, 194, 197, 245, 227, 29, 108, 36, 112, 112,
					184, 170, 76, 108, 223, 132, 241, 76, 98, 13, 75, 0, 92, 176, 69, 14,
					63, 121, 165, 165, 149, 49, 235, 254, 57, 38, 163, 100, 101, 155, 178, 157,
					10, 232, 123, 158, 230, 33, 198, 246, 57, 125, 32, 160, 147, 182, 248, 160,
					154, 102, 151, 71, 36, 43, 178, 156, 167, 225, 108, 67, 58, 219, 206, 146,
					125, 34, 245, 138, 106, 98, 163, 216, 9, 85, 90, 176, 131, 40, 180, 168}),
				new(Rotor).New(1789, 60, 444, []byte{
					153, 247, 84, 55, 230, 25, 164, 59, 126, 236, 221, 4, 209, 236, 84, 230,
					139, 111, 14, 137, 166, 224, 21, 71, 61, 4, 55, 178, 10, 241, 108, 204,
					229, 46, 100, 97, 187, 8, 111, 222, 202, 106, 189, 168, 13, 181, 118, 253,
					43, 150, 161, 40, 126, 233, 0, 54, 55, 210, 74, 111, 5, 82, 47, 112,
					26, 131, 8, 189, 195, 11, 151, 141, 200, 190, 217, 133, 65, 222, 160, 17,
					44, 232, 159, 225, 245, 125, 118, 66, 216, 192, 61, 218, 116, 2, 226, 9,
					72, 67, 215, 216, 198, 104, 221, 241, 168, 224, 17, 64, 59, 199, 214, 3,
					201, 216, 211, 128, 22, 161, 185, 192, 106, 78, 34, 196, 7, 106, 138, 160,
					207, 31, 16, 243, 134, 12, 20, 176, 104, 235, 86, 122, 175, 14, 121, 142,
					16, 249, 201, 10, 1, 8, 56, 74, 175, 200, 86, 5, 92, 167, 83, 227,
					52, 24, 20, 93, 183, 161, 250, 148, 17, 232, 31, 194, 23, 76, 70, 43,
					247, 124, 76, 64, 210, 181, 209, 78, 80, 141, 103, 69, 34, 75, 50, 30,
					76, 55, 111, 164, 133, 26, 223, 59, 10, 92, 133, 147, 142, 0, 116, 141,
					229, 65, 233, 130, 71, 141, 147, 157, 207, 27, 93, 221, 135, 196, 30, 47,
					243, 158, 234, 198, 60, 131, 116, 199, 143, 189, 155, 32, 154, 157, 202, 124,
					241, 205, 33, 209, 20, 188, 226, 168, 135, 224, 70, 86, 33, 158, 141, 249}),
				new(Permutator).New(256, []byte{
					95, 35, 229, 40, 44, 144, 131, 167, 39, 54, 121, 118, 94, 69, 223, 53,
					156, 210, 65, 187, 241, 101, 66, 67, 123, 201, 119, 14, 37, 115, 108, 205,
					173, 171, 231, 208, 226, 185, 129, 158, 81, 19, 190, 91, 186, 206, 239, 179,
					255, 42, 85, 46, 15, 192, 122, 57, 38, 189, 1, 23, 49, 242, 55, 166,
					8, 58, 36, 9, 180, 150, 92, 244, 139, 217, 176, 0, 140, 93, 30, 72,
					254, 199, 163, 151, 61, 228, 116, 17, 77, 110, 126, 64, 50, 184, 252, 56,
					111, 235, 2, 146, 41, 43, 249, 34, 143, 112, 134, 5, 132, 164, 63, 138,
					232, 207, 162, 227, 197, 165, 161, 220, 82, 105, 253, 174, 181, 175, 4, 149,
					100, 196, 48, 11, 79, 70, 10, 152, 74, 26, 25, 250, 75, 21, 33, 136,
					212, 155, 230, 245, 234, 145, 62, 117, 28, 45, 99, 183, 203, 137, 221, 80,
					120, 109, 73, 97, 170, 141, 29, 247, 246, 13, 68, 125, 209, 113, 172, 124,
					237, 198, 248, 130, 59, 214, 233, 71, 178, 16, 218, 6, 215, 128, 219, 204,
					114, 211, 96, 224, 90, 193, 106, 60, 236, 191, 89, 51, 213, 107, 160, 148,
					31, 159, 200, 76, 251, 18, 32, 127, 98, 154, 104, 157, 84, 222, 225, 86,
					182, 142, 202, 194, 22, 24, 7, 133, 195, 147, 88, 177, 188, 238, 103, 87,
					12, 78, 240, 47, 243, 27, 153, 169, 216, 52, 3, 20, 83, 135, 102, 168}),
				new(Rotor).New(1777, 1356, 175, []byte{
					184, 70, 115, 167, 255, 132, 209, 239, 51, 223, 64, 14, 59, 13, 235, 3,
					194, 36, 57, 227, 177, 80, 131, 110, 64, 134, 29, 57, 205, 63, 233, 230,
					113, 155, 231, 226, 3, 205, 205, 41, 226, 43, 203, 66, 228, 43, 253, 60,
					133, 242, 124, 93, 98, 21, 99, 163, 28, 41, 96, 195, 214, 72, 243, 137,
					48, 255, 13, 217, 128, 43, 109, 108, 98, 228, 106, 240, 114, 106, 243, 101,
					217, 227, 189, 78, 227, 12, 105, 29, 184, 243, 38, 21, 142, 242, 171, 37,
					101, 182, 150, 16, 67, 172, 210, 113, 52, 97, 214, 68, 154, 157, 4, 11,
					161, 168, 162, 207, 196, 37, 203, 105, 187, 225, 110, 2, 230, 132, 47, 101,
					111, 85, 28, 149, 183, 0, 253, 30, 177, 90, 249, 170, 31, 221, 145, 5,
					132, 180, 173, 52, 206, 9, 199, 71, 89, 101, 30, 26, 19, 22, 40, 66,
					219, 90, 24, 38, 63, 38, 13, 109, 152, 3, 15, 27, 216, 247, 48, 8,
					187, 98, 103, 47, 16, 29, 79, 25, 20, 157, 44, 120, 52, 193, 200, 244,
					107, 56, 73, 211, 96, 129, 29, 215, 111, 58, 63, 247, 214, 75, 167, 225,
					18, 50, 56, 89, 131, 222, 241, 211, 90, 51, 95, 42, 93, 81, 113, 141,
					230, 78, 255, 9, 163, 223, 103, 190, 129, 28, 118, 26, 214, 7, 132, 73,
					114, 198, 99, 161, 6, 221, 128, 12, 59, 114, 154, 127, 210, 205, 1, 152}),
				new(Rotor).New(1759, 555, 890, []byte{
					73, 82, 157, 119, 21, 201, 186, 243, 15, 28, 109, 175, 153, 178, 80, 87,
					116, 136, 226, 139, 78, 251, 2, 69, 65, 209, 93, 155, 39, 231, 11, 176,
					82, 233, 42, 233, 10, 48, 72, 70, 36, 155, 223, 255, 218, 123, 187, 251,
					216, 173, 67, 18, 9, 223, 248, 215, 218, 216, 9, 219, 205, 219, 71, 249,
					141, 126, 90, 219, 180, 54, 206, 201, 197, 5, 4, 50, 130, 17, 247, 29,
					10, 0, 250, 13, 225, 196, 172, 5, 183, 7, 207, 47, 129, 6, 245, 239,
					225, 13, 164, 145, 106, 54, 34, 160, 172, 62, 223, 218, 228, 17, 153, 195,
					110, 28, 189, 122, 38, 169, 248, 153, 252, 242, 227, 252, 84, 30, 200, 241,
					84, 74, 94, 24, 66, 199, 31, 198, 115, 98, 127, 193, 254, 220, 90, 250,
					73, 25, 50, 60, 183, 156, 175, 80, 91, 230, 23, 71, 71, 162, 5, 139,
					91, 42, 182, 251, 250, 141, 13, 138, 240, 100, 243, 62, 210, 252, 76, 23,
					78, 132, 213, 81, 241, 122, 60, 64, 69, 38, 212, 96, 252, 101, 251, 0,
					25, 31, 181, 156, 44, 139, 13, 77, 72, 138, 197, 59, 244, 97, 92, 245,
					71, 233, 168, 88, 124, 162, 17, 140, 175, 111, 2, 192, 36, 169, 206, 187,
					138, 100, 221, 249, 7, 142, 182, 215, 76, 89, 168, 43, 58, 68, 241, 69,
					167, 125, 129, 162, 160, 232, 174, 205, 147, 243, 5, 216, 148, 15, 55, 222}),
				new(Permutator).New(256, []byte{
					95, 35, 229, 40, 44, 144, 131, 167, 39, 54, 121, 118, 94, 69, 223, 53,
					156, 210, 65, 187, 241, 101, 66, 67, 123, 201, 119, 14, 37, 115, 108, 205,
					173, 171, 231, 208, 226, 185, 129, 158, 81, 19, 190, 91, 186, 206, 239, 179,
					255, 42, 85, 46, 15, 192, 122, 57, 38, 189, 1, 23, 49, 242, 55, 166,
					8, 58, 36, 9, 180, 150, 92, 244, 139, 217, 176, 0, 140, 93, 30, 72,
					254, 199, 163, 151, 61, 228, 116, 17, 77, 110, 126, 64, 50, 184, 252, 56,
					111, 235, 2, 146, 41, 43, 249, 34, 143, 112, 134, 5, 132, 164, 63, 138,
					232, 207, 162, 227, 197, 165, 161, 220, 82, 105, 253, 174, 181, 175, 4, 149,
					100, 196, 48, 11, 79, 70, 10, 152, 74, 26, 25, 250, 75, 21, 33, 136,
					212, 155, 230, 245, 234, 145, 62, 117, 28, 45, 99, 183, 203, 137, 221, 80,
					120, 109, 73, 97, 170, 141, 29, 247, 246, 13, 68, 125, 209, 113, 172, 124,
					237, 198, 248, 130, 59, 214, 233, 71, 178, 16, 218, 6, 215, 128, 219, 204,
					114, 211, 96, 224, 90, 193, 106, 60, 236, 191, 89, 51, 213, 107, 160, 148,
					31, 159, 200, 76, 251, 18, 32, 127, 98, 154, 104, 157, 84, 222, 225, 86,
					182, 142, 202, 194, 22, 24, 7, 133, 195, 147, 88, 177, 188, 238, 103, 87,
					12, 78, 240, 47, 243, 27, 153, 169, 216, 52, 3, 20, 83, 135, 102, 168}),
				new(Rotor).New(1787, 299, 1184, []byte{
					237, 65, 181, 11, 141, 148, 228, 255, 37, 147, 238, 78, 193, 183, 45, 85,
					44, 240, 12, 117, 182, 179, 62, 187, 43, 205, 109, 113, 231, 101, 26, 103,
					210, 122, 16, 45, 87, 189, 114, 209, 86, 23, 236, 47, 2, 61, 81, 70,
					94, 115, 133, 92, 62, 250, 108, 164, 63, 166, 61, 248, 85, 253, 208, 218,
					163, 24, 36, 171, 85, 9, 164, 6, 8, 167, 153, 29, 43, 244, 139, 15,
					81, 222, 187, 28, 241, 165, 83, 112, 220, 67, 10, 228, 236, 186, 90, 195,
					162, 108, 157, 194, 45, 10, 24, 202, 27, 222, 93, 114, 7, 109, 212, 248,
					222, 208, 228, 221, 164, 198, 202, 234, 218, 20, 2, 184, 122, 6, 37, 94,
					143, 55, 150, 122, 11, 62, 242, 54, 86, 184, 180, 114, 203, 215, 131, 215,
					108, 240, 250, 93, 5, 65, 10, 8, 98, 255, 181, 177, 32, 14, 67, 143,
					82, 42, 7, 235, 82, 78, 36, 22, 194, 208, 117, 77, 146, 33, 186, 114,
					174, 7, 160, 253, 79, 59, 189, 202, 210, 175, 176, 172, 156, 177, 130, 217,
					114, 106, 234, 193, 193, 13, 195, 202, 201, 244, 25, 121, 116, 51, 24, 135,
					255, 197, 2, 171, 26, 43, 84, 65, 197, 177, 115, 250, 32, 46, 246, 107,
					15, 170, 93, 104, 164, 36, 255, 47, 153, 116, 119, 10, 190, 109, 169, 98,
					129, 103, 168, 179, 157, 245, 217, 93, 105, 110, 139, 59, 47, 211, 56, 211}),
				new(Rotor).New(1783, 285, 1183, []byte{
					65, 195, 116, 124, 166, 169, 204, 117, 120, 26, 62, 143, 136, 111, 26, 246,
					217, 250, 214, 28, 60, 178, 219, 221, 184, 64, 64, 66, 151, 6, 123, 135,
					226, 83, 44, 91, 49, 78, 176, 246, 100, 149, 12, 145, 188, 102, 143, 253,
					11, 94, 179, 186, 150, 252, 93, 32, 117, 238, 46, 160, 127, 18, 203, 198,
					177, 153, 11, 137, 37, 14, 169, 14, 236, 114, 83, 102, 254, 169, 131, 108,
					192, 47, 81, 153, 42, 41, 142, 86, 210, 60, 34, 107, 203, 198, 21, 198,
					177, 33, 255, 11, 250, 157, 96, 46, 166, 73, 93, 21, 179, 171, 144, 85,
					137, 127, 166, 51, 253, 49, 85, 17, 253, 185, 84, 179, 227, 130, 53, 95,
					44, 72, 215, 106, 33, 184, 189, 126, 221, 77, 144, 142, 217, 95, 227, 8,
					234, 78, 87, 237, 146, 147, 189, 44, 165, 111, 60, 166, 128, 82, 136, 135,
					63, 133, 21, 186, 185, 56, 254, 217, 137, 243, 57, 32, 112, 251, 56, 66,
					161, 212, 87, 38, 19, 100, 195, 117, 249, 26, 53, 172, 11, 131, 137, 2,
					214, 17, 56, 205, 9, 41, 233, 221, 152, 69, 82, 6, 142, 118, 148, 208,
					220, 118, 234, 243, 173, 183, 87, 245, 28, 141, 90, 133, 23, 139, 195, 160,
					97, 58, 62, 211, 84, 230, 58, 60, 13, 159, 71, 196, 55, 13, 251, 108,
					125, 107, 14, 30, 217, 237, 110, 92, 32, 32, 161, 75, 131, 189, 195, 209}),
				cnter},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tntMachine.engine; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TntEngine.Engine() = %v, want %v", got, tt.want[:8])
			}
		})
	}
}

func TestTntEngine_MaximalStates(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	want, _ := new(big.Int).SetString("2046922266175282266177536", 10)
	tests := []struct {
		name string
		want *big.Int
	}{
		{
			name: "tteset1",
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			if got := e.MaximalStates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TntEngine.MaximalStates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createProFormaMachine(t *testing.T) {
	tests := []struct {
		name string
		want *[]Crypter
	}{
		{
			name: "tcpfm1",
			want: &[]Crypter{Rotor1, Rotor2, Permutator1, Rotor3, Rotor4, Permutator1, Rotor5, Rotor6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createProFormaMachine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createProFormaMachine() = %v, want %v", got, tt.want)
			}
		})
	}
}
