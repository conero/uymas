package main

import (
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
)

func main() {
	app := gen.AsCommand(new(defaultApp))
	lgr.ErrorIf(app.Run())
}
