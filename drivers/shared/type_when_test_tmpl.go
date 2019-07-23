package shared

const (
	TypeWhenTestTmpl =
`
import (
	"testing"
)

func TestTypeWhenNothing(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		Nothing().WhenNothing(func(){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var value {{.Type}}
		Something(value).WhenNothing(func(){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}
}

func TestTypeWhenSomething(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		Nothing().WhenSomething(func(datum {{.Type}}){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var datum {{.Type}}
		Something(datum).WhenSomething(func(datum {{.Type}}){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}
}
`
)
