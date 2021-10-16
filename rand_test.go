// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"reflect"
	"testing"
)

func TestNewRand(t *testing.T) {
	type args struct {
		src *TntEngine
	}
	tests := []struct {
		name string
		args args
		want *Rand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRand(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Intn(t *testing.T) {
	type fields struct {
		tntMachine *TntEngine
		idx        int
		blk        CypherBlock
	}
	type args struct {
		max int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := &Rand{
				tntMachine: tt.fields.tntMachine,
				idx:        tt.fields.idx,
				blk:        tt.fields.blk,
			}
			if got := rnd.Intn(tt.args.max); got != tt.want {
				t.Errorf("Rand.Intn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Int63n(t *testing.T) {
	type fields struct {
		tntMachine *TntEngine
		idx        int
		blk        CypherBlock
	}
	type args struct {
		n int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := &Rand{
				tntMachine: tt.fields.tntMachine,
				idx:        tt.fields.idx,
				blk:        tt.fields.blk,
			}
			if got := rnd.Int63n(tt.args.n); got != tt.want {
				t.Errorf("Rand.Int63n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Perm(t *testing.T) {
	type fields struct {
		tntMachine *TntEngine
		idx        int
		blk        CypherBlock
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := &Rand{
				tntMachine: tt.fields.tntMachine,
				idx:        tt.fields.idx,
				blk:        tt.fields.blk,
			}
			if got := rnd.Perm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rand.Perm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Read(t *testing.T) {
	type fields struct {
		tntMachine *TntEngine
		idx        int
		blk        CypherBlock
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := &Rand{
				tntMachine: tt.fields.tntMachine,
				idx:        tt.fields.idx,
				blk:        tt.fields.blk,
			}
			gotN, err := rnd.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rand.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Rand.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
