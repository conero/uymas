package main

import (
	"gitee.com/conero/uymas/v2/cli"
)

// zero 代码测试
//
// 压缩模式执行（删除调试信息/可用于生产）： go build -ldflags "-s -w" ./example/cli/zero/
// 普通打包： go build ./example/cli/zero/
func main() {
	app := cli.NewCli()
	_ = app.Run()
}
