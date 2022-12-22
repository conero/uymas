# Uymas

Golang 常用包，快速实现命令行程序开发、struct合并、随机数等生成.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gitee.com/conero/uymas?tab=doc)  [![Goproxy.cn](https://goproxy.cn/stats/gitee.com/conero/uymas/badges/download-count.svg)](https://goproxy.cn)  [![](https://goreportcard.com/badge/gitee.com/uymas/conero)](https://goreportcard.com/report/gitee.com/conero/uymas)  [![Go](https://github.com/conero/uymas/actions/workflows/go.yml/badge.svg)](https://github.com/conero/uymas/actions/workflows/go.yml)



**代码仓库介绍**

- [~~github~~](https://github.com/conero/uymas) 由于网络原因取消该站点
- [gitee](https://gitee.com/conero/uymas)




### 项目介绍
go 语言工具库
go-version： *v1.11.1*

- source
    - bin    命令行解析工具



**分支介绍**

- master 项目主分支
- develop 开发分支
- nestling  雏形分支，包含实验性的代码



```powershell
# 分支合并顺序
nestling --> develop -->master
```





### 安装

```ini
# github
$ go get -u gitee.com/conero/uymas

```



#### bin

> 命令行行语法
>
> `$ [command] [<options>]`

```ini
$ [command] [<options>]
# [<options>] 格式如下


# 1. 配置参数；全称以及简写
--set[=true]
# - 表示单字符; -x; -xy => -x -y ; -xzy => x=zy
-short[=true]


# 2. 二级命令(紧接着 [command])
$ [command] [<sub-command>] [<options>]
```



#### tinygo

`Experimental/Try`  尝试实验性支持 [tinygo](https://github.com/tinygo-org/tinygo)

- [ ] **进行中**（since 2022-12-22）





- （`"reflect is not fully implemented"`）That Fprintln appears to be using reflection, which is not well supported under tinygo yet. ([E2935](https://github.com/tinygo-org/tinygo/issues/2935))



### 使用

```go
package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
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

