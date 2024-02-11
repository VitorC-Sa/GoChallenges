package moveZerosToTheEnd

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "OK",
			args: args{
				arr: []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1},
			},
			want: []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0},
		},
		{
			name: "Only Zeros",
			args: args{
				arr: []int{0, 0, 0},
			},
			want: []int{0, 0, 0},
		},
		{
			name: "Zeros already in the End",
			args: args{
				arr: []int{1, 2, 0},
			},
			want: []int{1, 2, 0},
		},
		{
			name: "With Multiples Zeros in Sequence",
			args: args{
				arr: []int{5, 0, 6, 7, 8, 1, 1, 3, 1, 9, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []int{5, 6, 7, 8, 1, 1, 3, 1, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "With Negative Numbers",
			args: args{
				arr: []int{5, 0, 6, -1, 7, 8, 1, -2, 1, 3, -3, 1, 9, 0, -4, 0, 9, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []int{5, 6, -1, 7, 8, 1, -2, 1, 3, -3, 1, 9, -4, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveZeros(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
