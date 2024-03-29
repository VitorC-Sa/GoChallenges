package rot_13_decipher

import (
	"testing"

	"github.com/VitorC-Sa/GoChallenges/utils/test"
)

func TestMySolution(t *testing.T) {
	tName := test.Name()

	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: tName(), args: args{msg: "EBG13 rknzcyr."}, want: "ROT13 example."},
		{name: tName(), args: args{msg: "How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf."}, want: "Ubj pna lbh gryy na rkgebireg sebz na\nvagebireg ng AFN? In the elevators,\nthe extrovert looks at the OTHER guy's shoes."},
		{name: tName(), args: args{msg: "123"}, want: "123"},
		{name: tName(), args: args{msg: "Guvf vf npghnyyl gur svefg xngn V rire znqr. Gunaxf sbe svavfuvat vg! :)"}, want: "This is actually the first kata I ever made. Thanks for finishing it! :)"},
		{name: tName(), args: args{msg: "@[`{"}, want: "@[`{"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.msg); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
