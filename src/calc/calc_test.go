package rpn

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCalc(t *testing.T) {
	res := Calc("+", 4, 6, 7)
	if res != 17 {
		t.Errorf("Error expected %f, but found %f", float64(17), res)
	}

	res = Calc("-", 4, 6, 7)
	if res != -9 {
		t.Errorf("Error expected %f, but found %f", float64(-9), res)
	}

	res = Calc("*", 4, 6, 7)
	if res != 168 {
		t.Errorf("Error expected %f, but found %f", float64(168), res)
	}

	res = Calc("/", 6, 6, 4)
	if res != 0.25 {
		t.Errorf("Error expected %f, but found %f", 0.25, res)
	}
}
