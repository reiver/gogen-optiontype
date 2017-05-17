package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_nullable_type_scan_int64.go"

	renderer := gendriver.DefaultRenderer{
		FileName:    fileName,
		FileImports: shared.NullableTypeScanInt64Imports,
		FileTmpl:    shared.NullableTypeScanInt64Tmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
