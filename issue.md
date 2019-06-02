# Issue

> 2019-6-2 15:04:56



> 状态标记

- <span style="color: red;">NeedToDo</span>   未完成
- <span style="color: green;">Done</span>   已完成



### I00120190602 bin 函数式接口中 `bin.GetApp()` 的实效性 <span style="color: green;">Done</span>

```go
app := bin.GetApp()

bin.RegisterFunc("config", func() {
    app2 := bin.GetApp()
    // app 与 app2 的可能存在失效， app2 为最新对象
    // ...
})

// ...

// 启动命令行函数
bin.Run()
```

