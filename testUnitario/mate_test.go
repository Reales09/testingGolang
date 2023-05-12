package testUnitario

import "testing"

// func TestSumar(t *testing.T) {
// 	total := sumar(5, 5)
// 	if total != 11 {
// 		t.Errorf("Suma incorrecta, se esperaba %d y se obtuvo %d", 11, total)
// 	}
// }

func TestSumar(t *testing.T){
	tabla := []struct {
		a int
		b int
		c int
	}
	{
		{1, 2, 3},
		{1, 1, 2},
		{25, 25,50},
	}

	for _, test := range tabla {
		if total := sumar(test.a, test.b); total != test.c {
			t.Errorf("Suma incorrecta, se esperaba %d y se obtuvo %d", test.c, total)
		}
	}
}
