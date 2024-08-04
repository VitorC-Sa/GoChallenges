package countonesbinary

import "testing"

func TestCountOnes(t *testing.T) {
	type args struct {
		left  uint64
		right uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "t1",
			args: args{
				left:  5,
				right: 7,
			},
			want: 7,
		},
		{
			name: "t2",
			args: args{
				left:  12,
				right: 29,
			},
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOnes(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("CountOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
