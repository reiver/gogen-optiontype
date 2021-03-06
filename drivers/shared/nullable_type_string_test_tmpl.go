package shared

const (
	NullableTypeStringTestTmpl =
`
import (
	"fmt"
	"math"

	"testing"
)

func TestNullableTypeString(t *testing.T) {
	tests := []struct {
		Datum    NullableType
		Expected string
	}{
		{
			Datum: SomethingNullable("apple"),
			Expected:                "apple",
		},
		{
			Datum: SomethingNullable("BANANA"),
			Expected:                "BANANA",
		},
		{
			Datum: SomethingNullable("Cherry"),
			Expected:                "Cherry",
		},
		{
			Datum: SomethingNullable("dATE"),
			Expected:                "dATE",
		},



		{
			Datum: SomethingNullable("Hello world!"),
			Expected:                "Hello world!",
		},



		{
			Datum: SomethingNullable(fmt.Sprintf("%d", math.MinInt64)),
			Expected:                fmt.Sprintf("%d", math.MinInt64),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", -5)),
			Expected:                fmt.Sprintf("%d", -5),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", -4)),
			Expected:                fmt.Sprintf("%d", -4),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", -3)),
			Expected:                fmt.Sprintf("%d", -3),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", -2)),
			Expected:                fmt.Sprintf("%d", -2),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", -1)),
			Expected:                fmt.Sprintf("%d", -1),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", 0)),
			Expected:                fmt.Sprintf("%d", 0),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", 1)),
			Expected:                fmt.Sprintf("%d", 1),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", 2)),
			Expected:                fmt.Sprintf("%d", 2),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", 3)),
			Expected:                fmt.Sprintf("%d", 3),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", 4)),
			Expected:                fmt.Sprintf("%d", 4),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", 5)),
			Expected:                fmt.Sprintf("%d", 5),
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", math.MaxInt64)),
			Expected:                fmt.Sprintf("%d", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum: SomethingNullable(s),
				Expected:                s,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum: SomethingNullable(s),
				Expected:                s,
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		actual, err := datum.String()
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

func TestNullableTypeStringNothingNullable(t *testing.T) {

	datum := NothingNullable()

	_, err := datum.String()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNothingNullable, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}

func TestNullableTypeStringNull(t *testing.T) {

	datum := Null()

	_, err := datum.String()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %v", err)
		return
	}

	if expected, actual := errNull, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}
}
`
)
