package shared

const (
	NullableTypeMarshaljsonInt64TestTmpl =
`
import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"testing"
)

func TestNullableTypeMarshalJSONStruct(t *testing.T) {
	type MyThing struct {
		ID          NullableType `+"`"+`json:"id"`+"`"+`
		EMail       string       `+"`"+`json:"email"`+"`"+`
		WhenCreated time.Time    `+"`"+`json:"when_created"`+"`"+`
		ParentID    NullableType `+"`"+`json:"parent_id"`+"`"+`
	}

	var datum MyThing

	datum.ID = SomethingNullable(5)
	datum.EMail = "joeblow@example.net"
	{
		tm, err := time.Parse(time.RFC3339, "2017-04-12T13:02:28Z")
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		datum.WhenCreated = tm
	}
	datum.ParentID = SomethingNullable(2)

	actualBytes, err := json.Marshal(datum)
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	{
		expected := `+"`"+`{`+"`"+` +
			`+"`"+`"id":5,`+"`"+` +
			`+"`"+`"email":"joeblow@example.net",`+"`"+` +
			`+"`"+`"when_created":"2017-04-12T13:02:28Z",`+"`"+` +
			`+"`"+`"parent_id":2`+"`"+` +
			`+"`"+`}`+"`"+`

		if actual := string(actualBytes); expected != actual {
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			return
		}
	}
}

func TestNullableTypeMarshalJSON(t *testing.T) {
	tests := []struct {
		Datum    NullableType
		Expected string
	}{
		{
			Datum:    SomethingNullable(math.MinInt64),
			Expected: fmt.Sprintf("%d", math.MinInt64),
		},
		{
			Datum: SomethingNullable(5),
			Expected:               "5",
		},
		{
			Datum: SomethingNullable(-4),
			Expected:               "-4",
		},
		{
			Datum: SomethingNullable(-3),
			Expected:               "-3",
		},
		{
			Datum: SomethingNullable(-2),
			Expected:               "-2",
		},
		{
			Datum: SomethingNullable(-1),
			Expected:               "-1",
		},
		{
			Datum: SomethingNullable(0),
			Expected:               "0",
		},
		{
			Datum: SomethingNullable(1),
			Expected:               "1",
		},
		{
			Datum: SomethingNullable(2),
			Expected:               "2",
		},
		{
			Datum: SomethingNullable(3),
			Expected:               "3",
		},
		{
			Datum: SomethingNullable(4),
			Expected:               "4",
		},
		{
			Datum: SomethingNullable(5),
			Expected:               "5",
		},
		{
			Datum:    SomethingNullable(math.MaxInt64),
			Expected: fmt.Sprintf("%d", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum:    SomethingNullable(x),
				Expected: fmt.Sprintf("%d", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    NullableType
				Expected string
			}{
				Datum:    SomethingNullable(y),
				Expected: fmt.Sprintf("%d", y),
			}
			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Datum)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, string(actualBytes); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}
	}
}

func TestNullableTypeMarshalJSONNothingNullable(t *testing.T) {
	datum := NothingNullable()

	_, err := datum.MarshalJSON()
	if nil == err {
		t.Errorf("Expected an error, but not actually get one: (%T) %v", err, err)
		return
	}
}

func TestNullableTypeMarshalJSONull(t *testing.T) {
	datum := Null()

	actualBytes, err := datum.MarshalJSON()
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := "null", string(actualBytes); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}
`
)
