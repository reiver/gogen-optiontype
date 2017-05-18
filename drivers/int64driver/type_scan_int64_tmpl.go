package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_type_scan_int64.go"

	renderer := gendriver.DefaultRenderer{
		FileIsNotTest: true,
		FileName:    fileName,
		FileImports: shared.TypeScanInt64Imports,
		FileTmpl:    shared.TypeScanInt64Tmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
