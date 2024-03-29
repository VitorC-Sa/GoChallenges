package maximum_subarray_sum

import "testing"

func TestMySolution(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{numbers: []int{}}, want: 0},
		{name: "t2", args: args{numbers: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "t3", args: args{numbers: []int{-2, -1, -3, -4, -1, -2, -1, -5, -4}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.numbers); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
