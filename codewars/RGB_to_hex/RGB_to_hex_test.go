package rgb_to_hex

import (
	"testing"

	"github.com/VitorC-Sa/GoChallenges/utils/test"
)

type Args struct {
	r int
	g int
	b int
}

type Test = struct {
	name string
	args Args
	want string
}

func getTestCase() []Test {
	tName := test.Name()
	return []Test{
		{name: tName(), args: Args{r: 0, g: 0, b: 0}, want: "000000"},
		{name: tName(), args: Args{r: 1, g: 2, b: 3}, want: "010203"},
		{name: tName(), args: Args{r: 255, g: 255, b: 255}, want: "FFFFFF"},
		{name: tName(), args: Args{r: 254, g: 253, b: 252}, want: "FEFDFC"},
		{name: tName(), args: Args{r: -20, g: 275, b: 125}, want: "00FF7D"},
	}
}

func TestMySolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySolution(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
				t.Errorf("MySolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnotherSolution(t *testing.T) {
	tests := getTestCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnotherSolution(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
				t.Errorf("AnotherSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
