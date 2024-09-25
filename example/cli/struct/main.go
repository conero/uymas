package main

import (
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/logger/lgr"
)

func main() {
	app := evolve.FromStruct(new(defaultApp))
	lgr.ErrorIf(app.Run())
}
