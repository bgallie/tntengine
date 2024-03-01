// This is free and unencumbered software released into the public domain.

// See the UNLICENSE file for details.

package tntengine

import (
	"reflect"
	"testing"
)

func TestCypherBlock_String(t *testing.T) {
	tests := []struct {
		name string
		cblk CipherBlock
		want string
	}{
		{
			name: "tcbs1",
			cblk: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 21, 32},
			want: "CipherBlock:\t     Length: 32\n            \t   Capacity: 32\n            \t       Data:\t01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F 10\n            \t\t\t11 12 13 14 15 16 17 18 19 1A 1B 1C 1D 1E 15 20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cblk.String(); got != tt.want {
				t.Errorf("CypherBlock.String() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

// func TestCounter_SetIndex(t *testing.T) {
// 	var tntMachine TntEngine
// 	tntMachine.Init([]byte("SecretKey"))
// 	iCnt, _ := new(Counter).SetString("1234567890")
// 	type args struct {
// 		index *Counter
// 	}
// 	tests := []struct {
// 		name string
// 		cntr TntEngine
// 		args args
// 		want *Counter
// 	}{
// 		{
// 			name: "tcsi1",
// 			cntr: tntMachine,
// 			args: args{iCnt},
// 			want: iCnt,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.cntr.SetIndex(tt.args.index)
// 			if got := tt.cntr.Index(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Counter.Index() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// 	var blk CipherBlock
// 	tntMachine.left <- blk
// 	<-tntMachine.right
// 	tntMachine = *new(TntEngine)
// }

func TestCounter_Index(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	iCnt, _ := new(Counter).SetString("1234567890")
	tntMachine.SetIndex(iCnt)
	tests := []struct {
		name string
		cntr TntEngine
		want *Counter
	}{
		{
			name: "tci1",
			cntr: tntMachine,
			want: iCnt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cntr.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.Index() = %v, want %v", got, tt.want)
			}
		})
	}
	var blk CipherBlock
	tntMachine.left <- blk
	<-tntMachine.right
	tntMachine = *new(TntEngine)
}

func TestCounter_ApplyF(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetIndex(BigZero)
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name  string
		cntr  Crypter
		args  args
		want  CipherBlock
		want2 *Counter
	}{
		{
			name:  "tcaf1",
			cntr:  tntMachine.engine[len(tntMachine.engine)-1],
			args:  args{blk: *new(CipherBlock)},
			want:  *new(CipherBlock),
			want2: BigOne,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cntr.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.ApplyF() = %v, want %v", got, tt.want)
			}
			if got := tt.cntr.Index(); !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("Counter.Index() = %v, want %v", got, tt.want2)
			}
		})
	}
	var blk CipherBlock
	tntMachine.left <- blk
	<-tntMachine.right
	tntMachine = *new(TntEngine)
}

func TestCounter_ApplyG(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetIndex(BigZero)
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name  string
		cntr  Crypter
		args  args
		want  CipherBlock
		want2 *Counter
	}{
		{
			name:  "tcaf1",
			cntr:  tntMachine.engine[len(tntMachine.engine)-1],
			args:  args{blk: *new(CipherBlock)},
			want:  *new(CipherBlock),
			want2: BigZero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cntr.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.ApplyG() = %v, want %v", got, tt.want)
			}
			if got := tt.cntr.Index(); !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("Counter.Index() = %v, want %v", got, tt.want2)
			}
		})
	}
	var blk CipherBlock
	tntMachine.left <- blk
	<-tntMachine.right
	tntMachine = *new(TntEngine)
}

func TestSubBlock(t *testing.T) {
	type args struct {
		blk CipherBlock
		key CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tsb1",
			args: args{
				blk: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				key: []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []byte{
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			name: "tsb2",
			args: args{
				blk: []byte{0x59, 0xEF, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				key: []byte{0xC3, 0x0A, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []byte{0x96, 0xE4, 0x2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubBlock(tt.args.blk, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddBlock(t *testing.T) {
	type args struct {
		blk CipherBlock
		key CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tab1",
			args: args{
				blk: []byte{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
				key: []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "tab2",
			args: args{
				blk: []byte{0x96, 0xE4, 0x2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				key: []byte{0xC3, 0x0A, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []byte{0x59, 0xEF, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBlock(tt.args.blk, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptMachine(t *testing.T) {
	type args struct {
		ecm  Crypter
		left chan CipherBlock
	}
	tests := []struct {
		name string
		args args
		want chan CipherBlock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptMachine(tt.args.ecm, tt.args.left); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptMachine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecryptMachine(t *testing.T) {
	type args struct {
		ecm  Crypter
		left chan CipherBlock
	}
	tests := []struct {
		name string
		args args
		want chan CipherBlock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecryptMachine(tt.args.ecm, tt.args.left); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecryptMachine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createEncryptMachine(t *testing.T) {
	type args struct {
		ecms []Crypter
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  chan CipherBlock
		wantRight chan CipherBlock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := createEncryptMachine(tt.args.ecms...)
			if !reflect.DeepEqual(gotLeft, tt.wantLeft) {
				t.Errorf("createEncryptMachine() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if !reflect.DeepEqual(gotRight, tt.wantRight) {
				t.Errorf("createEncryptMachine() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}

func Test_createDecryptMachine(t *testing.T) {
	type args struct {
		ecms []Crypter
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  chan CipherBlock
		wantRight chan CipherBlock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := createDecryptMachine(tt.args.ecms...)
			if !reflect.DeepEqual(gotLeft, tt.wantLeft) {
				t.Errorf("createDecryptMachine() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if !reflect.DeepEqual(gotRight, tt.wantRight) {
				t.Errorf("createDecryptMachine() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}

func TestCounter_Add(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		cntr *Counter
		args args
		want *Counter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cntr.Add(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_Div(t *testing.T) {
	type args struct {
		divisor uint64
	}
	tests := []struct {
		name     string
		dividend *Counter
		args     args
		want     *Counter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dividend.Div(tt.args.divisor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_DivMod(t *testing.T) {
	type args struct {
		divisor uint64
	}
	tests := []struct {
		name     string
		dividend *Counter
		args     args
		want     *Counter
		want1    uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.dividend.DivMod(tt.args.divisor)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.DivMod() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Counter.DivMod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCounter_IsZero(t *testing.T) {
	tests := []struct {
		name  string
		index *Counter
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.index.IsZero(); got != tt.want {
				t.Errorf("Counter.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_Mod(t *testing.T) {
	type args struct {
		divisor uint64
	}
	tests := []struct {
		name     string
		dividend Counter
		args     args
		want     uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dividend.Mod(tt.args.divisor); got != tt.want {
				t.Errorf("Counter.Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_Mul(t *testing.T) {
	type args struct {
		multiplier uint64
	}
	tests := []struct {
		name         string
		multiplicand *Counter
		args         args
		want         *Counter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.multiplicand.Mul(tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_SetString(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name  string
		index *Counter
		args  args
		want  *Counter
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.index.SetString(tt.args.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.SetString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Counter.SetString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCounter_String(t *testing.T) {
	tests := []struct {
		name  string
		index Counter
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.index.String(); got != tt.want {
				t.Errorf("Counter.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_SetIndex(t *testing.T) {
	type args struct {
		val *Counter
	}
	tests := []struct {
		name  string
		index *Counter
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.index.SetIndex(tt.args.val)
		})
	}
}

func TestCounter_Sub(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		cntr *Counter
		args args
		want *Counter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cntr.Sub(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
