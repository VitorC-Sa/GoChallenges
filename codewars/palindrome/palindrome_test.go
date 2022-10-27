package palindrome

import "testing"

func TestMySolution(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "t0", args: args{"Banana"}, want: false},
		{name: "t1", args: args{"osso"}, want: true},
		{name: "t2", args: args{"Osso"}, want: true},
		{name: "t3", args: args{"OSSO"}, want: true},
		{name: "t5", args: args{"ovo"}, want: true},
		{name: "t6", args: args{"çéâãô"}, want: false},
		{name: "t7", args: args{"A frase não é um palíndromo"}, want: false},
		{name: "t8", args: args{"Laço bacana para panaca boçal"}, want: true},
		{name: "t9", args: args{`Aí, Lima falou: "Olá, família!"`}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.text); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
