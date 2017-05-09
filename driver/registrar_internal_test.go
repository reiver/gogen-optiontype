package gendriver

import (
	"bytes"

	"testing"
)

func TestInternalRegistrarRegister(t *testing.T) {

	tests := []struct{
		Data []struct{
			TypeName string
			FileName string
			FileTmpl string
			Expected string
		}
	}{
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
					Expected: `package thusid`+"\n\n"+`// This is a test for int64`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
					Expected: `package thusid`+"\n\n"+`// This is a test for int64`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
					Expected: `package thusid`+"\n\n"+`// This is a test for int64`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},



		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "string",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "string",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "string",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "string",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},



		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "string",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "float64",
					FileName: "something.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},
		{
			Data: []struct{
				TypeName string
				FileName string
				FileTmpl string
				Expected string
			}{
				{
					TypeName: "int64",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "int64",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "string",
					FileName: "apple.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "banana.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "cherry.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "string",
					FileName: "date.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},


				{
					TypeName: "float64",
					FileName: "something.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
				{
					TypeName: "float64",
					FileName: "something_test.go",
					FileTmpl: `package {{.Pkg}}`+"\n\n"+`// This is a test for {{.Type}}`,
				},
			},
		},
	}


	for testNumber, test := range tests {

		registry := new(internalRegistrar)
		{
			m := registry.m

			if expected, actual := 0, len(m); expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				continue
			}
		}

		{
			var registrar Registrar = registry

			for datumNumber, datum := range test.Data {

				var renderer Renderer = DefaultRenderer{
					FileName: datum.FileName,
					FileTmpl: datum.FileTmpl,
				}

				err := registrar.Register(datum.TypeName, renderer)
				if nil != err {
					t.Errorf("For test #%d and datum #%d, did not expect an error, but actually got one: (%T) %v", testNumber, datumNumber, err, err)
					continue
				}
			}
		}
		{
			expected := map[string][]string{}

			for _, datum := range test.Data {
				if _, ok := expected[datum.TypeName]; !ok {
					expected[datum.TypeName] = []string{}
				}

				expected[datum.TypeName] = append(expected[datum.TypeName], datum.FileName)
			}

			ExpectedLoop: for typeName, fileNames := range expected {
				var registrar Registrar = registry

				pkgName := "thusid"

				iterator, err := registrar.Iterator(IteratorParams{
					Pkg:  pkgName,
					Type: typeName,
				})
				if nil != err {
					t.Errorf("For test #%d and type name %q, did not expect an error, but actually got one: (%T) %v", testNumber, typeName, err, err)
					continue
				}

				var numIterations int
				for iterator.Next() {
					numIterations++

					renderer, err := iterator.Renderer()
					if nil != err {
						t.Errorf("For test #%d and type name %q and iteration count #%d, did not expect an error, but actually got one: (%T) %v", testNumber, typeName, numIterations, err, err)
						continue ExpectedLoop
					}

					actualFileName, actualWriterTo, err := renderer.WriterTo(RendererParams{
						Pkg:  pkgName,
						Type: typeName,
					})
					if nil != err {
						t.Errorf("For test #%d and type name %q and iteration count #%d, did not expect an error, but actually got one: (%T) %v", testNumber, typeName, numIterations, err, err)
						continue ExpectedLoop
					}

					{
						found := false
						for _, fileName := range fileNames {
							if fileName == actualFileName {
								found = true
							}
						}
						if !found {
							t.Errorf("For test #%d and type name %q and iteration count #%d, expected to find %q, but actually didn't.", testNumber, typeName, numIterations, actualFileName)
							continue ExpectedLoop
						}
					}
					{
						var buffer bytes.Buffer

						_, err := actualWriterTo.WriteTo(&buffer)
						if nil != err {
							t.Errorf("For test #%d and type name %q and iteration count #%d, did not expect an error, but actually got one: (%T) %v", testNumber, typeName, numIterations, err, err)
							continue ExpectedLoop
						}

//@TODO: confirm value of `n` and `buffer.String()`
					}
				}
				if err := iterator.Err(); nil != err {
					t.Errorf("For test #%d and type name %q, did not expect an error, but actually got one: (%T) %v", testNumber, typeName, err, err)
					continue
				}
				if expected, actual := len(fileNames), numIterations; expected != actual {
					t.Errorf("For test #%d and type name %q, expected %d, but actually got %d.", testNumber, typeName, expected, actual)
					continue
				}
			}

		}


	}
}
