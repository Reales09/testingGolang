package testUnitario

import "testing"

func TestSumar(t *testing.T) {
	total := sumar(5, 5)
	if total != 11 {
		t.Errorf("Suma incorrecta, se esperaba %d y se obtuvo %d", 11, total)
	}
}
