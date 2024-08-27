# Uymas

Golang 常用包，快速实现命令行程序开发、struct合并、随机数等生成.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gitee.com/conero/uymas/v2?tab=doc)   [![Report Card](https://goreportcard.com/badge/gitee.com/conero/uymas/v2)](https://goreportcard.com/report/gitee.com/conero/uymas/v2)   [![Goproxy.cn](https://goproxy.cn/stats/gitee.com/conero/uymas/v2/badges/download-count.svg)](https://goproxy.cn)  [![](https://goreportcard.com/badge/gitee.com/uymas/conero)](https://goreportcard.com/report/gitee.com/conero/uymas)  [![Go](https://github.com/conero/uymas/actions/workflows/go.yml/badge.svg)](https://github.com/conero/uymas/actions/workflows/go.yml)  [![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/conero/uymas?label=Latest%20Version&color=teal)](https://github.com/conero/uymas/releases/latest)
[![Go](https://img.shields.io/badge/go-1.20-cyan.svg)](https://golang.org) 



**代码仓库介绍**

- [~~github~~](https://github.com/conero/uymas) 由于网络原因取消该站点
- [gitee](https://gitee.com/conero/uymas)




### 项目介绍
go 语言工具库
go-version： *v1.20* （使其兼容 windows7）

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
$ go get -u gitee.com/conero/uymas/v2

```



编译程序

```shell
# 压缩化编译
go build -o ./dist -ldflags "-s -w" ./cmd/...
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



##### plugin sub command（PSC）

插件式子命令

通过扫描二进制所在目录"\$/"及“$/plg"下可执行文件，若存在将其视为PSC。

支持命名格式：

- \$/\$app-name           name
- \$/\$app_name           name
- \$/name                      name
- $/plg/\$app-name     name
- \$/plg/\$app_name     name



#### tinygo

`Experimental/Try`  尝试实验性支持 [tinygo](https://github.com/tinygo-org/tinygo)

- [ ] **进行中**（since 2022-12-22）





- （`"reflect is not fully implemented"`）That Fprintln appears to be using reflection, which is not well supported under tinygo yet. ([E2935](https://github.com/tinygo-org/tinygo/issues/2935))



### 使用

```go
package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/evolve"
)

// command struct
type test struct {
    Command
}

func main() {
	evl := evolve.NewEvolve()

    // register func
    evl.Command(func() {
        fmt.Println("Evolution For Index.")
    }, "index")

    // register struct
    evl.Command(new(test), "test", "t")
    log.Fatal(evl.Run())
}
```

