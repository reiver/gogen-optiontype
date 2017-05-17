package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "string"
	const fileName = "gen_type_scan_string.go"

	renderer := gendriver.DefaultRenderer{
		FileName:    fileName,
		FileImports: shared.TypeScanStringImports,
		FileTmpl:    shared.TypeScanStringTmpl,
	}


	if err := gendriver.Registry.Register(typeName, false, renderer); nil != err {
		panic(err)
	}
}
