package shared

const (
	NullableTypeScanStringTestTmpl =
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
			Datum:            Something("apple"),
			Expected: SomethingNullable("apple"),
		},
		{
			Datum:            Something("BANANA"),
			Expected: SomethingNullable("BANANA"),
		},
		{
			Datum:            Something("Cherry"),
			Expected: SomethingNullable("Cherry"),
		},
		{
			Datum:            Something("dATE"),
			Expected: SomethingNullable("dATE"),
		},



		{
			Datum:    SomethingNullable("apple"),
			Expected: SomethingNullable("apple"),
		},
		{
			Datum:    SomethingNullable("BANANA"),
			Expected: SomethingNullable("BANANA"),
		},
		{
			Datum:    SomethingNullable("Cherry"),
			Expected: SomethingNullable("Cherry"),
		},
		{
			Datum:    SomethingNullable("dATE"),
			Expected: SomethingNullable("dATE"),
		},



		{
			Datum:                 "apple",
			Expected: SomethingNullable("apple"),
		},
		{
			Datum:                 "BANANA",
			Expected: SomethingNullable("BANANA"),
		},
		{
			Datum:                 "Cherry",
			Expected: SomethingNullable("Cherry"),
		},
		{
			Datum:                 "dATE",
			Expected: SomethingNullable("dATE"),
		},



		{
			Datum:                 fmt.Sprintf("%d", math.MinInt64),
			Expected: SomethingNullable(fmt.Sprintf("%d", math.MinInt64)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -5),
			Expected: SomethingNullable(fmt.Sprintf("%d", -5)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -4),
			Expected: SomethingNullable(fmt.Sprintf("%d", -4)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -3),
			Expected: SomethingNullable(fmt.Sprintf("%d", -3)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -2),
			Expected: SomethingNullable(fmt.Sprintf("%d", -2)),
		},
		{
			Datum:                 fmt.Sprintf("%d", -1),
			Expected: SomethingNullable(fmt.Sprintf("%d", -1)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 0),
			Expected: SomethingNullable(fmt.Sprintf("%d", 0)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 1),
			Expected: SomethingNullable(fmt.Sprintf("%d", 1)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 2),
			Expected: SomethingNullable(fmt.Sprintf("%d", 2)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 3),
			Expected: SomethingNullable(fmt.Sprintf("%d", 3)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 4),
			Expected: SomethingNullable(fmt.Sprintf("%d", 4)),
		},
		{
			Datum:                 fmt.Sprintf("%d", 5),
			Expected: SomethingNullable(fmt.Sprintf("%d", 5)),
		},
		{
			Datum:                 fmt.Sprintf("%d", math.MaxInt64),
			Expected: SomethingNullable(fmt.Sprintf("%d", math.MaxInt64)),
		},



		{
			Datum:          []byte(fmt.Sprintf("%d", math.MinInt64)),
			Expected: SomethingNullable(fmt.Sprintf("%d", math.MinInt64)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -5)),
			Expected: SomethingNullable(fmt.Sprintf("%d", -5)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -4)),
			Expected: SomethingNullable(fmt.Sprintf("%d", -4)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -3)),
			Expected: SomethingNullable(fmt.Sprintf("%d", -3)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -2)),
			Expected: SomethingNullable(fmt.Sprintf("%d", -2)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", -1)),
			Expected: SomethingNullable(fmt.Sprintf("%d", -1)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 0)),
			Expected: SomethingNullable(fmt.Sprintf("%d", 0)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 1)),
			Expected: SomethingNullable(fmt.Sprintf("%d", 1)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 2)),
			Expected: SomethingNullable(fmt.Sprintf("%d", 2)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 3)),
			Expected: SomethingNullable(fmt.Sprintf("%d", 3)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 4)),
			Expected: SomethingNullable(fmt.Sprintf("%d", 4)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", 5)),
			Expected: SomethingNullable(fmt.Sprintf("%d", 5)),
		},
		{
			Datum:          []byte(fmt.Sprintf("%d", math.MaxInt64)),
			Expected: SomethingNullable(fmt.Sprintf("%d", math.MaxInt64)),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    interface{}
				Expected NullableType
			}{
				Datum:                 s,
				Expected: SomethingNullable(s),
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
				Datum:                 s,
				Expected: SomethingNullable(s),
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
				Datum:                 b,
				Expected: SomethingNullable(s),
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
				Datum:                 b,
				Expected: SomethingNullable(s),
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
			Datum: int64(-5),
		},
		{
			Datum: int64(-4),
		},
		{
			Datum: int64(-3),
		},
		{
			Datum: int64(-2),
		},
		{
			Datum: int64(-1),
		},
		{
			Datum: int64(0),
		},
		{
			Datum: int64(1),
		},
		{
			Datum: int64(2),
		},
		{
			Datum: int64(3),
		},
		{
			Datum: int64(4),
		},
		{
			Datum: int64(5),
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
