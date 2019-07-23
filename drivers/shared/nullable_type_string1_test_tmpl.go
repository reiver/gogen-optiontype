package shared

const (
	NullableTypeString1TestTmpl =
`
import (
	"fmt"
	"math"
	"testing"
)

func TestNullableTypeString(t *testing.T) {
	tests := []struct{
		Datum    NullableType
		Expected string
	}{
		{
			Datum:                                SomethingNullable(math.MinInt64),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", math.MinInt64),
		},
		{
			Datum                               : SomethingNullable(-5),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", -5),
		},
		{
			Datum                               : SomethingNullable(-4),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", -4),
		},
		{
			Datum                               : SomethingNullable(-3),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", -3),
		},
		{
			Datum                               : SomethingNullable(-2),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", -2),
		},
		{
			Datum                               : SomethingNullable(-1),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", -1),
		},
		{
			Datum                               : SomethingNullable(0),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", 0),
		},
		{
			Datum                               : SomethingNullable(1),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", 1),
		},
		{
			Datum                               : SomethingNullable(2),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", 2),
		},
		{
			Datum                               : SomethingNullable(3),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", 3),
		},
		{
			Datum                               : SomethingNullable(4),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", 4),
		},
		{
			Datum                               : SomethingNullable(5),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", 5),
		},
		{
			Datum:                                SomethingNullable(math.MaxInt64),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum NullableType
				Expected string
			}{
				Datum:                                SomethingNullable(x),
				Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum NullableType
				Expected string
			}{
				Datum:                                SomethingNullable(y),
				Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", y),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		if expected, actual := test.Expected, datum.String(); expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestNullableTypeStringNothingNullable(t *testing.T) {

	datum := NothingNullable()

	if expected, actual := "{{.Pkg}}.NothingNullable()", datum.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}

func TestNullableTypeStringNull(t *testing.T) {

	datum := Null()

	if expected, actual := "{{.Pkg}}.Null()", datum.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}
`
)
