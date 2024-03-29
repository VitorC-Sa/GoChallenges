package primewithevendigits

import (
	"testing"

	"github.com/VitorC-Sa/GoChallenges/utils/test"
)

func TestF(t *testing.T) {
	tName := test.Name()
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: tName(),
			args: args{
				n: 1000,
			},
			want: 887,
		},
		{
			name: tName(),
			args: args{
				n: 1210,
			},
			want: 1201,
		},
		{
			name: tName(),
			args: args{
				n: 10000,
			},
			want: 8887,
		},
		{
			name: tName(),
			args: args{
				n: 500,
			},
			want: 487,
		},
		{
			name: tName(),
			args: args{
				n: 487,
			},
			want: 467,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := F(tt.args.n); got != tt.want {
				t.Errorf("F() = %v, want %v", got, tt.want)
			}
		})
	}
}
