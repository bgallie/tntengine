// Package cryptors - bit manipulation routines tests
package tntengine

import (
	"reflect"
	"testing"
)

func TestSetBit(t *testing.T) {
	type args struct {
		ary []byte
		bit uint
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "Set bit 0", args: args{ary: []byte{0x00, 0x00}, bit: 0}, want: []byte{0x01, 0x00}},
		{name: "Set bit 4", args: args{ary: []byte{0x00, 0x00}, bit: 4}, want: []byte{0x10, 0x00}},
		{name: "Set bit 8", args: args{ary: []byte{0x00, 0x00}, bit: 8}, want: []byte{0x00, 0x01}},
		{name: "Set bit 12", args: args{ary: []byte{0x00, 0x00}, bit: 12}, want: []byte{0x00, 0x10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBit(tt.args.ary, tt.args.bit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClrBit(t *testing.T) {
	type args struct {
		ary []byte
		bit uint
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "Clear bit 0", args: args{ary: []byte{0xFF, 0xFF}, bit: 0}, want: []byte{0xFE, 0xFF}},
		{name: "Clear bit 4", args: args{ary: []byte{0xFF, 0xFF}, bit: 4}, want: []byte{0xEF, 0xFF}},
		{name: "Clear bit 8", args: args{ary: []byte{0xFF, 0xFF}, bit: 8}, want: []byte{0xFF, 0xFE}},
		{name: "Clear bit 12", args: args{ary: []byte{0xFF, 0xFF}, bit: 12}, want: []byte{0xFF, 0xEF}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClrBit(tt.args.ary, tt.args.bit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClrBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBit(t *testing.T) {
	type args struct {
		ary []byte
		bit uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Get bit 0", args: args{ary: []byte{0x55, 0x55}, bit: 0}, want: true},
		{name: "Get bit 1", args: args{ary: []byte{0x55, 0x55}, bit: 1}, want: false},
		{name: "Get bit 2", args: args{ary: []byte{0x55, 0x55}, bit: 2}, want: true},
		{name: "Get bit 3", args: args{ary: []byte{0x55, 0x55}, bit: 3}, want: false},
		{name: "Get bit 4", args: args{ary: []byte{0x55, 0x55}, bit: 4}, want: true},
		{name: "Get bit 5", args: args{ary: []byte{0x55, 0x55}, bit: 5}, want: false},
		{name: "Get bit 6", args: args{ary: []byte{0x55, 0x55}, bit: 6}, want: true},
		{name: "Get bit 7", args: args{ary: []byte{0x55, 0x55}, bit: 7}, want: false},
		{name: "Get bit 8", args: args{ary: []byte{0x55, 0x55}, bit: 8}, want: true},
		{name: "Get bit 9", args: args{ary: []byte{0x55, 0x55}, bit: 9}, want: false},
		{name: "Get bit 10", args: args{ary: []byte{0x55, 0x55}, bit: 10}, want: true},
		{name: "Get bit 11", args: args{ary: []byte{0x55, 0x55}, bit: 11}, want: false},
		{name: "Get bit 12", args: args{ary: []byte{0x55, 0x55}, bit: 12}, want: true},
		{name: "Get bit 13", args: args{ary: []byte{0x55, 0x55}, bit: 13}, want: false},
		{name: "Get bit 14", args: args{ary: []byte{0x55, 0x55}, bit: 14}, want: true},
		{name: "Get bit 16", args: args{ary: []byte{0x55, 0x55}, bit: 15}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBit(tt.args.ary, tt.args.bit); got != tt.want {
				t.Errorf("GetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
