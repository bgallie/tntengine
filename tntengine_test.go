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
	tests := []struct {
		name string
		want []Crypter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tntMachine
			if got := e.Engine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TntEngine.Engine() = %v, want %v", got, tt.want)
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

func Test_updateRotor(t *testing.T) {
	type args struct {
		r      *Rotor
		random *Rand
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r.Update(tt.args.random)
		})
	}
}

func Test_updatePermutator(t *testing.T) {
	type args struct {
		p      *Permutator
		random *Rand
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.p.Update(tt.args.random)
		})
	}
}
