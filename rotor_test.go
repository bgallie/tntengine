// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tntengine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewRotor(t *testing.T) {
	type args struct {
		size  int
		start int
		step  int
		rotor []byte
	}
	tests := []struct {
		name string
		args args
		want *Rotor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new(Rotor).New(tt.args.size, tt.args.start, tt.args.step, tt.args.rotor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRotor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_Update(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	type args struct {
		random *Rand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			r.Update(tt.args.random)
		})
	}
}

func TestRotor_sliceRotor(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			r.sliceRotor()
		})
	}
}

func TestRotor_SetIndex(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	type args struct {
		idx *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			r.SetIndex(tt.args.idx)
		})
	}
}

func TestRotor_Index(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   *big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			if got := r.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_ApplyF(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	type args struct {
		blk *[CypherBlockBytes]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[CypherBlockBytes]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			if got := r.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_ApplyG(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	type args struct {
		blk *[CypherBlockBytes]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[CypherBlockBytes]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			if got := r.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.ApplyG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_String(t *testing.T) {
	type fields struct {
		Size    int
		Start   int
		Step    int
		Current int
		Rotor   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				Size:    tt.fields.Size,
				Start:   tt.fields.Start,
				Step:    tt.fields.Step,
				Current: tt.fields.Current,
				Rotor:   tt.fields.Rotor,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Rotor.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
