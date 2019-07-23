package shared

const (
	NullableTypeValueStringTestTmpl =
`
import (
	"database/sql/driver"
	"fmt"
	"math"
	"testing"
)

func TestNullableTypeValue(t *testing.T) {
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
			Datum: SomethingNullable("1"),
			Expected:                "1",
		},
		{
			Datum: SomethingNullable("12"),
			Expected:                "12",
		},
		{
			Datum: SomethingNullable("-5"),
			Expected:                "-5",
		},
		{
			Datum: SomethingNullable(fmt.Sprintf("%d", math.MaxInt64)),
			Expected:                fmt.Sprintf("%d", math.MaxInt64),
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

func TestNullableTypeValueNothingNullable(t *testing.T) {
	datum := NothingNullable()

	_, err := datum.Value()
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: (%T) %v", err, err)
		return
	}

}

func TestNullableTypeValueNull(t *testing.T) {
	datum := Null()

	actual, err := datum.Value()
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected := driver.Value(nil); expected != actual {
		t.Errorf("Expected %v, but actually got %v", expected, actual)
		return
	}
}
`
)
