package stringEnds

import (
	"testing"
)

func Test_mySolution(t *testing.T) {
	type args struct {
		str    string
		ending string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "t1", args: args{str: "", ending: ""}, want: true},
		{name: "t2", args: args{str: " ", ending: ""}, want: true},
		{name: "t3", args: args{str: "abc", ending: "c"}, want: true},
		{name: "t4", args: args{str: "banana", ending: "ana"}, want: true},
		{name: "t5", args: args{str: "a", ending: "z"}, want: false},
		{name: "t6", args: args{str: "", ending: "t"}, want: false},
		{name: "t7", args: args{str: "$a = $b + 1", ending: "+1"}, want: false},
		{name: "t8", args: args{str: "    ", ending: "   "}, want: true},
		{name: "t9", args: args{str: "1oo", ending: "100"}, want: false},
		{name: "t9", args: args{str: "abc", ending: "zzabc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySolution(tt.args.str, tt.args.ending); got != tt.want {
				t.Errorf("mySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bestSolution(t *testing.T) {
	type args struct {
		str    string
		ending string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "t1", args: args{str: "", ending: ""}, want: true},
		{name: "t2", args: args{str: " ", ending: ""}, want: true},
		{name: "t3", args: args{str: "abc", ending: "c"}, want: true},
		{name: "t4", args: args{str: "banana", ending: "ana"}, want: true},
		{name: "t5", args: args{str: "a", ending: "z"}, want: false},
		{name: "t6", args: args{str: "", ending: "t"}, want: false},
		{name: "t7", args: args{str: "$a = $b + 1", ending: "+1"}, want: false},
		{name: "t8", args: args{str: "    ", ending: "   "}, want: true},
		{name: "t9", args: args{str: "1oo", ending: "100"}, want: false},
		{name: "t9", args: args{str: "abc", ending: "zzabc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestSolution(tt.args.str, tt.args.ending); got != tt.want {
				t.Errorf("bestSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
