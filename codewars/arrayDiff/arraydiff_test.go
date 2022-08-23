package arraydiff

import (
	"reflect"
	"testing"
)

type args struct {
	a []int
	b []int
}

type tests struct {
	name string
	args args
	want []int
}

func getTestCase() []tests {
	return []tests{
		{name: "t1", args: args{a: []int{1, 2}, b: []int{1}}, want: []int{2}},
		{name: "t2", args: args{a: []int{1, 2, 2}, b: []int{1}}, want: []int{2, 2}},
		{name: "t3", args: args{a: []int{1, 2, 2}, b: []int{2}}, want: []int{1}},
		{name: "t4", args: args{a: []int{1, 2, 2}, b: []int{}}, want: []int{1, 2, 2}},
		{name: "t5", args: args{a: []int{}, b: []int{1, 2}}, want: []int{}},
		{name: "t6", args: args{a: []int{1, 2, 3}, b: []int{1, 2}}, want: []int{3}},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
