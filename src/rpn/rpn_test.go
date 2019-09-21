package rpn

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRpn(t *testing.T) {
	express := map[string]string{
		"1 + 2": "1 2 +",
		"1.5 + 3.2": "1.5 3.2 +",
		"3 + ( 4 * 2 )": "3 4 2 * +",
		"( 1 + 2 ) * 4 + 3": "1 2 + 4 * 3 +",
		"15 + 23": "15 23 +",
		"3 + 4 * 2 / ( 1 - 5 ) ^ 2": "3 4 2 * 1 5 - 2 ^ / +",
	}

	for k, v := range express {
		rpn := Rpn(k)
		if v != rpn {
			t.Errorf("Error expected %s, but found %s", v, rpn)
		}
	}
}

func TestCompute(t *testing.T) {
	express := map[string]float64{
		"1 + 2": 3,
		"1.5 + 3.2": 4.7,
		"3 + ( 4 * 2 )": 11,
		"( 1 + 2 ) * 4 + 3": 15,
		"15 + 23": 38,
		"3 + 4 * 2 / ( 1 - 5 ) ^ 2": 3.5,
	}

	for k, v := range express {
		res := Compute(k)
		if v != res {
			t.Errorf("Error expected %f, but found %f", v, res)
		}
	}
}
