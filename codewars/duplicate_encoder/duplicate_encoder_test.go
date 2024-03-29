package duplicate_encoder

import "testing"

type args struct {
	word string
}

type tests struct {
	name string
	args args
	want string
}

func getTestCase() []tests {
	return []tests{
		{name: "t0", args: args{word: "din"}, want: "((("},
		{name: "t1", args: args{word: "recede"}, want: "()()()"},
		{name: "t2", args: args{word: "(( @"}, want: "))(("},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.word); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
