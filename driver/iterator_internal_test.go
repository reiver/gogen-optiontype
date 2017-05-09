package gendriver

import (
	"bytes"

	"testing"
)

func TestInternalIterator(t *testing.T) {

	tests := []struct{
		RendererParams RendererParams
		Renderers    []Renderer
		Expected  map[string]string
	}{
		{
			RendererParams: RendererParams{
				Pkg:  "thatid",
				Type: "int64",
			},
			Renderers: []Renderer{
				DefaultRenderer{
					FileName: "apple.go",
					FileTmpl:
`package {{.Pkg}}

type Type struct {
	value {{.Type}}
}
`,
				},
			},
			Expected: map[string]string{
				"apple.go":
`package thatid

type Type struct {
	value int64
}
`,
			},
		},

		{
			RendererParams: RendererParams{
				Pkg:  "thoseid",
				Type: "string",
			},
			Renderers: []Renderer{
				DefaultRenderer{
					FileName: "banana.go",
					FileTmpl:
`package {{.Pkg}}

type Type struct {
	value {{.Type}}
}
`,
				},
				DefaultRenderer{
					FileName: "banana_test.go",
					FileTmpl:
`package {{.Pkg}}

import (
	"testing"
)

func TestType{{.Type}}(t *testing.T) {
	t.Errorf("This should not happen.")
}
`,
				},
			},
			Expected: map[string]string{
				"banana.go":
`package thoseid

type Type struct {
	value string
}
`,
				"banana_test.go":
`package thoseid

import (
	"testing"
)

func TestTypestring(t *testing.T) {
	t.Errorf("This should not happen.")
}
`,
			},
		},
	}


	Loop: for testNumber, test := range tests {

		var internalIter internalIterator

		if err := internalIter.copyFrom(test.Renderers); nil != err {
			t.Errorf("For test #%d, ", testNumber)
			continue
		}

		{
			var iterator Iterator = &internalIter

			var numIterations int
			for iterator.Next() {
				numIterations++

				renderer, err := iterator.Renderer()
				if nil != err {
					t.Errorf("For test #%d and iteration count #%d, did not expect an error, but actually got one: (%T) %v", testNumber, numIterations, err, err)
					continue Loop
				}

				actualFileName, actualWriterTo, err := renderer.WriterTo(test.RendererParams)
				if nil != err {
					t.Errorf("For test #%d and iteration count #%d, did not expect an error, but actually got one: (%T) %v", testNumber, numIterations, err, err)
					continue Loop
				}

				if _, ok := test.Expected[actualFileName]; !ok {
					t.Errorf("For test #%d and iteration count #%d, did not expect file name %q, but actually got it.", testNumber, numIterations, actualFileName)
					continue Loop
				}
				{
					var buffer bytes.Buffer

					n, err := actualWriterTo.WriteTo(&buffer)
					if nil != err {
						t.Errorf("For test #%d and iteration count #%d, did not expect an error, but actually got one: (%T) %v", testNumber, numIterations, err, err)
						continue Loop
					}
					if expected, actual := int64(len(test.Expected[actualFileName])), n; expected != actual {
						t.Errorf("For test #%d and iteration count #%d, expected %d, but actually got %d.", testNumber, numIterations, expected, actual)
						t.Errorf("EXPECTED: %q", test.Expected[actualFileName])
						t.Errorf("ACTUAL:   %q", buffer.String())
						continue Loop
					}
					if expected, actual := test.Expected[actualFileName], buffer.String(); expected != actual {
						t.Errorf("For test #%d and iteration count #%d...", testNumber, numIterations)
						t.Errorf("EXPECTED: %q", expected)
						t.Errorf("ACTUAL:   %q", actual)
						continue Loop
					}
				}
			}
			if err := iterator.Err(); nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber)
				continue
			}
			if expected, actual := len(test.Expected), numIterations; expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				continue
			}
		}
	}
}
