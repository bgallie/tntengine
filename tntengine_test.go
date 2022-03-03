// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"io"
	"math/big"
	"reflect"
	"testing"
)

func TestTntEngine_Left(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.BuildCipherMachine()
	tests := []struct {
		name string
		want chan CypherBlock
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
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestTntEngine_Right(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.BuildCipherMachine()
	tests := []struct {
		name string
		want chan CypherBlock
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
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestTntEngine_CounterKey(t *testing.T) {
	var tntMachine TntEngine
	tests := []struct {
		name             string
		key              string
		proFormaFileName string
		want             string
	}{
		{
			name:             "ttec1",
			key:              "SecretKey",
			proFormaFileName: "",
			want:             "e922e73a0f662987531e0950e7f8f11093f6b7a8bb043306b1feb723b19ef61b",
		},
		{
			name:             "ttec2",
			key:              "SecretKey",
			proFormaFileName: "test.proforma.json",
			want:             "0c9eba881bf288ccbbb4001229b7700fcbc90ec5ef08613946c4b629c111194d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tntMachine.Init([]byte(tt.key), tt.proFormaFileName)
			if got := tntMachine.CounterKey(); got != tt.want {
				t.Errorf("TntEngine.CounterKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTntEngine_Index(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"), "")
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
	tntMachine.Init([]byte("SecretKey"), "")
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
	tntMachine.Init([]byte("SecretKey"), "")
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
	tntMachine.Init([]byte("SecretKey"), "")
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
	tntMachine.Init([]byte("SecretKey"), "")
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
	tntMachine.Init([]byte("SecretKey"), "")
	tests := []struct {
		name string
		want *big.Int
	}{
		// TODO: Add test cases.
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
	tntMachine.Init([]byte("SecretKey"), "")
	type args struct {
		secret           []byte
		proFormaFileName string
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
			e.Init(tt.args.secret, tt.args.proFormaFileName)
		})
	}
}

func TestTntEngine_BuildCipherMachine(t *testing.T) {
	var tntMachine TntEngine
	tntMachine.Init([]byte("SecretKey"), "")
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
	type args struct {
		pfmReader io.Reader
	}
	tests := []struct {
		name string
		args args
		want *[]Crypter
	}{
		{
			name: "tcpfm1",
			args: args{pfmReader: nil},
			want: &[]Crypter{Rotor1, Rotor2, Permutator1, Rotor3, Rotor4, Permutator2, Rotor5, Rotor6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createProFormaMachine(tt.args.pfmReader); !reflect.DeepEqual(got, tt.want) {
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
