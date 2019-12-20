## uymas/bin

> 命令行生成工具





### 格式

```shell
# 命令解析
$ [command] <option>

# 简单选项
$ <option>

#
# -- 与 - 的区别，参考Linux常用命令格式
--fix	# 选项全拼
-fix    # 选择简写，等同于 -f -i -x

#
# 命令行数据格式
--name='Joshua Conero'
--name 'Joshua Conero'

# only-option 作为无值选项
--only-option --last-name Conero
# 短标签映射关系（需要建立映射关系）
# -N,--name
# -O,--only-option
# 数组型参数
--persons Conero Jahn Lile Any --prex

# 实现属性严格检查开关(此时需要注册所有选择)

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





#### 解析算法实现

##### python

```python
def option_parse(args, strict_option_list=None):
    '''
    见: _example/design/option-parse/option-parse.py
    '''
    pass
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



*新式函数命令工具*

```go
FRdata
```



#### 数据加载

*支持多次数据加载*

```shell
# 大量数据加载实现
./uymas-bin.exe --load-json '{"json":"json 字符串"}' --load-json='{"json2": "方法二"}'
./uymas-bin.exe --load-url-style 'key=value&k2=v2&k3=v3'
./uymas-bin.exe --load-session-style 'key:value; k2:v2; k3:v3;'

# 不同数据加载
# 长选项
--key 'value'
--key='value'
--command-style

# 单选项
-P 'value'
-C
```



##### json

##### url-style

##### session-style

