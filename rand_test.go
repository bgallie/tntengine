// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"reflect"
	"testing"
)

func TestNewRand(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetIndex(BigZero)
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
			want:  &Rand{tntMachine, 32, emptyBlk},
			wantK: "a24d7bb490d928b8c9bf74b509c31d5d1d79b2039c38e3252be0a857ac619be2",
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
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestRand_Intn(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetIndex(BigZero)
	rnd := NewRand(tntMachine)
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
			want:  186,
			wantK: "a24d7bb490d928b8c9bf74b509c31d5d1d79b2039c38e3252be0a857ac619be2",
			wantR: &Rand{tntMachine, 32, emptyBlk},
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
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestRand_Int63n(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetIndex(BigZero)
	rnd := NewRand(tntMachine)
	type args struct {
		n int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		wantK string
	}{
		{
			name:  "Int63n Test 1",
			args:  args{1000000000},
			want:  79361042,
			wantK: "a24d7bb490d928b8c9bf74b509c31d5d1d79b2039c38e3252be0a857ac619be2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
			}
			if got := rnd.Int63n(tt.args.n); got != tt.want {
				t.Errorf("Rand.Int63n() = %v, want %v", got, tt.want)
			}
		})
	}
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestRand_Perm(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetIndex(BigZero)
	rnd := NewRand(tntMachine)
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
			want:  []int{3, 5, 0, 1, 6, 7, 8, 9, 2, 4},
			wantK: "a24d7bb490d928b8c9bf74b509c31d5d1d79b2039c38e3252be0a857ac619be2",
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
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}

func TestRand_Read(t *testing.T) {
	tntMachine := new(TntEngine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetIndex(BigZero)
	rnd := NewRand(tntMachine)
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
				4, 186, 244, 18, 244, 202, 69, 96, 213, 124, 18, 95, 8, 54, 220, 120,
				88, 183, 18, 103, 109, 21, 171, 51, 86, 129, 184, 228, 53, 195, 90, 185,
				23, 151, 67, 217},
			wantN:   36,
			wantErr: false,
			wantK:   "a24d7bb490d928b8c9bf74b509c31d5d1d79b2039c38e3252be0a857ac619be2",
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
	var blk CypherBlock
	tntMachine.left <- blk
	<-tntMachine.right
}
