package shared

const (
	NullableTypeWhenTestTmpl =
`
import (
	"testing"
)

func TestNullableTypeWhenNothing(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		NothingNullable().WhenNothing(func(){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		Null().WhenNothing(func(){
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
		SomeNullable(value).WhenNothing(func(){
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

		NothingNullable().WhenNull(func(){
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

		NothingNullable().WhenSome(func(datum {{.Type}}){
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
