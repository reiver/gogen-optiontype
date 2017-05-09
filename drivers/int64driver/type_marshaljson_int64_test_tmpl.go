package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_type_marshaljson_int64_test.go"

	renderer := gendriver.DefaultRenderer{
		FileName: fileName,
		FileTmpl: shared.TypeMarshaljsonInt64TestTmpl,
	}


	if err := gendriver.Registry.Register(typeName, renderer); nil != err {
		panic(err)
	}
}
