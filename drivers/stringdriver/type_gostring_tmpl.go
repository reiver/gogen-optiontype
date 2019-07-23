package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "string"
	const fileName = "gen_type_gostring.go"

	renderer := gendriver.DefaultRenderer{
		FileIsNotTest: true,
		FileName:    fileName,
		FileImports: shared.TypeGoStringStringImports,
		FileTmpl:    shared.TypeGoStringStringTmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
