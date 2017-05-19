package shared

const (
	NullableTypeWhenTestTmpl =
`
import (
	"testing"
)

func TestNullableTypeWhenNone(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		NoneNullable().WhenNone(func(){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		Null().WhenNone(func(){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var value {{.Type}}
		SomeNullable(value).WhenNone(func(){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}
}

func TestNullableTypeWhenNull(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		NoneNullable().WhenNull(func(){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		Null().WhenNull(func(){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var value {{.Type}}
		SomeNullable(value).WhenNull(func(){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}
}

func TestNullableTypeWhenSome(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		NoneNullable().WhenSome(func(datum {{.Type}}){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		Null().WhenSome(func(datum {{.Type}}){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var datum {{.Type}}
		SomeNullable(datum).WhenSome(func(datum {{.Type}}){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}
}
`
)
