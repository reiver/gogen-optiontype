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
		Pkg     string
		Trial   bool
		Type    string
	}
	{
		flag.StringVar(&flags.Pkg,  "pkg",   "main", "Package name. Ex: --pkg=main")
		flag.BoolVar(&flags.Trial,  "trial", false,  "Trial run. Ex: --trial")
		flag.StringVar(&flags.Type, "type",  "",     "Type. Ex: --type=int64")

		flag.Parse()

		if "" == flags.Type {
			fmt.Fprintf(os.Stderr, "ERROR: \"type\" may not be empty.\n")
			return
		}
	}

	{
		registry := gendriver.Registry
		if nil == registry {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Could not find registry.\n")
			return
		}

		iterator, err := registry.Iterator(struct{
			Pkg  string
			Type string
		}{
			Pkg:  flags.Pkg,
			Type: flags.Type,
		})
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Could get an iterator from the registry, because: %v\n", err)
			return
		}
		if nil == iterator {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Received a bad iterator from the registry.\n")
			return
		}


		for iterator.Next() {

			renderer, err := iterator.Renderer()
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem getting renderer, because: %v\n", err)
				return
			}

			fileName, writerTo, err := renderer. WriterTo(gendriver.RendererParams{
				Pkg:  flags.Pkg,
				Type: flags.Type,
			})
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem preparing to write file, because: %v\n", err)
				return
			}


			var buffer bytes.Buffer
			n, err := writerTo.WriteTo(&buffer)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing file %q contents to buffer, because: %v\n", fileName, err)
				return
			}

			switch flags.Trial {
			case true:
				fmt.Printf("FILE NAME: %q\n", fileName)
				fmt.Printf("FILE SIZE: %d\n", n)
				fmt.Printf("FILE CONTENTS:\n%s\n", buffer.String())
			default:
				file, err := os.Create(fileName)
				if nil != err {
					fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem create file %q, because: %v\n", fileName, err)
					return
				}

				n2, err := buffer.WriteTo(file)
				if nil != err {
					fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: %v\n", fileName, err)
					return
				}
				if expected, actual := n, n2; expected != actual {
					fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem writing contents to file %q, because: expected to write %d bytes, but actually wrote %d bytes.\n", fileName, expected, actual)
					return
				}
			}


		}
		if err := iterator.Err(); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Problem iterating through files, because: %s\n", err)
			return
		}
	}
}
