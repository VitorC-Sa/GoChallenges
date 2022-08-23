package multiplesof

import (
	"testing"
)

type args struct {
	number int
}

type tests struct {
	name string
	args args
	want int
}

func getTestCase() []tests {
	return []tests{
		{name: "t0", args: args{number: -1}, want: 0},
		{name: "t1", args: args{number: 10}, want: 23},
		{name: "t2", args: args{number: 20}, want: 78},
		{name: "t2", args: args{number: 2}, want: 0},
		{name: "t2", args: args{number: 0}, want: 0},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.number); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestSolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestSolution(tt.args.number); got != tt.want {
				t.Errorf("BestSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
