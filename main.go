package main

import (
	"github.com/reiver/gogen-optiontype/driver"
	_ "github.com/reiver/gogen-optiontype/drivers"

	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {

	var flags struct {
		NoTests bool
		OneFile bool
		Pkg     string
		Trial   bool
		Type    string
	}
	{
		flag.BoolVar(&flags.NoTests, "no-tests", false,  "No tests generated. Ex: --no-tests")
		flag.BoolVar(&flags.OneFile, "one-file", false,  "One file is generated (not counting tests). Ex: --one-file")
		flag.StringVar(&flags.Pkg,   "pkg",      "main", "Package name. Ex: --pkg=main or -pkg=accountid or --pkg=itemid")
		flag.BoolVar(&flags.Trial,   "trial",    false,  "Trial run. Ex: --trial")
		flag.StringVar(&flags.Type,  "type",     "",     "Type. Ex: --type=int64 or --type=string")

		flag.Parse()

		if "" == flags.Type {
			fmt.Fprintf(os.Stderr, "ERROR: \"type\" may not be empty.\n")
			return
		}
	}

	var iterator gendriver.Iterator
	{
		registry := gendriver.Registry
		if nil == registry {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Could not find registry.\n")
			return
		}

		var err error

		iterator, err = registry.Iterator(struct{
			NoTests bool
			Pkg     string
			Type    string
		}{
			NoTests: flags.NoTests,
			Pkg:     flags.Pkg,
			Type:    flags.Type,
		})
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Could get an iterator from the registry, because: %v\n", err)
			return
		}
		if nil == iterator {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Received a bad iterator from the registry.\n")
			return
		}
	}

	{
		oneFileImports := map[string]string{}
		var oneFileBuffer bytes.Buffer
		for iterator.Next() {

			renderer, err := iterator.Renderer()
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem getting renderer, because: %v\n", err)
				return
			}

			fileName, writerTo, err := renderer. WriterTo(gendriver.RendererParams{
				NoHeader: flags.OneFile,
				Pkg:      flags.Pkg,
				Type:     flags.Type,
			})
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem preparing to write file, because: %v\n", err)
				return
			}


			var buffer bytes.Buffer
			if flags.Trial || !flags.OneFile || renderer.IsTest() {
				gendriver.WriteHeader(&buffer, renderer.Imports(), struct{
					Pkg  string
					Type string
				}{
					Pkg:  flags.Pkg,
					Type: flags.Type,
				})
			}
			n0 := buffer.Len()
			n, err := writerTo.WriteTo(&buffer)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing file %q contents to buffer, because: %v\n", fileName, err)
				return
			}
			n += int64(n0)

			switch {
			case flags.Trial:
				fmt.Printf("FILE NAME: %q\n", fileName)
				fmt.Printf("FILE SIZE: %d\n", n)
				fmt.Printf("FILE CONTENTS:\n%s\n", buffer.String())
			case flags.OneFile && !renderer.IsTest():
				n2, err := buffer.WriteTo(&oneFileBuffer)
				if nil != err {
					fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: %v\n", fileName, err)
					return
				}
				if expected, actual := n, n2; expected != actual {
					fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: expected to write %d bytes, but actually wrote %d bytes.\n", fileName, expected, actual)
					return
				}

				for path, name := range renderer.Imports() {
					//@TODO: What happens if the same import is done in different ways?
					//       Ex:
					//       	import "apple/banana/cherry"
					//       vs:
					//       	import _ "apple/banana/cherry"
					//       vs:
					//       	import . "apple/banana/cherry"
					//       vs:
					//       	import food "apple/banana/cherry"
					oneFileImports[path] = name
				}
			default:
				func(){
					file, err := os.Create(fileName)
					if nil != err {
						fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem creating file %q, because: %v\n", fileName, err)
						return
					}
					defer file.Close()

					n2, err := buffer.WriteTo(file)
					if nil != err {
						fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: %v\n", fileName, err)
						return
					}
					if expected, actual := n, n2; expected != actual {
						fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: expected to write %d bytes, but actually wrote %d bytes.\n", fileName, expected, actual)
						return
					}
				}()
			}
		}
		if err := iterator.Err(); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem iterating through files, because: %s\n", err)
			return
		}
		if flags.OneFile {
			const fileName = "gen_optiontype.go"

			file, err := os.Create(fileName)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem creating file %q, because: %v\n", fileName, err)
				return
			}
			defer file.Close()

			gendriver.WriteHeader(file, oneFileImports, struct{
				Pkg  string
				Type string
			}{
				Pkg:  flags.Pkg,
				Type: flags.Type,
			})

			nExpected := int64(oneFileBuffer.Len())

			n, err := oneFileBuffer.WriteTo(file)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: %v\n", fileName, err)
				return
			}
			if expected, actual := nExpected, n; expected != actual {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: expected to write %d bytes, but actually wrote %d bytes.\n", fileName, expected, actual)
				return
			}
		}
	}
}
