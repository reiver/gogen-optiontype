package gendriver

import (
	"bytes"

	"testing"
)

func TestDefaultRenderer(t *testing.T) {

	tests := []struct{
		FileName       string
		FileTmpl       string
		RendererParams RendererParams
		Expected       string
	}{
		{
			FileName: "apple.go",
			FileTmpl:
`type Type struct {
	value {{.Type}}
}
`,
			RendererParams: RendererParams{
				Pkg:  "itemid",
				Type: "int64",
			},
			Expected:
`package itemid

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

type Type struct {
	value int64
}
`,
		},

		{
			FileName: "banana_test.go",
			FileTmpl:
`import (
	"testing"
)

func TestType{{.Type}}(t *testing.T) {

	t.Errorf("Uh oh!")

}
`,
			RendererParams: RendererParams{
				Pkg:  "somethingid",
				Type: "string",
			},
			Expected:
`package somethingid

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"testing"
)

func TestTypestring(t *testing.T) {

	t.Errorf("Uh oh!")

}
`,
		},
	}


	for testNumber, test := range tests {

		renderer := DefaultRenderer{
			FileName: test.FileName,
			FileTmpl: test.FileTmpl,
		}

		actualFileName, actualWriterTo, err := renderer.WriterTo(test.RendererParams)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.FileName, actualFileName; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}

		{
			var buffer bytes.Buffer

			n, err := actualWriterTo.WriteTo(&buffer)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
				continue
			}

			if expected, actual := n, int64(len(test.Expected)); expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				t.Errorf("EXPECTED: %d", expected)
				t.Errorf("ACTUAL:   %d", actual)

				t.Errorf("EXPECTED:\n%s", test.Expected)
				t.Errorf("ACTUAL:\n%s", buffer.String())
				continue
			}
			if expected, actual := test.Expected, buffer.String(); expected != actual {
				t.Errorf("For test #%d...", testNumber)
				t.Errorf("EXPECTED:\n%v", expected)
				t.Errorf("ACTUAL:\n%v", actual)
				continue
			}
		}
	}
}
