package shared

const (
	NullableTypeUnmarshaljsonStringTestTmpl =
`
import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"testing"
)

func TestNullableTypeUnmarshalJSONStruct(t *testing.T) {

	jsonString := `+"`"+`{
		"id":"some-ID",
		"email":"joeblow@example.net",
		"when_created": "2017-04-12T13:02:28Z",
		"parent_id": "another-ID"
	}`+"`"+`

	type MyThing struct {
		ID          NullableType `+"`"+`json:"id"`+"`"+`
		EMail       string       `+"`"+`json:"email"`+"`"+`
		WhenCreated time.Time    `+"`"+`json:"when_created"`+"`"+`
		ParentID    NullableType `+"`"+`json:"parent_id"`+"`"+`
	}

	var datum MyThing

	if err := json.Unmarshal([]byte(jsonString), &datum); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := SomethingNullable("some-ID"), datum.ID; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}

	if expected, actual := "joeblow@example.net", datum.EMail; expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}

	{
		expected, err := time.Parse(time.RFC3339, "2017-04-12T13:02:28Z")
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if actual := datum.WhenCreated; !expected.Equal(actual) {
			t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
			return
		}
	}

	if expected, actual := SomethingNullable("another-ID"),
		datum.ParentID; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}

}

func TestNullableTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		Datum    []byte
		Expected NullableType
	}{
		{
			Datum: []byte("null"),
			Expected:      Null(),
		},



		{
			Datum:        []byte(`+"`"+`"apple"`+"`"+`),
			Expected: SomethingNullable("apple"),
		},
		{
			Datum:        []byte(`+"`"+`"BANANA"`+"`"+`),
			Expected: SomethingNullable("BANANA"),
		},
		{
			Datum:        []byte(`+"`"+`"Cherry"`+"`"+`),
			Expected: SomethingNullable("Cherry"),
		},
		{
			Datum:        []byte(`+"`"+`"dATE"`+"`"+`),
			Expected: SomethingNullable("dATE"),
		},



		{
			Datum:        []byte(`+"`"+`"Hello world!"`+"`"+`),
			Expected: SomethingNullable("Hello world!"),
		},



		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(math.MinInt64))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(math.MinInt64))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(-5))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(-5))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(-4))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",  int64(-4))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(-3))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(-3))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(-2))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(-2))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(-1))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(-1))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(0))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(0))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(1))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(1))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(2))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(2))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(3))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(3))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(4))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(4))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(5))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(5))),
		},
		{
			Datum:               []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, int64(math.MaxInt64))),
			Expected: SomethingNullable(fmt.Sprintf(       "%d",        int64(math.MaxInt64))),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    []byte
				Expected NullableType
			}{
				Datum:         []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, x)),
				Expected: SomethingNullable(fmt.Sprintf( "%d",  x)),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    []byte
				Expected NullableType
			}{
				Datum:         []byte(fmt.Sprintf(`+"`"+`"%d"`+"`"+`, y)),
				Expected: SomethingNullable(fmt.Sprintf( "%d",  y)),
			}
			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {
		var datum NullableType

		if err := datum.UnmarshalJSON(test.Datum); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, datum; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestNullableTypeUnmarshalJSONFail(t *testing.T) {
	tests := []struct {
		Datum    []byte
	}{
		{
			Datum: []byte("false"),
		},
		{
			Datum: []byte("true"),
		},



		{
			Datum: []byte("-5"),
		},
		{
			Datum: []byte("-4"),
		},
		{
			Datum: []byte("-3"),
		},
		{
			Datum: []byte("-2"),
		},
		{
			Datum: []byte("-1"),
		},
		{
			Datum: []byte("0"),
		},
		{
			Datum: []byte("1"),
		},
		{
			Datum: []byte("2"),
		},
		{
			Datum: []byte("3"),
		},
		{
			Datum: []byte("4"),
		},
		{
			Datum: []byte("5"),
		},



		{
			Datum: []byte("-1.1"),
		},
		{
			Datum: []byte("1.1"),
		},



		{
			Datum: []byte(`+"`"+`{"apple":2}`+"`"+`),
		},
	}


	for testNumber, test := range tests {
		var datum NullableType

		if err := datum.UnmarshalJSON(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
`
)
