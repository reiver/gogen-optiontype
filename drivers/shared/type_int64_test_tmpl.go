package shared

const (
	TypeInt64TestTmpl =
`
import (
	"math"

	"testing"
)

func TestTypeInt64(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected int64
	}{
		{
			Datum: Something(math.MinInt64),
			Expected:        math.MinInt64,
		},
		{
			Datum: Something(-5),
			Expected:        -5,
		},
		{
			Datum: Something(-4),
			Expected:        -4,
		},
		{
			Datum: Something(-3),
			Expected:        -3,
		},
		{
			Datum: Something(-2),
			Expected:        -2,
		},
		{
			Datum: Something(-1),
			Expected:        -1,
		},
		{
			Datum: Something(0),
			Expected:        0,
		},
		{
			Datum: Something(1),
			Expected:        1,
		},
		{
			Datum: Something(2),
			Expected:        2,
		},
		{
			Datum: Something(3),
			Expected:        3,
		},
		{
			Datum: Something(4),
			Expected:        4,
		},
		{
			Datum: Something(5),
			Expected:        5,
		},
		{
			Datum: Something(math.MaxInt64),
			Expected:        math.MaxInt64,
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    Type
				Expected int64
			}{
				Datum: Something(x),
				Expected:        x,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    Type
				Expected int64
			}{
				Datum: Something(y),
				Expected:        y,
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		actual, err := datum.Int64()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeInt64Nothing(t *testing.T) {

	datum := Nothing()

	_, err := datum.Int64()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNothing, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}
`
)
