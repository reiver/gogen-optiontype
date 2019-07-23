package shared

const (
	TypeScanStringTestTmpl =
`
import (
	"fmt"
	"time"

	"testing"
)

func TestTypeScan(t *testing.T) {
	tests := []struct {
		Datum    interface{}
		Expected Type
	}{
		{
			Datum:    Something("apple"),
			Expected: Something("apple"),
		},
		{
			Datum:    Something("BANANA"),
			Expected: Something("BANANA"),
		},
		{
			Datum:    Something("Cherry"),
			Expected: Something("Cherry"),
		},
		{
			Datum:    Something("dATE"),
			Expected: Something("dATE"),
		},



		{
			Datum:              "apple",
			Expected: Something("apple"),
		},
		{
			Datum:              "BANANA",
			Expected: Something("BANANA"),
		},
		{
			Datum:              "Cherry",
			Expected: Something("Cherry"),
		},
		{
			Datum:              "dATE",
			Expected: Something("dATE"),
		},



		{
			Datum:       []byte("apple"),
			Expected: Something("apple"),
		},
		{
			Datum:       []byte("BANANA"),
			Expected: Something("BANANA"),
		},
		{
			Datum:       []byte("Cherry"),
			Expected: Something("Cherry"),
		},
		{
			Datum:       []byte("dATE"),
			Expected: Something("dATE"),
		},



		{
			Datum:    Something("Hello world!"),
			Expected: Something("Hello world!"),
		},



		{
			Datum:              "Hello world!",
			Expected: Something("Hello world!"),
		},



		{
			Datum:       []byte("Hello world!"),
			Expected: Something("Hello world!"),
		},



		{
			Datum:              "-5",
			Expected: Something("-5"),
		},
		{
			Datum:              "-4",
			Expected: Something("-4"),
		},
		{
			Datum:              "-3",
			Expected: Something("-3"),
		},
		{
			Datum:              "-3",
			Expected: Something("-3"),
		},
		{
			Datum:              "-2",
			Expected: Something("-2"),
		},
		{
			Datum:              "-1",
			Expected: Something("-1"),
		},
		{
			Datum:              "0",
			Expected: Something("0"),
		},
		{
			Datum:              "1",
			Expected: Something("1"),
		},
		{
			Datum:              "2",
			Expected: Something("2"),
		},
		{
			Datum:              "3",
			Expected: Something("3"),
		},
		{
			Datum:              "4",
			Expected: Something("4"),
		},
		{
			Datum:              "5",
			Expected: Something("5"),
		},



		{
			Datum:       []byte("-5"),
			Expected: Something("-5"),
		},
		{
			Datum:       []byte("-4"),
			Expected: Something("-4"),
		},
		{
			Datum:       []byte("-3"),
			Expected: Something("-3"),
		},
		{
			Datum:       []byte("-2"),
			Expected: Something("-2"),
		},
		{
			Datum:       []byte("-1"),
			Expected: Something("-1"),
		},
		{
			Datum:       []byte("0"),
			Expected: Something("0"),
		},
		{
			Datum:       []byte("1"),
			Expected: Something("1"),
		},
		{
			Datum:       []byte("2"),
			Expected: Something("2"),
		},
		{
			Datum:       []byte("3"),
			Expected: Something("3"),
		},
		{
			Datum:       []byte("4"),
			Expected: Something("4"),
		},
		{
			Datum:       []byte("5"),
			Expected: Something("5"),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			s := fmt.Sprintf("%d", x)

			test := struct {
				Datum    interface{}
				Expected Type
			}{
				Datum:              s,
				Expected: Something(s),
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
				Expected: Something(s),
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
				Expected: Something(s),
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
				Expected: Something(s),
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
		var datum Type

		if err := datum.Scan(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
`
)
