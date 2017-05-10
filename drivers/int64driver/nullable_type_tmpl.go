package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_nullable_type_tmpl.go"

	renderer := gendriver.DefaultRenderer{
		FileName: fileName,
		FileTmpl: shared.NullableTypeTmpl,
	}


	if err := gendriver.Registry.Register(typeName, renderer); nil != err {
		panic(err)
	}
}
