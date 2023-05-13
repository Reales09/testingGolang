package testUnitario

import "testing"

// func TestSumar(t *testing.T) {
// 	total := sumar(5, 5)
// 	if total != 11 {
// 		t.Errorf("Suma incorrecta, se esperaba %d y se obtuvo %d", 11, total)
// 	}
// }

func TestSumar(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{1, 1, 2},
		{25, 25, 50},
	}

	for _, test := range tabla {
		if total := sumar(test.a, test.b); total != test.c {
			t.Errorf("Suma incorrecta, se esperaba %d y se obtuvo %d", test.c, total)
		}
	}
}

func TestGetMax(t *testing.T) {

	tabla := []struct {
		a int
		b int

		c int
	}{
		{4, 2, 4},
		{5, 3, 5},
		{10, 10, 10},
	}
	for _, item := range tabla {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("Maximo incorrecto, se esperaba %d y se obtuvo %d", item.c, max)
		}
	}
}

func TestFibonnaci(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
		{20, 6765},
	}

	for _, item := range tabla {
		if r := Fibonnaci(item.n); r != item.r {
			t.Errorf("Fibonnaci incorrecto, se esperaba %d y se obtuvo %d", item.r, r)
		}
	}
}
