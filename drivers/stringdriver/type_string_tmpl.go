package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "string"
	const fileName = "gen_type_string.go"

	renderer := gendriver.DefaultRenderer{
		FileName:    fileName,
		FileImports: shared.TypeStringImports,
		FileTmpl:    shared.TypeStringTmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
