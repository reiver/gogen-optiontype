# gogen-optiontype

Option types (also known as maybe types) for Go programming language, via `go generate`.

# Usage

```go
package itemid

//go:generate gogen-optiontype --pkg=itemid --type=int64

```

```go
package thingid

//go:generate gogen-optiontype --pkg=thingid --type=string

```
