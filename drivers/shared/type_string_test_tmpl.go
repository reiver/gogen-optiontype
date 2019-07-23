package shared

const (
	TypeStringTestTmpl =
`
import (
	"fmt"

	"testing"
)

func TestTypeString(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected string
	}{
		{
			Datum: Something("apple"),
			Expected:        "apple",
		},
		{
			Datum: Something("BANANA"),
			Expected:        "BANANA",
		},
		{
			Datum: Something("Cherry"),
			Expected:        "Cherry",
		},
		{
			Datum: Something("dATE"),
			Expected:        "dATE",
		},



		{
			Datum: Something("Hello world!"),
			Expected:        "Hello world!",
		},



		{
			Datum: Something("-5"),
			Expected:        "-5",
		},
		{
			Datum: Something("-4"),
			Expected:        "-4",
		},
		{
			Datum: Something("-3"),
			Expected:        "-3",
		},
		{
			Datum: Something("-2"),
			Expected:        "-2",
		},
		{
			Datum: Something("-1"),
			Expected:        "-1",
		},
		{
			Datum: Something("0"),
			Expected:        "0",
		},
		{
			Datum: Something("1"),
			Expected:        "1",
		},
		{
			Datum: Something("2"),
			Expected:        "2",
		},
		{
			Datum: Something("3"),
			Expected:        "3",
		},
		{
			Datum: Something("4"),
			Expected:        "4",
		},
		{
			Datum: Something("5"),
			Expected:        "5",
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    Type
				Expected string
			}{
				Datum: Something(s),
				Expected:   s,
			}
			tests = append(tests, test)
		}

		{
			y := -x

			s := fmt.Sprintf("%d", y)

			test := struct {
				Datum    Type
				Expected string
			}{
				Datum: Something(s),
				Expected:   s,
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

func TestTypeStringNothing(t *testing.T) {

	datum := Nothing()

	_, err := datum.String()
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
