package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_type_gostring.go"

	renderer := gendriver.DefaultRenderer{
		FileIsNotTest: true,
		FileName:    fileName,
		FileImports: shared.TypeGoStringInt64Imports,
		FileTmpl:    shared.TypeGoStringInt64Tmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
