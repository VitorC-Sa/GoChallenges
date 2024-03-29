package isbn_10_validation

import (
	"testing"

	"github.com/VitorC-Sa/GoChallenges/utils/test"
)

func TestMySolution(t *testing.T) {

	tName := test.Name()

	type args struct {
		isbn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: tName(), args: args{isbn: "1112223339"}, want: true},
		{name: tName(), args: args{isbn: "048665088X"}, want: true},
		{name: tName(), args: args{isbn: "1293000000"}, want: true},
		{name: tName(), args: args{isbn: "1234554321"}, want: true},
		{name: tName(), args: args{isbn: "1234512345"}, want: false},
		{name: tName(), args: args{isbn: "1293"}, want: false},
		{name: tName(), args: args{isbn: "X123456788"}, want: false},
		{name: tName(), args: args{isbn: "ABCDEFGHIJ"}, want: false},
		{name: tName(), args: args{isbn: "XXXXXXXXXX"}, want: false},
		{name: tName(), args: args{isbn: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.isbn); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
