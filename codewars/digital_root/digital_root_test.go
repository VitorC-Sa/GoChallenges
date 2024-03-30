package digital_root

import (
	"testing"

	"github.com/VitorC-Sa/GoChallenges/utils/test"
)

func TestDigitalRoot(t *testing.T) {
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
				n: 16,
			},
			want: 7,
		},
		{
			name: tName(),
			args: args{
				n: 195,
			},
			want: 6,
		},
		{
			name: tName(),
			args: args{
				n: 992,
			},
			want: 2,
		},
		{
			name: tName(),
			args: args{
				n: 167346,
			},
			want: 9,
		},
		{
			name: tName(),
			args: args{
				n: 493193,
			},
			want: 2,
		},
		{
			name: tName(),
			args: args{
				n: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitalRoot(tt.args.n); got != tt.want {
				t.Errorf("DigitalRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
