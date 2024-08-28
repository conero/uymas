# Uymas

Golang 常用包，快速实现命令行程序开发、struct合并、随机数等生成.

[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/conero/uymas?label=Latest%20Version&color=teal)](https://github.com/conero/uymas/releases/latest)
[![Go](https://img.shields.io/badge/go-1.18-cyan.svg)](https://golang.org)  [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gitee.com/conero/uymas?tab=doc) [![Report Card](https://goreportcard.com/badge/gitee.com/conero/uymas)](https://goreportcard.com/report/gitee.com/conero/uymas)   [![Goproxy.cn](https://goproxy.cn/stats/gitee.com/conero/uymas/badges/download-count.svg)](https://goproxy.cn)  [![](https://goreportcard.com/badge/gitee.com/uymas/conero)](https://goreportcard.com/report/gitee.com/conero/uymas)  [![Go](https://github.com/conero/uymas/actions/workflows/go.yml/badge.svg)](https://github.com/conero/uymas/actions/workflows/go.yml) 



**代码仓库介绍**

- [~~github~~](https://github.com/conero/uymas) 由于网络原因取消该站点
- [gitee](https://gitee.com/conero/uymas)




### 项目介绍
go 语言工具库
go-version： *v1.18*

- source
    - bin    命令行解析工具



支持 *[golangci-lint](https://github.com/golangci/golangci-lint)* 推荐规范

```shell
# 执行规范推荐
golangci-lint run ./...

# 执行所有测试用例
go test ./...
```



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



内置环境变量

```ini
# logger/lgr 包，设置日志级别：
UYMAS_LGR_LEVEL = info

# uymas 命令
UYMAS_CMD_UYMAS_LONG = true
UYMAS_CMD_UYMAS_COLON = false
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

`Experimental/Try`  尝试实验性支持 [tinygo](https://github.com/tinygo-org/tinygo)  支持版本不低于 `v0.31.0`



如编译：

```shell
# 编译 tiny 包
tinygo build ./cmd/tiny
```



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

