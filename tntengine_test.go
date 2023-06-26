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
				new(Rotor).New(1787, 1379, 1579, []byte{
					148, 120, 9, 223, 187, 219, 81, 176, 67, 255, 147, 198, 62, 191, 137, 248,
					193, 125, 255, 243, 71, 239, 195, 192, 66, 213, 100, 195, 162, 116, 210, 227,
					197, 26, 4, 167, 139, 117, 116, 114, 144, 217, 182, 55, 154, 43, 7, 236,
					216, 131, 136, 236, 113, 246, 143, 183, 114, 71, 145, 42, 78, 247, 177, 209,
					222, 108, 233, 93, 88, 227, 215, 165, 179, 155, 173, 114, 159, 75, 242, 29,
					162, 143, 237, 228, 226, 37, 251, 235, 46, 55, 172, 120, 48, 10, 158, 180,
					165, 70, 19, 9, 100, 68, 124, 57, 98, 65, 229, 200, 206, 24, 211, 236,
					2, 41, 182, 159, 149, 158, 246, 151, 43, 77, 199, 166, 157, 112, 107, 80,
					151, 188, 19, 204, 141, 194, 57, 54, 134, 130, 208, 37, 126, 218, 223, 111,
					182, 74, 220, 128, 138, 95, 53, 200, 193, 172, 2, 131, 26, 243, 117, 234,
					157, 164, 41, 193, 152, 179, 107, 81, 18, 204, 211, 21, 156, 150, 53, 187,
					228, 31, 94, 6, 194, 178, 65, 180, 22, 201, 9, 198, 31, 124, 35, 75,
					109, 105, 39, 85, 175, 11, 80, 86, 157, 181, 2, 56, 76, 156, 30, 117,
					236, 161, 252, 198, 34, 239, 141, 0, 169, 74, 107, 206, 177, 145, 87, 160,
					196, 75, 248, 222, 221, 142, 130, 29, 250, 159, 52, 246, 249, 77, 196, 15,
					238, 251, 159, 63, 122, 31, 6, 22, 170, 38, 27, 22, 165, 147, 30, 47}),
				new(Rotor).New(1759, 398, 1058, []byte{
					159, 190, 152, 254, 176, 246, 254, 249, 167, 93, 113, 172, 188, 86, 11, 125,
					9, 132, 47, 111, 130, 160, 251, 52, 176, 246, 62, 170, 202, 243, 73, 173,
					35, 16, 106, 47, 131, 155, 173, 27, 238, 183, 52, 252, 75, 12, 16, 234,
					54, 165, 109, 64, 100, 109, 10, 87, 16, 198, 97, 193, 14, 167, 32, 60,
					8, 21, 221, 88, 100, 176, 226, 228, 73, 1, 228, 230, 188, 82, 170, 167,
					54, 254, 9, 173, 41, 141, 76, 139, 188, 229, 250, 192, 85, 56, 124, 49,
					117, 63, 103, 235, 200, 132, 236, 54, 145, 78, 241, 183, 159, 176, 83, 36,
					163, 136, 134, 173, 78, 21, 79, 116, 216, 39, 171, 183, 65, 25, 246, 164,
					58, 72, 227, 143, 94, 148, 113, 153, 254, 187, 55, 157, 154, 105, 226, 180,
					10, 53, 192, 123, 253, 153, 198, 182, 56, 239, 184, 250, 119, 138, 97, 189,
					188, 232, 104, 1, 3, 51, 241, 207, 215, 128, 161, 156, 244, 121, 109, 91,
					241, 4, 32, 2, 157, 173, 237, 251, 176, 168, 74, 253, 138, 174, 47, 112,
					68, 79, 110, 188, 163, 7, 6, 88, 124, 188, 29, 186, 177, 51, 125, 0,
					217, 26, 248, 80, 160, 8, 79, 170, 87, 188, 125, 165, 79, 95, 76, 127,
					88, 123, 255, 252, 211, 174, 56, 86, 94, 171, 133, 190, 4, 194, 151, 55,
					65, 208, 125, 26, 88, 123, 31, 85, 229, 249, 164, 86, 161, 171, 41, 68}),
				new(Permutator).New(256, []byte{
					2, 54, 236, 59, 71, 24, 38, 151, 82, 45, 214, 178, 37, 135, 46, 72,
					226, 87, 160, 25, 188, 122, 155, 109, 57, 95, 127, 93, 23, 202, 159, 217,
					175, 213, 198, 210, 34, 44, 230, 207, 197, 30, 77, 36, 120, 29, 154, 153,
					194, 62, 58, 240, 255, 158, 80, 123, 139, 114, 163, 149, 252, 136, 65, 143,
					215, 61, 88, 220, 182, 204, 205, 171, 174, 164, 73, 208, 78, 199, 8, 116,
					35, 250, 157, 172, 167, 113, 144, 166, 173, 242, 219, 203, 165, 132, 104, 195,
					92, 227, 251, 11, 129, 15, 209, 168, 76, 161, 243, 41, 22, 118, 111, 212,
					26, 5, 53, 187, 124, 231, 239, 97, 200, 229, 191, 10, 169, 98, 196, 145,
					179, 89, 4, 225, 216, 3, 128, 9, 112, 96, 186, 133, 221, 148, 162, 223,
					68, 237, 40, 28, 81, 121, 74, 52, 48, 189, 42, 99, 235, 150, 218, 241,
					101, 125, 108, 253, 248, 152, 137, 67, 55, 84, 211, 110, 50, 249, 134, 60,
					245, 185, 126, 56, 63, 206, 106, 18, 147, 20, 107, 69, 0, 234, 115, 176,
					21, 192, 117, 14, 16, 140, 170, 47, 138, 247, 105, 49, 228, 100, 64, 102,
					184, 7, 66, 233, 90, 183, 13, 232, 94, 12, 156, 91, 254, 246, 141, 27,
					222, 6, 19, 39, 224, 190, 238, 31, 32, 51, 33, 79, 146, 131, 119, 201,
					103, 180, 75, 130, 86, 83, 142, 17, 244, 85, 70, 181, 1, 177, 193, 43}),
				new(Rotor).New(1789, 60, 444, []byte{
					153, 247, 84, 55, 230, 25, 164, 59, 126, 236, 221, 4, 209, 236, 84, 230,
					139, 111, 14, 137, 166, 224, 21, 71, 61, 4, 139, 183, 101, 78, 5, 174,
					94, 15, 206, 125, 202, 48, 228, 146, 191, 224, 132, 235, 224, 60, 18, 57,
					21, 103, 237, 61, 115, 105, 9, 46, 125, 96, 212, 84, 255, 99, 31, 143,
					174, 229, 231, 55, 87, 140, 24, 250, 254, 223, 11, 189, 225, 224, 128, 159,
					129, 148, 123, 81, 207, 187, 23, 163, 124, 181, 23, 212, 9, 96, 174, 161,
					11, 53, 94, 248, 208, 200, 122, 100, 102, 100, 45, 238, 228, 196, 203, 100,
					117, 33, 216, 253, 197, 158, 136, 53, 250, 119, 30, 142, 163, 134, 22, 166,
					102, 243, 12, 64, 209, 28, 91, 75, 121, 204, 74, 86, 124, 140, 230, 173,
					195, 100, 183, 11, 232, 204, 89, 52, 21, 92, 150, 81, 254, 153, 112, 222,
					32, 98, 73, 34, 119, 148, 167, 165, 12, 11, 53, 195, 237, 110, 79, 173,
					225, 142, 171, 222, 45, 213, 183, 40, 170, 51, 39, 178, 180, 11, 10, 218,
					75, 126, 31, 114, 113, 104, 95, 47, 116, 114, 106, 239, 24, 178, 51, 215,
					50, 253, 234, 33, 107, 175, 142, 70, 143, 130, 146, 226, 96, 45, 123, 45,
					243, 158, 234, 198, 60, 131, 116, 199, 143, 189, 155, 32, 154, 157, 202, 124,
					241, 205, 33, 209, 20, 188, 226, 168, 135, 96, 241, 182, 204, 169, 192, 245}),
				new(Rotor).New(1753, 753, 408, []byte{
					136, 191, 189, 56, 116, 210, 228, 111, 198, 3, 198, 21, 42, 121, 197, 79,
					63, 180, 253, 253, 172, 101, 180, 152, 99, 189, 4, 104, 132, 178, 53, 215,
					92, 70, 70, 87, 94, 39, 64, 240, 15, 107, 137, 146, 30, 85, 105, 70,
					211, 197, 82, 109, 206, 194, 82, 224, 205, 25, 251, 4, 105, 216, 230, 12,
					100, 169, 249, 236, 74, 196, 227, 117, 11, 49, 154, 165, 43, 49, 189, 157,
					130, 225, 5, 185, 183, 190, 153, 171, 197, 72, 231, 95, 173, 227, 173, 220,
					79, 166, 226, 97, 156, 197, 93, 27, 185, 205, 8, 240, 224, 7, 6, 6,
					121, 152, 121, 250, 18, 152, 102, 120, 26, 156, 77, 198, 136, 234, 181, 197,
					238, 60, 135, 211, 217, 91, 115, 121, 91, 53, 250, 217, 219, 129, 216, 120,
					238, 237, 196, 33, 185, 70, 218, 171, 73, 217, 141, 243, 74, 143, 57, 28,
					240, 163, 145, 78, 0, 188, 23, 195, 168, 227, 100, 59, 200, 89, 179, 96,
					7, 96, 40, 118, 38, 180, 118, 227, 127, 209, 16, 15, 79, 22, 26, 95,
					103, 19, 190, 30, 134, 196, 82, 104, 215, 182, 68, 135, 73, 187, 98, 103,
					70, 228, 20, 224, 63, 80, 255, 200, 31, 107, 36, 16, 127, 123, 113, 232,
					164, 201, 223, 140, 7, 140, 43, 84, 242, 138, 159, 126, 104, 251, 251, 89,
					203, 104, 49, 199, 122, 9, 208, 8, 101, 107, 174, 193, 87, 68, 97, 164}),
				new(Permutator).New(256, []byte{
					2, 54, 236, 59, 71, 24, 38, 151, 82, 45, 214, 178, 37, 135, 46, 72,
					226, 87, 160, 25, 188, 122, 155, 109, 57, 95, 127, 93, 23, 202, 159, 217,
					175, 213, 198, 210, 34, 44, 230, 207, 197, 30, 77, 36, 120, 29, 154, 153,
					194, 62, 58, 240, 255, 158, 80, 123, 139, 114, 163, 149, 252, 136, 65, 143,
					215, 61, 88, 220, 182, 204, 205, 171, 174, 164, 73, 208, 78, 199, 8, 116,
					35, 250, 157, 172, 167, 113, 144, 166, 173, 242, 219, 203, 165, 132, 104, 195,
					92, 227, 251, 11, 129, 15, 209, 168, 76, 161, 243, 41, 22, 118, 111, 212,
					26, 5, 53, 187, 124, 231, 239, 97, 200, 229, 191, 10, 169, 98, 196, 145,
					179, 89, 4, 225, 216, 3, 128, 9, 112, 96, 186, 133, 221, 148, 162, 223,
					68, 237, 40, 28, 81, 121, 74, 52, 48, 189, 42, 99, 235, 150, 218, 241,
					101, 125, 108, 253, 248, 152, 137, 67, 55, 84, 211, 110, 50, 249, 134, 60,
					245, 185, 126, 56, 63, 206, 106, 18, 147, 20, 107, 69, 0, 234, 115, 176,
					21, 192, 117, 14, 16, 140, 170, 47, 138, 247, 105, 49, 228, 100, 64, 102,
					184, 7, 66, 233, 90, 183, 13, 232, 94, 12, 156, 91, 254, 246, 141, 27,
					222, 6, 19, 39, 224, 190, 238, 31, 32, 51, 33, 79, 146, 131, 119, 201,
					103, 180, 75, 130, 86, 83, 142, 17, 244, 85, 70, 181, 1, 177, 193, 43}),
				new(Rotor).New(1777, 1460, 1034, []byte{
					185, 71, 250, 58, 68, 93, 60, 193, 109, 52, 219, 97, 233, 10, 158, 27,
					152, 52, 57, 137, 24, 14, 178, 182, 230, 20, 20, 93, 64, 172, 173, 253,
					33, 171, 102, 177, 131, 119, 70, 18, 82, 64, 67, 96, 243, 14, 79, 240,
					220, 95, 21, 237, 214, 175, 208, 92, 6, 144, 179, 236, 55, 12, 12, 146,
					201, 218, 22, 33, 51, 131, 100, 23, 205, 95, 56, 34, 154, 249, 228, 121,
					205, 84, 110, 177, 226, 70, 140, 232, 30, 87, 115, 147, 141, 192, 83, 171,
					241, 82, 130, 208, 96, 35, 167, 90, 2, 45, 106, 159, 107, 20, 251, 33,
					182, 221, 218, 157, 182, 131, 182, 248, 77, 196, 193, 193, 165, 110, 132, 173,
					115, 17, 230, 100, 89, 232, 228, 2, 173, 119, 140, 144, 117, 127, 148, 226,
					218, 16, 248, 187, 115, 254, 5, 15, 156, 82, 198, 179, 87, 87, 77, 40,
					193, 184, 241, 184, 177, 22, 176, 2, 53, 51, 13, 95, 74, 60, 223, 212,
					113, 173, 110, 14, 172, 56, 124, 80, 164, 18, 125, 93, 37, 61, 44, 47,
					55, 144, 187, 95, 82, 184, 151, 185, 115, 194, 239, 169, 118, 166, 19, 185,
					81, 238, 99, 207, 57, 20, 105, 93, 67, 251, 15, 63, 118, 102, 115, 143,
					244, 117, 136, 186, 120, 130, 219, 104, 182, 195, 210, 21, 60, 55, 48, 105,
					114, 18, 49, 28, 100, 109, 205, 41, 40, 186, 128, 88, 91, 251, 1, 243}),
				new(Rotor).New(1783, 201, 1696, []byte{
					162, 147, 81, 248, 75, 180, 188, 113, 27, 252, 48, 249, 253, 40, 61, 171,
					51, 229, 49, 227, 248, 93, 26, 177, 107, 230, 69, 140, 45, 105, 52, 157,
					151, 123, 129, 98, 159, 81, 177, 194, 29, 184, 224, 104, 216, 6, 120, 28,
					88, 46, 5, 132, 100, 96, 98, 183, 249, 113, 85, 32, 246, 196, 46, 202,
					155, 28, 231, 171, 50, 174, 54, 157, 194, 69, 59, 26, 70, 29, 126, 152,
					72, 98, 49, 171, 133, 161, 66, 109, 153, 81, 83, 176, 65, 11, 9, 249,
					227, 224, 115, 114, 160, 142, 46, 120, 164, 48, 177, 134, 185, 242, 173, 208,
					141, 91, 54, 137, 64, 252, 52, 6, 179, 54, 156, 165, 11, 7, 105, 55,
					237, 35, 248, 124, 67, 161, 145, 136, 89, 248, 26, 203, 104, 196, 234, 139,
					220, 93, 169, 143, 177, 194, 34, 195, 161, 5, 8, 229, 12, 216, 133, 51,
					203, 168, 230, 32, 92, 171, 145, 109, 117, 31, 224, 245, 199, 253, 184, 174,
					152, 191, 81, 90, 224, 66, 117, 18, 117, 197, 202, 110, 85, 232, 24, 128,
					245, 75, 164, 202, 64, 178, 241, 81, 77, 197, 119, 208, 125, 33, 234, 113,
					111, 62, 53, 182, 15, 139, 51, 135, 47, 187, 0, 80, 155, 62, 55, 209,
					201, 40, 252, 37, 90, 222, 184, 13, 126, 152, 252, 126, 148, 158, 213, 153,
					242, 152, 113, 252, 46, 141, 216, 53, 243, 34, 198, 150, 52, 154, 78, 9}),
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
	want, _ := new(big.Int).SetString("7995790102247196352256", 10)
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
