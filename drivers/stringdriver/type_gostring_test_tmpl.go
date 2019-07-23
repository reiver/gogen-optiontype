package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "string"
	const fileName = "gen_type_gostring_test.go"

	renderer := gendriver.DefaultRenderer{
		FileName: fileName,
		FileTmpl: shared.TypeGoStringStringTestTmpl,
	}


	if err := gendriver.Registry.Register(typeName, true, renderer); nil != err {
		panic(err)
	}
}
