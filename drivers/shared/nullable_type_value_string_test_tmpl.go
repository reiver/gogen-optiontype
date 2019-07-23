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
			Datum: SomeNullable("apple"),
			Expected:           "apple",
		},
		{
			Datum: SomeNullable("BANANA"),
			Expected:           "BANANA",
		},
		{
			Datum: SomeNullable("Cherry"),
			Expected:           "Cherry",
		},
		{
			Datum: SomeNullable("dATE"),
			Expected:           "dATE",
		},



		{
			Datum: SomeNullable("Hello world!"),
			Expected:           "Hello world!",
		},



		{
			Datum: SomeNullable("1"),
			Expected:           "1",
		},
		{
			Datum: SomeNullable("12"),
			Expected:           "12",
		},
		{
			Datum: SomeNullable("-5"),
			Expected:           "-5",
		},
		{
			Datum: SomeNullable(fmt.Sprintf("%d", math.MaxInt64)),
			Expected:           fmt.Sprintf("%d", math.MaxInt64),
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
