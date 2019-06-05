## uymas/bin

> 命令行生成工具





### 格式

```shell
# 命令解析
$ [command] <option>

# 简单选项
$ <option>
```



> 数据格式

```shell
# 字符串
$ 'the data string.'	# the data string.
$ "the data string."	# the data string.

# 数字类型
$ 8						# int
$ 8.88					# float

# 数组
# 分割符号(separator)  默认","
$ 'a','b','c','d'		# array [a, b, c, d]
# "," 分割
$ --separator-comma 
$ -spt-c
```





### 教程

*路由状态分为：命令匹配成功 、空命令状态、自定义函数路由成功状态。*

*命令行程序可实现 `对象式` 和 `函数式`， 同时持：`对象式/函数式混合风格`*



#### 对象式

```go
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

```





#### 函数式

```go
package main

import (
	"github.com/conero/uymas/bin"
)


func main() {
	// 项目注册
	bin.RegisterFunc("name", func() {
		fmt.Println(" conero/uymas/bin example with Base.")
	})

	// 未知命令
	bin.UnfindFunc(func(cmd string) {
		fmt.Println(cmd + "unfind（functional）")
	})

	// 空函数
	bin.EmptyFunc(func() {
		fmt.Println("empty（functional）")
	})

	bin.Run()
}

```

