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
				new(Rotor).New(1783, 201, 1696, []byte{
					162, 147, 81, 248, 75, 180, 188, 20, 149, 19, 88, 64, 129, 204, 166, 94,
					177, 188, 48, 70, 127, 234, 235, 114, 94, 158, 209, 76, 168, 213, 33, 126,
					226, 225, 199, 129, 114, 211, 99, 163, 7, 211, 115, 227, 19, 27, 83, 104,
					239, 147, 24, 50, 60, 186, 116, 188, 132, 147, 170, 201, 220, 129, 26, 71,
					243, 139, 78, 231, 148, 125, 208, 190, 185, 254, 186, 105, 70, 229, 125, 122,
					227, 158, 165, 34, 136, 213, 17, 39, 197, 239, 214, 178, 245, 251, 174, 116,
					215, 102, 141, 122, 199, 44, 161, 15, 139, 9, 84, 131, 222, 248, 236, 196,
					133, 237, 156, 191, 4, 164, 56, 106, 222, 237, 210, 242, 173, 20, 245, 93,
					98, 52, 114, 108, 19, 113, 243, 81, 156, 166, 189, 174, 243, 119, 79, 41,
					225, 64, 121, 8, 209, 139, 131, 140, 84, 199, 137, 161, 174, 121, 25, 176,
					161, 137, 64, 152, 46, 178, 24, 69, 233, 39, 186, 167, 190, 86, 64, 228,
					147, 233, 97, 205, 214, 231, 54, 63, 166, 109, 151, 67, 246, 215, 227, 200,
					132, 243, 17, 168, 102, 173, 209, 55, 53, 22, 240, 32, 138, 218, 147, 184,
					178, 210, 150, 41, 133, 11, 135, 203, 185, 107, 58, 162, 221, 119, 56, 209,
					201, 40, 252, 37, 90, 94, 138, 202, 9, 44, 160, 64, 102, 83, 175, 88,
					94, 24, 163, 63, 245, 117, 57, 47, 207, 104, 38, 212, 234, 16, 191, 143}),
				new(Rotor).New(1777, 608, 1259, []byte{
					149, 98, 118, 183, 12, 47, 237, 205, 33, 65, 254, 108, 62, 240, 199, 165,
					196, 94, 222, 204, 229, 249, 58, 84, 142, 210, 144, 57, 191, 110, 42, 56,
					178, 114, 35, 195, 179, 226, 19, 86, 213, 76, 225, 122, 26, 47, 206, 137,
					151, 121, 167, 136, 154, 244, 95, 200, 50, 100, 253, 37, 48, 50, 38, 135,
					200, 196, 141, 123, 9, 165, 203, 17, 187, 148, 27, 188, 251, 30, 165, 34,
					154, 218, 194, 214, 27, 35, 231, 128, 124, 49, 183, 80, 69, 202, 66, 69,
					179, 224, 53, 205, 167, 196, 1, 148, 63, 252, 54, 209, 191, 39, 1, 31,
					201, 68, 27, 23, 96, 157, 102, 245, 127, 4, 8, 20, 222, 53, 83, 209,
					186, 104, 171, 251, 71, 67, 42, 27, 140, 250, 220, 12, 78, 43, 102, 73,
					0, 101, 62, 198, 172, 108, 145, 214, 230, 91, 232, 85, 249, 18, 233, 89,
					127, 99, 54, 158, 144, 28, 210, 77, 197, 123, 60, 206, 218, 171, 94, 3,
					217, 107, 239, 190, 130, 86, 240, 26, 82, 208, 123, 30, 181, 181, 15, 76,
					220, 3, 158, 78, 243, 208, 223, 228, 94, 1, 157, 159, 143, 137, 89, 191,
					105, 192, 222, 5, 175, 163, 186, 45, 184, 241, 90, 98, 211, 12, 43, 197,
					236, 110, 25, 94, 218, 155, 67, 130, 252, 217, 124, 224, 143, 75, 137, 189,
					188, 153, 203, 243, 117, 168, 28, 165, 33, 115, 126, 221, 84, 112, 42, 170}),
				new(Permutator).New(256, []byte{
					64, 236, 3, 176, 114, 180, 182, 34, 78, 224, 89, 73, 211, 21, 133, 84,
					106, 96, 76, 63, 190, 193, 218, 58, 155, 113, 115, 132, 69, 141, 179, 30,
					120, 202, 217, 213, 75, 52, 127, 229, 37, 181, 62, 123, 85, 170, 13, 57,
					196, 101, 31, 246, 215, 212, 135, 168, 239, 60, 91, 1, 252, 199, 152, 110,
					172, 28, 118, 208, 67, 187, 185, 223, 20, 232, 197, 66, 138, 157, 191, 192,
					15, 188, 240, 124, 137, 53, 108, 207, 216, 153, 83, 59, 251, 220, 95, 74,
					70, 165, 40, 4, 158, 175, 247, 107, 166, 82, 139, 198, 55, 248, 146, 26,
					116, 100, 51, 241, 25, 225, 245, 203, 231, 46, 235, 129, 72, 145, 167, 32,
					149, 2, 0, 184, 112, 111, 250, 178, 39, 99, 222, 228, 92, 49, 253, 206,
					219, 186, 144, 134, 6, 143, 8, 68, 131, 14, 27, 23, 177, 35, 33, 122,
					201, 226, 109, 255, 161, 18, 234, 130, 117, 171, 204, 243, 209, 29, 164, 254,
					97, 38, 88, 244, 121, 242, 230, 104, 156, 233, 221, 169, 128, 42, 173, 200,
					105, 210, 98, 22, 90, 125, 50, 119, 147, 148, 44, 163, 162, 12, 43, 56,
					126, 9, 24, 65, 159, 81, 154, 17, 5, 47, 41, 61, 71, 80, 36, 249,
					189, 7, 205, 183, 136, 238, 19, 54, 87, 45, 102, 16, 11, 194, 103, 93,
					86, 77, 140, 10, 150, 79, 151, 237, 142, 174, 48, 227, 160, 214, 94, 195}),
				new(Rotor).New(1759, 1037, 1014, []byte{
					191, 73, 197, 240, 99, 142, 62, 186, 195, 80, 248, 230, 114, 36, 76, 129,
					196, 204, 134, 17, 129, 231, 156, 77, 138, 220, 150, 218, 223, 128, 130, 248,
					170, 108, 27, 103, 28, 171, 53, 197, 202, 188, 27, 143, 203, 90, 73, 234,
					136, 137, 61, 244, 69, 116, 154, 44, 79, 72, 142, 155, 21, 99, 144, 231,
					72, 159, 245, 132, 31, 80, 130, 111, 106, 96, 211, 20, 129, 120, 36, 66,
					183, 252, 121, 94, 167, 100, 33, 196, 176, 172, 202, 153, 135, 176, 169, 30,
					45, 229, 138, 230, 197, 6, 189, 90, 87, 27, 19, 228, 92, 161, 130, 18,
					65, 33, 215, 154, 12, 227, 26, 172, 11, 33, 175, 31, 242, 185, 56, 161,
					148, 50, 227, 39, 26, 11, 111, 79, 33, 222, 198, 84, 171, 224, 155, 2,
					1, 161, 158, 212, 139, 204, 90, 255, 231, 26, 66, 189, 117, 115, 150, 167,
					234, 184, 144, 244, 113, 242, 188, 97, 0, 195, 82, 253, 5, 238, 205, 97,
					0, 247, 96, 134, 121, 175, 243, 199, 74, 75, 156, 88, 77, 84, 255, 90,
					62, 219, 190, 61, 108, 49, 23, 21, 14, 150, 184, 236, 222, 233, 246, 152,
					245, 241, 195, 122, 219, 135, 188, 126, 173, 100, 6, 243, 223, 164, 98, 248,
					49, 71, 31, 221, 97, 40, 124, 115, 57, 18, 166, 64, 98, 102, 195, 136,
					192, 115, 206, 38, 69, 110, 75, 237, 111, 64, 65, 252, 12, 225, 92, 189}),
				new(Rotor).New(1753, 862, 893, []byte{
					30, 47, 74, 29, 212, 73, 198, 8, 231, 188, 102, 164, 252, 30, 122, 154,
					161, 201, 3, 16, 250, 176, 133, 103, 130, 154, 164, 85, 65, 197, 109, 121,
					17, 151, 147, 22, 70, 167, 65, 109, 233, 100, 125, 163, 252, 44, 245, 153,
					94, 29, 182, 94, 84, 23, 22, 107, 189, 30, 107, 137, 222, 255, 79, 245,
					125, 155, 109, 199, 24, 151, 133, 18, 190, 160, 139, 247, 219, 6, 107, 136,
					56, 4, 210, 13, 132, 129, 52, 2, 11, 246, 231, 44, 104, 84, 19, 188,
					247, 125, 172, 162, 124, 27, 184, 151, 98, 115, 100, 189, 85, 81, 165, 244,
					205, 134, 27, 78, 226, 102, 184, 120, 58, 68, 32, 175, 228, 153, 98, 32,
					241, 99, 216, 96, 129, 163, 124, 158, 108, 117, 229, 58, 247, 124, 35, 67,
					90, 60, 146, 228, 205, 64, 11, 47, 221, 156, 218, 202, 218, 243, 162, 153,
					222, 208, 255, 49, 55, 64, 181, 254, 24, 45, 130, 155, 14, 119, 225, 131,
					216, 227, 157, 135, 247, 195, 86, 190, 168, 123, 71, 78, 215, 239, 42, 214,
					55, 142, 240, 41, 30, 11, 75, 14, 192, 5, 54, 241, 186, 127, 247, 16,
					121, 44, 138, 85, 235, 255, 7, 135, 229, 58, 16, 60, 94, 148, 58, 168,
					147, 140, 17, 206, 121, 205, 72, 249, 61, 244, 52, 67, 147, 7, 32, 244,
					97, 11, 207, 4, 53, 73, 171, 130, 138, 219, 242, 200, 187, 190, 172, 185}),
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

func TestTntEngine_EngineType(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			if got := e.EngineType(); got != tt.want {
				t.Errorf("TntEngine.EngineType() = %v, want %v", got, tt.want)
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

func TestTntEngine_Init(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	type args struct {
		secret []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			e.Init(tt.args.secret)
		})
	}
}

func TestTntEngine_BuildCipherMachine(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			e.BuildCipherMachine()
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
