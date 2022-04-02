package main // o teste deve estar no mesmo package do arquivo a ser testado

import "testing"

// tudo que possuir o prefixo Test será testado, é necessário também injetar a variavel testing.T
// para rodar basta utilizar go test -v | v de verbose (é necessário ter um módulo)
// para ver a cobertura  go test -cover
func TestSum(t *testing.T) {
	a := 5
	b := 6
	expected := 11
	res := sum(a, b)
	if res != expected {
		t.Errorf("Our sum function doesn't work, %d+%d isn't %d\n", a, b, res)
	}
}
