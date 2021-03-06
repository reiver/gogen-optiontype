package shared

const (
	TypeValueInt64TestTmpl =
`
import (
	"math"
	"testing"
)

func TestTypeValue(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected int64
	}{
		{
			Datum: Something(1),
			Expected:  int64(1),
		},
		{
			Datum: Something(12),
			Expected:  int64(12),
		},
		{
			Datum: Something(-5),
			Expected:  int64(-5),
		},
		{
			Datum: Something(math.MaxInt64),
			Expected:  int64(math.MaxInt64),
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
