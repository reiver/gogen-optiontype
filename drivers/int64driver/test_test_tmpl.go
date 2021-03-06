package verboten

import (
	"github.com/reiver/gogen-optiontype/driver"
	"github.com/reiver/gogen-optiontype/drivers/shared"
)


func init() {
	const typeName = "int64"
	const fileName = "gen_test_test.go"

	renderer := gendriver.DefaultRenderer{
		FileName: fileName,
		FileTmpl: shared.TestTestTmpl,
	}


	if err := gendriver.Registry.Register(typeName, true, renderer); nil != err {
		panic(err)
	}
}
