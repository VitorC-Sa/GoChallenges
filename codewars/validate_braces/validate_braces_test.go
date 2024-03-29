package validate_braces

import (
	"testing"
)

type args struct {
	str string
}

type tests struct {
	name string
	args args
	want bool
}

func getTestCase() []tests {
	return []tests{
		{name: "t0", args: args{str: "()"}, want: true},
		{name: "t1", args: args{str: "(){}[]"}, want: true},
		{name: "t2", args: args{str: "([{}])"}, want: true},
		{name: "t3", args: args{str: "(}"}, want: false},
		{name: "t4", args: args{str: "[(])"}, want: false},
		{name: "t5", args: args{str: "[({)](]"}, want: false},
		{name: "t6", args: args{str: "(("}, want: false},
		{name: "t7", args: args{str: "(((({{}})[]))"}, want: false},
		{name: "t8", args: args{str: "()({{}}"}, want: false},
		{name: "t9", args: args{str: "()({{}}))))))))"}, want: false},
		{name: "t10", args: args{str: "}}}}}}"}, want: false},
		{name: "t11", args: args{str: ")))))"}, want: false},
		{name: "t12", args: args{str: "]]]]]]"}, want: false},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.str); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyNewSolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyNewSolution(tt.args.str); got != tt.want {
				t.Errorf("MyNewSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestSolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestSolution(tt.args.str); got != tt.want {
				t.Errorf("BestSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
