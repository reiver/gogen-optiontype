package shared

const (
	TypeMarshaljsonStringTestTmpl =
`
import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"testing"
)

func TestTypeMarshalJSONStruct(t *testing.T) {
	type MyThing struct {
		ID          Type      `+"`"+`json:"id"`+"`"+`
		EMail       string    `+"`"+`json:"email"`+"`"+`
		WhenCreated time.Time `+"`"+`json:"when_created"`+"`"+`
		ParentID    Type      `+"`"+`json:"parent_id"`+"`"+`
	}

	var datum MyThing

	datum.ID = Something("some-ID")
	datum.EMail = "joeblow@example.net"
	{
		tm, err := time.Parse(time.RFC3339, "2017-04-12T13:02:28Z")
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		datum.WhenCreated = tm
	}
	datum.ParentID = Something("another-ID")

	actualBytes, err := json.Marshal(datum)
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	{
		expected := `+"`"+`{`+"`"+` +
			`+"`"+`"id":"some-ID",`+"`"+` +
			`+"`"+`"email":"joeblow@example.net",`+"`"+` +
			`+"`"+`"when_created":"2017-04-12T13:02:28Z",`+"`"+` +
			`+"`"+`"parent_id":"another-ID"`+"`"+` +
			`+"`"+`}`+"`"+`

		if actual := string(actualBytes); expected != actual {
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			return
		}
	}
}

func TestTypeMarshalJSON(t *testing.T) {
	tests := []struct {
		Datum    Type
		Expected string
	}{
		{
			Datum: Something("apple"),
			Expected: `+"`"+`"apple"`+"`"+`,
		},
		{
			Datum: Something("BANANA"),
			Expected: `+"`"+`"BANANA"`+"`"+`,
		},
		{
			Datum: Something("Cherry"),
			Expected: `+"`"+`"Cherry"`+"`"+`,
		},
		{
			Datum: Something("dATE"),
			Expected: `+"`"+`"dATE"`+"`"+`,
		},



		{
			Datum: Something("Hello world!"),
			Expected: `+"`"+`"Hello world!"`+"`"+`,
		},



		{
			Datum: Something(fmt.Sprintf(       "%d",        math.MinInt64)),
			Expected:        fmt.Sprintf(`+"`"+`"%d"`+"`"+`, math.MinInt64),
		},
		{
			Datum: Something("-5"),
			Expected: `+"`"+`"-5"`+"`"+`,
		},
		{
			Datum: Something("-4"),
			Expected: `+"`"+`"-4"`+"`"+`,
		},
		{
			Datum: Something("-3"),
			Expected: `+"`"+`"-3"`+"`"+`,
		},
		{
			Datum: Something("-2"),
			Expected: `+"`"+`"-2"`+"`"+`,
		},
		{
			Datum: Something("-1"),
			Expected: `+"`"+`"-1"`+"`"+`,
		},
		{
			Datum: Something("0"),
			Expected: `+"`"+`"0"`+"`"+`,
		},
		{
			Datum: Something("1"),
			Expected: `+"`"+`"1"`+"`"+`,
		},
		{
			Datum: Something("2"),
			Expected: `+"`"+`"2"`+"`"+`,
		},
		{
			Datum: Something("3"),
			Expected: `+"`"+`"3"`+"`"+`,
		},
		{
			Datum: Something("4"),
			Expected: `+"`"+`"4"`+"`"+`,
		},
		{
			Datum: Something("5"),
			Expected: `+"`"+`"5"`+"`"+`,
		},
		{
			Datum: Something(fmt.Sprintf(       "%d",        math.MaxInt64)),
			Expected:        fmt.Sprintf(`+"`"+`"%d"`+"`"+`, math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    Type
				Expected string
			}{
				Datum: Something(fmt.Sprintf(       "%d",        x)),
				Expected:        fmt.Sprintf(`+"`"+`"%d"`+"`"+`, x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    Type
				Expected string
			}{
				Datum: Something(fmt.Sprintf(       "%d",        y)),
				Expected:        fmt.Sprintf(`+"`"+`"%d"`+"`"+`, y),
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

func TestTypeMarshalJSONNothing(t *testing.T) {
	datum := Nothing()

	_, err := datum.MarshalJSON()
	if nil == err {
		t.Errorf("Expected an error, but not actually get one: (%T) %v", err, err)
		return
	}
}
`
)
