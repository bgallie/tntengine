// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"reflect"
	"testing"
)

func closeTntMachine(e *TntEngine) {
	blk := new(CypherBlock)
	e.Left() <- *blk
	<-e.Right()
}
func TestNewRand(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	type args struct {
		src *TntEngine
	}
	tests := []struct {
		name  string
		args  args
		want  *Rand
		wantK string
	}{
		{
			name:  "NewRandTest 1",
			args:  args{tntMachine},
			want:  &Rand{tntMachine, CypherBlockBytes, emptyBlk},
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			if got := NewRand(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	type args struct {
		src *TntEngine
	}
	tests := []struct {
		name  string
		args  args
		want  *Rand
		wantK string
	}{
		{
			name:  "NewTest 1",
			args:  args{tntMachine},
			want:  &Rand{tntMachine, CypherBlockBytes, emptyBlk},
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			if got := new(Rand).New(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Intn(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		max int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		wantK string
		wantR *Rand
	}{
		{
			name:  "Intn Test 1",
			args:  args{1000},
			want:  28,
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
			wantR: &Rand{tntMachine, CypherBlockBytes, emptyBlk},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			if !reflect.DeepEqual(rnd, tt.wantR) {
				t.Errorf("NewRand() = %v, want %v", rnd, tt.wantR)
			}
			if got := rnd.Intn(tt.args.max); got != tt.want {
				t.Errorf("Rand.Intn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Int63n(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		n int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		wantK string
		wantR *Rand
	}{
		{
			name:  "Int63n Test 1",
			args:  args{1000000000},
			want:  1874693,
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
			wantR: &Rand{tntMachine, CypherBlockBytes, emptyBlk},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			if !reflect.DeepEqual(rnd, tt.wantR) {
				t.Errorf("New() = %v, want %v", rnd, tt.wantR)
			}
			if got := rnd.Int63n(tt.args.n); got != tt.want {
				t.Errorf("Rand.Int63n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Perm(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		wantK string
	}{
		{
			name:  "Prem Test 1",
			args:  args{10},
			want:  []int{1, 8, 9, 6, 7, 2, 5, 3, 4, 0},
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			if got := rnd.Perm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rand.Perm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Read(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantN   int
		wantErr bool
		wantK   string
	}{
		{
			name: "Read test 1",
			args: args{make([]byte, 36)},
			want: []byte{
				192, 28, 155, 5, 242, 35, 70, 248, 5, 254, 131, 110, 228, 171,
				200, 160, 60, 249, 157, 12, 222, 88, 255, 203, 243, 121, 103,
				241, 252, 81, 136, 5, 5, 218, 72, 61},
			wantN:   36,
			wantErr: false,
			wantK:   "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			gotN, err := rnd.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rand.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Rand.Read() count = %v, want %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(tt.args.p, tt.want) {
				t.Errorf("Rand.Read() = %v, want %v", tt.args.p, tt.want)
			}
		})
	}
}
