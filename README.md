# uymas

## 项目介绍
go 语言工具库
go-version： *v1.11.1*

- source
    - bin    命令行解析工具

## 安装

```ini
# github
$ go get -u  github.com/conero/uymas

```



### bin

> 命令行行语法
>
> `$ [command] [<options>]`

```ini
$ [command] [<options>]
# [<options>] 格式如下

# 1. 配置参数；全称以及简写
--set[=true]
-short[=true]

# 2. 二级命令(紧接着 [command])
$ [command] [<sub-command>] [<options>]
```







## 使用

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

func (test *Test) Run ()  {
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

