package shared

const (
	TypeGoStringTestTmpl =
`
import (
	"fmt"
	"math"
	"testing"
)

func TestTypeString(t *testing.T) {
	tests := []struct{
		Datum    Type
		Expected string
	}{
		{
			Datum:                                Something(math.MinInt64),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", math.MinInt64),
		},
		{
			Datum:                                Something(-5),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", -5),
		},
		{
			Datum:                                Something(-4),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", -4),
		},
		{
			Datum:                                Something(-3),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", -3),
		},
		{
			Datum:                                Something(-2),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", -2),
		},
		{
			Datum:                                Something(-1),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", -1),
		},
		{
			Datum:                                Something(0),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", 0),
		},
		{
			Datum:                                Something(1),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", 1),
		},
		{
			Datum:                                Something(2),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", 2),
		},
		{
			Datum:                                Something(3),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", 3),
		},
		{
			Datum:                                Something(4),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", 4),
		},
		{
			Datum:                                Something(5),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", 5),
		},
		{
			Datum:                                Something(math.MaxInt64),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", math.MaxInt64),
		},
	}

	const numRandTests = 40
	for i := 0; i < numRandTests; i++ {
		x := randomness.Int63()

		{
			test := struct {
				Datum Type
				Expected string
			}{
				Datum:                                Something(x),
				Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum Type
				Expected string
			}{
				Datum:                                Something(y),
				Expected: fmt.Sprintf("{{.Pkg}}.Something(%d)", y),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {
		datum := test.Datum

		if expected, actual := test.Expected, datum.GoString(); expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeStringNothing(t *testing.T) {

	datum := Nothing()

	if expected, actual := "{{.Pkg}}.Nothing()", datum.GoString(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}
`
)
