package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_nullable_type_string1.go"

	renderer := gendriver.DefaultRenderer{
		FileName:    fileName,
		FileImports: shared.NullableTypeString1Imports,
		FileTmpl:    shared.NullableTypeString1Tmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
