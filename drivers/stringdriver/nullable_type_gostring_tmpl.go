package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "string"
	const fileName = "gen_nullable_type_gostring.go"

	renderer := gendriver.DefaultRenderer{
		FileIsNotTest: true,
		FileName:    fileName,
		FileImports: shared.NullableTypeGoStringStringImports,
		FileTmpl:    shared.NullableTypeGoStringStringTmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
