// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for the details.

package tntengine

import (
	"reflect"
	"testing"
)

func TestNewRand(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer tntMachine.CloseCipherMachine()
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
			want:  &Rand{tntMachine, CipherBlockBytes, emptyBlk},
			wantK: "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
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
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer tntMachine.CloseCipherMachine()
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
			want:  &Rand{tntMachine, CipherBlockBytes, emptyBlk},
			wantK: "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
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
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer tntMachine.CloseCipherMachine()
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
			want:  442,
			wantK: "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
			wantR: &Rand{tntMachine, CipherBlockBytes, emptyBlk},
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
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer tntMachine.CloseCipherMachine()
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
			want:  29018107,
			wantK: "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
			wantR: &Rand{tntMachine, CipherBlockBytes, emptyBlk},
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
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer tntMachine.CloseCipherMachine()
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
			want:  []int{7, 8, 0, 5, 6, 4, 9, 3, 2, 1},
			wantK: "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
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
	tntMachine.Init([]byte("SecretKey"))
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer tntMachine.CloseCipherMachine()
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
				1, 186, 199, 251, 33, 196, 225, 173, 88, 200, 222, 71,
				173, 172, 53, 253, 48, 12, 142, 120, 207, 22, 70, 152,
				79, 171, 7, 31, 35, 7, 136, 135, 216, 83, 104, 18},
			wantN:   36,
			wantErr: false,
			wantK:   "30a7c225e88daa83416dee1970dc58b81f0c3771d6eb801ce23b49439357cc16",
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
