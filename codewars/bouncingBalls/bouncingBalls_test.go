package bouncingballs

import "testing"

type args struct {
	h      float64
	bounce float64
	window float64
}

type tests struct {
	name string
	args args
	want int
}

func getTestCase() []tests {
	return []tests{
		{name: "t0", args: args{h: 3, bounce: 0.66, window: 1.5}, want: 3},
		{name: "t1", args: args{h: 40, bounce: 0.4, window: 10}, want: 3},
		{name: "t2", args: args{h: 10, bounce: 0.6, window: 10}, want: -1},
		{name: "t3", args: args{h: 40, bounce: 1, window: 10}, want: -1},
		{name: "t4", args: args{h: 5, bounce: -1, window: 1.5}, want: -1},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.h, tt.args.bounce, tt.args.window); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
