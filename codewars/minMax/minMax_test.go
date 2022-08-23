package minmax

import (
	"reflect"
	"testing"
)

type args struct {
	arr []int
}

type tests struct {
	name string
	args args
	want [2]int
}

func getTestCase() []tests {
	return []tests{
		{name: "t0", args: args{arr: []int{1}}, want: [2]int{1, 1}},
		{name: "t1", args: args{arr: []int{1, 2, 3, 4, 5}}, want: [2]int{1, 5}},
		{name: "t2", args: args{arr: []int{2334454, 5}}, want: [2]int{5, 2334454}},
		{name: "t3", args: args{arr: []int{2, 4, 5, 1, 3}}, want: [2]int{1, 5}},
		{name: "t4", args: args{arr: []int{4, 4}}, want: [2]int{4, 4}},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
