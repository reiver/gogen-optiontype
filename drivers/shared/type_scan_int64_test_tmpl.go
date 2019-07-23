package shared

const (
	TypeScanInt64TestTmpl =
`
import (
	"fmt"
	"math"
	"time"

	"testing"
)

func TestTypeScan(t *testing.T) {
	tests := []struct {
		Datum    interface{}
		Expected Type
	}{
		{
			Datum:    Something(math.MinInt64),
			Expected: Something(math.MinInt64),
		},
		{
			Datum:    Something(-5),
			Expected: Something(-5),
		},
		{
			Datum:    Something(-4),
			Expected: Something(-4),
		},
		{
			Datum:    Something(-3),
			Expected: Something(-3),
		},
		{
			Datum:    Something(-2),
			Expected: Something(-2),
		},
		{
			Datum:    Something(-1),
			Expected: Something(-1),
		},
		{
			Datum:    Something(0),
			Expected: Something(0),
		},
		{
			Datum:    Something(1),
			Expected: Something(1),
		},
		{
			Datum:    Something(2),
			Expected: Something(2),
		},
		{
			Datum:    Something(3),
			Expected: Something(3),
		},
		{
			Datum:    Something(4),
			Expected: Something(4),
		},
		{
			Datum:    Something(5),
			Expected: Something(5),
		},
		{
			Datum:    Something(math.MaxInt64),
			Expected: Something(math.MaxInt64),
		},



		{
			Datum:         int64(math.MinInt64),
			Expected:  Something(math.MinInt64),
		},
		{
			Datum:         int64(-5),
			Expected:  Something(-5),
		},
		{
			Datum:         int64(-4),
			Expected:  Something(-4),
		},
		{
			Datum:         int64(-3),
			Expected:  Something(-3),
		},
		{
			Datum:         int64(-2),
			Expected:  Something(-2),
		},
		{
			Datum:         int64(-1),
			Expected:  Something(-1),
		},
		{
			Datum:         int64(0),
			Expected:  Something(0),
		},
		{
			Datum:         int64(1),
			Expected:  Something(1),
		},
		{
			Datum:         int64(2),
			Expected:  Something(2),
		},
		{
			Datum:         int64(3),
			Expected:  Something(3),
		},
		{
			Datum:         int64(4),
			Expected:  Something(4),
		},
		{
			Datum:         int64(5),
			Expected:  Something(5),
		},
		{
			Datum:         int64(math.MaxInt64),
			Expected:  Something(math.MaxInt64),
		},



		{
			Datum: fmt.Sprintf("%d", math.MinInt64),
			Expected:      Something(math.MinInt64),
		},
		{
			Datum: fmt.Sprintf("%d", -5),
			Expected:      Something(-5),
		},
		{
			Datum: fmt.Sprintf("%d", -4),
			Expected:      Something(-4),
		},
		{
			Datum: fmt.Sprintf("%d", -3),
			Expected:      Something(-3),
		},
		{
			Datum: fmt.Sprintf("%d", -2),
			Expected:      Something(-2),
		},
		{
			Datum: fmt.Sprintf("%d", -1),
			Expected:      Something(-1),
		},
		{
			Datum: fmt.Sprintf("%d", 0),
			Expected:      Something(0),
		},
		{
			Datum: fmt.Sprintf("%d", 1),
			Expected:      Something(1),
		},
		{
			Datum: fmt.Sprintf("%d", 2),
			Expected:      Something(2),
		},
		{
			Datum: fmt.Sprintf("%d", 3),
			Expected:      Something(3),
		},
		{
			Datum: fmt.Sprintf("%d", 4),
			Expected:      Something(4),
		},
		{
			Datum: fmt.Sprintf("%d", 5),
			Expected:      Something(5),
		},
		{
			Datum: fmt.Sprintf("%d", math.MaxInt64),
			Expected:      Something(math.MaxInt64),
		},



		{
			Datum: []byte(fmt.Sprintf("%d", math.MinInt64)),
			Expected:             Something(math.MinInt64),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -5)),
			Expected:             Something(-5),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -4)),
			Expected:             Something(-4),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -3)),
			Expected:             Something(-3),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -2)),
			Expected:             Something(-2),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -1)),
			Expected:             Something(-1),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 0)),
			Expected:             Something(0),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 1)),
			Expected:             Something(1),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 2)),
			Expected:             Something(2),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 3)),
			Expected:             Something(3),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 4)),
			Expected:             Something(4),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 5)),
			Expected:             Something(5),
		},
		{
			Datum: fmt.Sprintf("%d", math.MaxInt64),
			Expected:      Something(math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              x,
				Expected: Something(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              y,
				Expected: Something(y),
			}
			tests = append(tests, test)
		}

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              s,
				Expected: Something(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              s,
				Expected: Something(y),
			}
			tests = append(tests, test)
		}

		{
			s := fmt.Sprintf("%d", x)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              b,
				Expected: Something(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              b,
				Expected: Something(y),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		var datum Type

		if err := datum.Scan(test.Datum); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, datum; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeScanFail(t *testing.T) {
	tests := []struct {
		Datum interface{}
	}{
		{
			Datum: nil,
		},



		{
			Datum: false,
		},
		{
			Datum: true,
		},



		{
			Datum: float64(-1.1),
		},
		{
			Datum: float64(-1.0),
		},
		{
			Datum: float64(-0.1),
		},
		{
			Datum: float64(0.0),
		},
		{
			Datum: float64(0.1),
		},
		{
			Datum: float64(1.0),
		},
		{
			Datum: float64(1.1),
		},



		{
			Datum: time.Now(),
		},
	}


	for testNumber, test := range tests {
		var datum Type

		if err := datum.Scan(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
`
)
