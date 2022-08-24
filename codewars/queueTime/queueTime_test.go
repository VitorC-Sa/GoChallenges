package queuetime

import (
	"testing"
)

type args struct {
	customers []int
	n         int
}

type tests struct {
	name string
	args args
	want int
}

func getTestCase() []tests {
	return []tests{
		{name: "t0", args: args{customers: []int{}, n: 1}, want: 0},
		{name: "t1", args: args{customers: []int{1, 2, 3, 4}, n: 1}, want: 10},
		{name: "t2", args: args{customers: []int{2, 2, 3, 3, 4, 4}, n: 2}, want: 9},
		{name: "t3", args: args{customers: []int{1, 2, 3, 4, 5}, n: 100}, want: 5},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.customers, tt.args.n); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
