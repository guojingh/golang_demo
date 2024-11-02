package base62

import (
	"testing"
)

func TestInt2String(t *testing.T) {
	tests := []struct {
		name string
		seq  uint64
		want string
	}{
		{name: "0", seq: 0, want: "0"},
		{name: "1", seq: 1, want: "1"},
		{name: "62", seq: 62, want: "10"},
		{name: "6347", seq: 6347, want: "1En"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.seq); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	tests := []struct {
		name    string
		S       string
		wantSeq uint64
	}{
		{name: "case 0", S: "0", wantSeq: 0},
		{name: "case 10", S: "10", wantSeq: 62},
		{name: "case 1En", S: "1En", wantSeq: 6347},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSeq := String2Int(tt.S); gotSeq != tt.wantSeq {
				t.Errorf("String2Int() = %v, want %v", gotSeq, tt.wantSeq)
			}
		})
	}
}
