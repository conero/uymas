# 帮助文档

> 2018年12月6日 星期四
>
> Joshua Conero



## `uymas/bin`

*bin.Command 模板*

```go
// 应用格式模板
type AppFormatTpl struct {
	bin.Command
}

// 项目初始化
func (a *AppFormatTpl) Init() {
	a.SCA = &bin.SubCmdAlias{
		Alias: map[string][]string{
			"list": []string{"l"},
		},
		Self: a,
	}
	a.Command.Init()
}

// 运行入口
func (a *AppFormatTpl) Run() {
	a.InnerDistribute()
	if !a.SCA.Matched {
		fmt.Println(" 日志管理应用. Need to do")
	}
}


// 帮助
func (a *AppFormatTpl) Help() {
	txt := ` $ log 项目日志管理` +
		Br + `  list,-l                      显示日志列表` +
		Br
	fmt.Println(txt)
}


```

