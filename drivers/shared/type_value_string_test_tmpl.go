package shared

const (
	TypeValueStringTestTmpl =
`
import (
	"testing"
)

func TestTypeValue(t *testing.T) {
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
			Datum: Something("1"),
			Expected:        "1",
		},
		{
			Datum: Something("12"),
			Expected:        "12",
		},
		{
			Datum: Something("-5"),
			Expected:        "-5",
		},
	}


	for testNumber, test := range tests {
		actual, err := test.Datum.Value()
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

func TestTypeValueNothing(t *testing.T) {
	datum := Nothing()

	_, err := datum.Value()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: (%T) %v", err, err)
		return
	}

}
`
)
