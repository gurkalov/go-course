package calc

import (
	"course/calc/extend"
	"course/calc/simple"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSimpleBase(t *testing.T) {

	var res float64
	res = Compute(simple.Calc{2, 5}, "+")
	if res != 7 {
		t.Errorf("Error expected %f, but found %f", float64(7), res)
	}

	res = Compute(simple.Calc{2, 5}, "-")
	if res != -3 {
		t.Errorf("Error expected %f, but found %f", float64(-3), res)
	}

	res = Compute(simple.Calc{2, 5}, "*")
	if res != 10 {
		t.Errorf("Error expected %f, but found %f", float64(10), res)
	}

	res = Compute(simple.Calc{2, 5}, "/")
	if res != 0.4 {
		t.Errorf("Error expected %f, but found %f", float64(0.4), res)
	}
}

func TestExtendBase(t *testing.T) {

	var res float64
	res = Compute(extend.Calc{2, 5}, "+")
	if res != 7 {
		t.Errorf("Error expected %f, but found %f", float64(7), res)
	}

	res = Compute(simple.Calc{2, 5}, "-")
	if res != -3 {
		t.Errorf("Error expected %f, but found %f", float64(-3), res)
	}

	res = Compute(simple.Calc{2, 5}, "*")
	if res != 10 {
		t.Errorf("Error expected %f, but found %f", float64(10), res)
	}

	res = Compute(simple.Calc{2, 5}, "/")
	if res != 0.4 {
		t.Errorf("Error expected %f, but found %f", float64(0.4), res)
	}
}

func TestExtendExtend(t *testing.T) {

	var res float64
	res = ComputeExtend(extend.Calc{2, 5}, "^")
	if res != 32 {
		t.Errorf("Error expected %f, but found %f", float64(32), res)
	}

	res = ComputeExtend(extend.Calc{2, 5}, "mod")
	if res != 2 {
		t.Errorf("Error expected %f, but found %f", float64(2), res)
	}
}
