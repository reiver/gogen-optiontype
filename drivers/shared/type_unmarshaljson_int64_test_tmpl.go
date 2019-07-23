package shared

const (
	TypeUnmarshaljsonInt64TestTmpl =
`
import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"testing"
)

func TestTypeUnmarshalJSONStruct(t *testing.T) {

	jsonString := `+"`"+`{
		"id":5,
		"email":"joeblow@example.net",
		"when_created": "2017-04-12T13:02:28Z",
		"parent_id": 2
	}`+"`"+`

	type MyThing struct {
		ID          Type      `+"`"+`json:"id"`+"`"+`
		EMail       string    `+"`"+`json:"email"`+"`"+`
		WhenCreated time.Time `+"`"+`json:"when_created"`+"`"+`
		ParentID    Type      `+"`"+`json:"parent_id"`+"`"+`
	}

	var datum MyThing

	if err := json.Unmarshal([]byte(jsonString), &datum); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := Something(5), datum.ID; expected != actual {
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

	if expected, actual := Something(2),
		datum.ParentID; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v", expected, expected, actual, actual)
		return
	}

}

func TestTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		Datum    []byte
		Expected Type
	}{
		{
			Datum: []byte(fmt.Sprintf("%d", int64(math.MinInt64))),
			Expected:                   Something(math.MinInt64),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(-5))),
			Expected:                   Something(-5),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(-4))),
			Expected:                   Something(-4),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(-3))),
			Expected:                   Something(-3),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(-2))),
			Expected:                   Something(-2),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(-1))),
			Expected:                   Something(-1),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(0))),
			Expected:                   Something(0),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(1))),
			Expected:                   Something(1),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(2))),
			Expected:                   Something(2),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(3))),
			Expected:                   Something(3),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(4))),
			Expected:                   Something(4),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(5))),
			Expected:                   Something(5),
		},
		{
			Datum: []byte(fmt.Sprintf("%d", int64(math.MaxInt64))),
			Expected:                   Something(math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum    []byte
				Expected Type
			}{
				Datum: []byte(fmt.Sprintf("%d", x)),
				Expected:             Something(x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum    []byte
				Expected Type
			}{
				Datum: []byte(fmt.Sprintf("%d", y)),
				Expected:             Something(y),
			}
			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {
		var datum Type

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

func TestTypeUnmarshalJSONNull(t *testing.T) {
	var datum Type

	if err := datum.UnmarshalJSON([]byte("null")); nil == err {
		t.Errorf("Expected an error, but did not actually get one: (%T) %v", err, err)
		return
	}
}

func TestTypeUnmarshalJSONFail(t *testing.T) {
	tests := []struct {
		Datum    []byte
	}{
		{
			Datum: []byte(`+"`"+`"apple"`+"`"+`),
		},
		{
			Datum: []byte(`+"`"+`"banana"`+"`"+`),
		},
		{
			Datum: []byte(`+"`"+`"cherry"`+"`"+`),
		},
		{
			Datum: []byte(`+"`"+`"date"`+"`"+`),
		},



		{
			Datum: []byte("false"),
		},
		{
			Datum: []byte("true"),
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



		{
			Datum: []byte("null"),
		},
	}


	for testNumber, test := range tests {
		var datum Type

		if err := datum.UnmarshalJSON(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
`
)
