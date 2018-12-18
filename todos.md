# todos-list

> 2018年12月18日 星期二
>
> Joshua Conero



*状态*

- <font color="red">Need</font>   待处理
- <font color="green">Done</font>  已完成



## v0.4

> add the v0.4 need todo



### (<font color="red">Need</font>) Optimization bin easy to get action key

> 优化`bin`简洁获取命令的值

```go
// old
// $ [app] action <key>
func (a *ActionApp) Action(){
    key := a.App.Next("kv")
    fmt.println(" get the action key like:" + key)
}

// new todo
// apply a Simplified method
```

