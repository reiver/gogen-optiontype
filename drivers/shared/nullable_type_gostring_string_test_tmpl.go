package shared

const (
	NullableTypeGoStringStringTestTmpl =
`
import (
	"fmt"
	"testing"
)

func TestNullableTypeGoString(t *testing.T) {
	tests := []struct{
		Datum    NullableType
		Expected string
	}{
		{
			Datum                               : SomethingNullable("apple"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "apple"),
		},
		{
			Datum                               : SomethingNullable("BANANA"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "BANANA"),
		},
		{
			Datum                               : SomethingNullable("Cherry"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "Cherry"),
		},
		{
			Datum                               : SomethingNullable("dATE"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "dATE"),
		},



		{
			Datum                               : SomethingNullable("Hello world!"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "Hello world!"),
		},



		{
			Datum                               : SomethingNullable("-5"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "-5"),
		},
		{
			Datum                               : SomethingNullable("-4"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "-4"),
		},
		{
			Datum                               : SomethingNullable("-3"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "-3"),
		},
		{
			Datum                               : SomethingNullable("-2"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "-2"),
		},
		{
			Datum                               : SomethingNullable("-1"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "-1"),
		},
		{
			Datum                               : SomethingNullable("0"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "0"),
		},
		{
			Datum                               : SomethingNullable("1"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "1"),
		},
		{
			Datum                               : SomethingNullable("2"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "2"),
		},
		{
			Datum                               : SomethingNullable("3"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "3"),
		},
		{
			Datum                               : SomethingNullable("4"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "4"),
		},
		{
			Datum                               : SomethingNullable("5"),
			Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", "5"),
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
				Datum:                  SomethingNullable(fmt.Sprintf("%d", x)),
				Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(\"%d\")", x),
			}
			tests = append(tests, test)
		}

		{
			y := -x

			test := struct {
				Datum NullableType
				Expected string
			}{
				Datum:                  SomethingNullable(fmt.Sprintf("%d", y)),
				Expected: fmt.Sprintf("{{.Pkg}}.SomethingNullable(\"%d\")", y),
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

func TestNullableTypeGoStringNothingNullable(t *testing.T) {

	datum := NothingNullable()

	if expected, actual := "{{.Pkg}}.NothingNullable()", datum.GoString(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}

func TestNullableTypeGoStringNull(t *testing.T) {

	datum := Null()

	if expected, actual := "{{.Pkg}}.Null()", datum.GoString(); expected != actual {
		t.Errorf("Expected %q, but actually got %q", expected, actual)
		return
	}
}
`
)
