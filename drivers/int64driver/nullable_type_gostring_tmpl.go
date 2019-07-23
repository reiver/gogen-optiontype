package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_nullable_type_gostring.go"

	renderer := gendriver.DefaultRenderer{
		FileIsNotTest: true,
		FileName:    fileName,
		FileImports: shared.NullableTypeGoStringInt64Imports,
		FileTmpl:    shared.NullableTypeGoStringInt64Tmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
