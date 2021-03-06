package shared

const (
	NullableTypeScanInt64TestTmpl =
`
import (
	"fmt"
	"math"
	"time"

	"testing"
)

func TestNullableTypeScan(t *testing.T) {
	tests := []struct {
		Datum    interface{}
		Expected NullableType
	}{
		{
			Datum:    nil,
			Expected: Null(),
		},



		{
			Datum:    Nothing(),
			Expected: NothingNullable(),
		},



		{
			Datum:            Something(math.MinInt64),
			Expected: SomethingNullable(math.MinInt64),
		},
		{
			Datum:            Something(-5),
			Expected: SomethingNullable(-5),
		},
		{
			Datum:            Something(-4),
			Expected: SomethingNullable(-4),
		},
		{
			Datum:            Something(-3),
			Expected: SomethingNullable(-3),
		},
		{
			Datum:            Something(-2),
			Expected: SomethingNullable(-2),
		},
		{
			Datum:            Something(-1),
			Expected: SomethingNullable(-1),
		},
		{
			Datum:            Something(0),
			Expected: SomethingNullable(0),
		},
		{
			Datum:            Something(1),
			Expected: SomethingNullable(1),
		},
		{
			Datum:            Something(2),
			Expected: SomethingNullable(2),
		},
		{
			Datum:            Something(3),
			Expected: SomethingNullable(3),
		},
		{
			Datum:            Something(4),
			Expected: SomethingNullable(4),
		},
		{
			Datum:            Something(5),
			Expected: SomethingNullable(5),
		},
		{
			Datum:            Something(math.MaxInt64),
			Expected: SomethingNullable(math.MaxInt64),
		},



		{
			Datum:    SomethingNullable(math.MinInt64),
			Expected: SomethingNullable(math.MinInt64),
		},
		{
			Datum:    SomethingNullable(-5),
			Expected: SomethingNullable(-5),
		},
		{
			Datum:    SomethingNullable(-4),
			Expected: SomethingNullable(-4),
		},
		{
			Datum:    SomethingNullable(-3),
			Expected: SomethingNullable(-3),
		},
		{
			Datum:    SomethingNullable(-2),
			Expected: SomethingNullable(-2),
		},
		{
			Datum:    SomethingNullable(-1),
			Expected: SomethingNullable(-1),
		},
		{
			Datum:    SomethingNullable(0),
			Expected: SomethingNullable(0),
		},
		{
			Datum:    SomethingNullable(1),
			Expected: SomethingNullable(1),
		},
		{
			Datum:    SomethingNullable(2),
			Expected: SomethingNullable(2),
		},
		{
			Datum:    SomethingNullable(3),
			Expected: SomethingNullable(3),
		},
		{
			Datum:    SomethingNullable(4),
			Expected: SomethingNullable(4),
		},
		{
			Datum:    SomethingNullable(5),
			Expected: SomethingNullable(5),
		},
		{
			Datum:    SomethingNullable(math.MaxInt64),
			Expected: SomethingNullable(math.MaxInt64),
		},



		{
			Datum:                int64(math.MinInt64),
			Expected: SomethingNullable(math.MinInt64),
		},
		{
			Datum:                int64(-5),
			Expected: SomethingNullable(-5),
		},
		{
			Datum:                int64(-4),
			Expected: SomethingNullable(-4),
		},
		{
			Datum:                int64(-3),
			Expected: SomethingNullable(-3),
		},
		{
			Datum:                int64(-2),
			Expected: SomethingNullable(-2),
		},
		{
			Datum:                int64(-1),
			Expected: SomethingNullable(-1),
		},
		{
			Datum:                int64(0),
			Expected: SomethingNullable(0),
		},
		{
			Datum:                int64(1),
			Expected: SomethingNullable(1),
		},
		{
			Datum:                int64(2),
			Expected: SomethingNullable(2),
		},
		{
			Datum:                int64(3),
			Expected: SomethingNullable(3),
		},
		{
			Datum:                int64(4),
			Expected: SomethingNullable(4),
		},
		{
			Datum:                int64(5),
			Expected: SomethingNullable(5),
		},
		{
			Datum:                int64(math.MaxInt64),
			Expected: SomethingNullable(math.MaxInt64),
		},



		{
			Datum:    fmt.Sprintf("%d", math.MinInt64),
			Expected: SomethingNullable(math.MinInt64),
		},
		{
			Datum:    fmt.Sprintf("%d", -5),
			Expected: SomethingNullable(-5),
		},
		{
			Datum:    fmt.Sprintf("%d", -4),
			Expected: SomethingNullable(-4),
		},
		{
			Datum:    fmt.Sprintf("%d", -3),
			Expected: SomethingNullable(-3),
		},
		{
			Datum:    fmt.Sprintf("%d", -2),
			Expected: SomethingNullable(-2),
		},
		{
			Datum:    fmt.Sprintf("%d", -1),
			Expected: SomethingNullable(-1),
		},
		{
			Datum:    fmt.Sprintf("%d", 0),
			Expected: SomethingNullable(0),
		},
		{
			Datum:    fmt.Sprintf("%d", 1),
			Expected: SomethingNullable(1),
		},
		{
			Datum:    fmt.Sprintf("%d", 2),
			Expected: SomethingNullable(2),
		},
		{
			Datum:    fmt.Sprintf("%d", 3),
			Expected: SomethingNullable(3),
		},
		{
			Datum:    fmt.Sprintf("%d", 4),
			Expected: SomethingNullable(4),
		},
		{
			Datum:    fmt.Sprintf("%d", 5),
			Expected: SomethingNullable(5),
		},
		{
			Datum:    fmt.Sprintf("%d", math.MaxInt64),
			Expected: SomethingNullable(math.MaxInt64),
		},



		{
			Datum: []byte(fmt.Sprintf("%d", math.MinInt64)),
			Expected:     SomethingNullable(math.MinInt64),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -5)),
			Expected:     SomethingNullable(-5),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -4)),
			Expected:     SomethingNullable(-4),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -3)),
			Expected:     SomethingNullable(-3),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -2)),
			Expected:     SomethingNullable(-2),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", -1)),
			Expected:     SomethingNullable(-1),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 0)),
			Expected:     SomethingNullable(0),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 1)),
			Expected:     SomethingNullable(1),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 2)),
			Expected:     SomethingNullable(2),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 3)),
			Expected:     SomethingNullable(3),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 4)),
			Expected:     SomethingNullable(4),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", 5)),
			Expected:     SomethingNullable(5),
		},
		{
			Datum:    fmt.Sprintf("%d", math.MaxInt64),
			Expected: SomethingNullable(math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                      x,
				Expected: SomethingNullable(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                      y,
				Expected: SomethingNullable(y),
			}
			tests = append(tests, test)
		}

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                      s,
				Expected: SomethingNullable(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                      s,
				Expected: SomethingNullable(y),
			}
			tests = append(tests, test)
		}

		{
			s := fmt.Sprintf("%d", x)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                      b,
				Expected: SomethingNullable(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)
			b := []byte(s)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                      b,
				Expected: SomethingNullable(y),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		var datum NullableType

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

func TestNullableTypeScanFail(t *testing.T) {
	tests := []struct {
		Datum interface{}
	}{
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
		var datum NullableType

		if err := datum.Scan(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
`
)
