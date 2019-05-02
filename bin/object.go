package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
)
// 命令 test
type Test struct {
	bin.Command
}
// 项目初始化
func (a *Test) Init ()  {
	// 重写方法时必先系统父结构体方法[!!]
	a.Command.Init()

	// todo ....
}
// 运行，执行内二级命令分发
func (a *Test) Run ()  {
	fmt.Println("ffff.")
}

// 命令 yang
type Yang struct {
	bin.Command
}


func main() {
	//router := &bin.Router{}
	//bin.Register("test", &Test{})
	//bin.Register("yang", &Yang{})
	//bin.Adapter(router)
	bin.RegisterApps(map[string]interface{}{
		"test": &Test{},
		"yang": &Yang{},
	})
	bin.Run()
}
