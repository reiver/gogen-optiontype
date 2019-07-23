package shared

const (
	TypeGoStringStringTestTmpl =
`
import (
	"fmt"
	"testing"
)

func TestTypeGoString(t *testing.T) {
	tests := []struct{
		Datum    Type
		Expected string
	}{
		{
			Datum:                                Something("apple"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "apple"),
		},
		{
			Datum:                                Something("BANANA"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "BANANA"),
		},
		{
			Datum:                                Something("Cherry"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "Cherry"),
		},
		{
			Datum:                                Something("dATE"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "dATE"),
		},



		{
			Datum:                                Something("Hello world!"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "Hello world!"),
		},



		{
			Datum:                                Something("-5"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "-5"),
		},
		{
			Datum:                                Something("-4"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "-4"),
		},
		{
			Datum:                                Something("-3"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "-3"),
		},
		{
			Datum:                                Something("-2"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "-2"),
		},
		{
			Datum:                                Something("-1"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "-1"),
		},
		{
			Datum:                                Something("0"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "0"),
		},
		{
			Datum:                                Something("1"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "1"),
		},
		{
			Datum:                                Something("2"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "2"),
		},
		{
			Datum:                                Something("3"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "3"),
		},
		{
			Datum:                                Something("4"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "4"),
		},
		{
			Datum:                                Something("5"),
			Expected: fmt.Sprintf("{{.Pkg}}.Something(%q)", "5"),
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
				Datum:                  Something(fmt.Sprintf("%d", x)),
				Expected: fmt.Sprintf("{{.Pkg}}.Something(\"%d\")", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum Type
				Expected string
			}{
				Datum:                  Something(fmt.Sprintf("%d", y)),
				Expected: fmt.Sprintf("{{.Pkg}}.Something(\"%d\")", y),
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

func TestTypeGoStringNothing(t *testing.T) {

	datum := Nothing()

	if expected, actual := "{{.Pkg}}.Nothing()", datum.GoString(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}
`
)
