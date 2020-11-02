## tcp 网络连接

> Joshua Conero
>
> 2020.11.02





#### 数据传输

>  简单的消息传递协议

```powershell
# 格式
协议://内容
```



> 客户端请求链接

```powershell
native-message://?username=${姓名}
# 认证
native-message://authorization?username=${姓名}
# 广播
native-message://broadcast?message=${信息}
# 发送到用户
native-message://send-message?username=?&&message=${信息}
```









#### 参考

- [golang_文件传输: go实现C/S构架下的文件传输系统](https://blog.csdn.net/weixin_43851310/article/details/87993308)
- [go sokect](https://www.cnblogs.com/bubu99/p/12521702.html)