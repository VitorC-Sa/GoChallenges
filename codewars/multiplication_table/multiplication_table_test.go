package multiplication_table

import (
	"reflect"
	"testing"

	"github.com/VitorC-Sa/GoChallenges/utils/test"
)

func TestMySolution(t *testing.T) {

	tName := test.Name()

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: tName(), args: args{n: 1}, want: [][]int{{1}}},
		{name: tName(), args: args{n: 2}, want: [][]int{{1, 2}, {2, 4}}},
		{name: tName(), args: args{n: 3}, want: [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
