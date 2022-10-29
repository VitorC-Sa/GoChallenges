package test

import "fmt"

func Name() func() string {
	i := 0
	return func() string {
		i++
		return fmt.Sprintf("Test N%d", i)
	}
}
