# ChangeLog(v2)

> Joshua Conero
>
> 20180819

## V2.1

### V2.1.0 - alpha

> todos

- [ ] 实现 *[toml](https://github.com/toml-lang/toml)* 基本解析(支持)
- [ ] 代码注释英文化，即便是开始比较困难;以及git commit尽量英文化



> 基本概述

- *搭建`toml`解析器的支持程序*

- (优化) *根据 godoc 规范优化注释文本*

- (优化) *更名 `container -> Container` 便于测试，使之设计得更加合理*



**inigo**

- (+) *`ParseValue` 方法按照设定的规则实现字符串字面值的解析*
- (+) *`StrClear` 方法实现字符串字面变量值得清洗*
- [修复] *inigo.NewParser 参数解析错误*
- **Container**
  - (+) *新方法 `GetDef` 使用简介的带默认参数的值获取*
  - (+) *方法`GetFunc`: 内部添加事件驱动时的动态值获取，即获取值时内部调用注册回调函数*
- (+) *新增方法 `Del` 删除容器中的键值*
  - (+) *新增方法 `Merge` 合并容器数据*
  - (+) *新增方法 `Reset` 重置中容器的数据*
- **Parser**
  - (+) *添加方法 `Driver() string` 用于获取当前的驱动名称；以及实现各个解析器的对应的方法*
  - (+) *添加与 Container 对应的 GetDef 方法*
  - (+) *新增方法 `ErrorMsg` 用以返回错误信息*
  - (+) *新增方法 `GetFunc(key string, regFn func() interface{}) Parser` 实现事件驱动式值获取*
  - (+) *新增方法 `Del` 用于删除配置中的键值*
  - (实现) *实现 `Raw` 方法获取原始的数据，可用于数据测试以及检测*
  - (优化) *调整 `Container` 更改的数据进行优化 (20190606)*
  - (修复) *内部参数 valid 无效的问题*





## v2.0

### v2.0.12/181222 - alpha

> Package

- `inigo_test.go` 测试 `TestNewParser` 用于golang的标准进行优化测试代码

- *文件重命令；使之更加容易编排查看*
  - `baseParser` -> `parserBase` 
  - `rongParser` -> `parserRong` 



### v2.0.11/181105

*初步实现，通过数据生成 ini 配置文件*

> **package**

`Parser`

- (+) 添加方法 `Save() bool` 会覆盖源文件
- (+) 添加方法 `SaveAsFile(filename string) bool` 使用当前多去的数据，生成新的文件

`BaseParser`

- (+) 实现`Parser` 中新增的方法



### v2.0.10/180930

> 设计调整

  ```
  设计: 
  	1. BaseParser -> container		(继承)
  				  -> Parser			(实现)
  				  
  	2. RongParser -> BaseParser     (继承)
  ```

> **package**


- (调整) 将旧版中 *LnReader* 移到新版中

- (优化) *NewParser* 函数的重写

- (+) *新增`container` 抽象容器·，实现对容器中数据的获取以及设置*

- `Parser`


    - (+) *添加方法Section，用于获取有关section的参数*

- `BaseParser`
  - (+) *实现对基本ini文件语法的支持，完成对文件的解析，并且获取到数据*

- `baseFileParse`


    - (+) 添加 *base-ini* 文件的读取与解析




### v2.0.1/180819

- **概述**
  - 删除历史旧代码
  - 不在区分子包， *全部包含在项目下*
- **package**
  - (+) 添加 *parser.go* 设计 ``Parser`` 解析器接口
    - (+) 实现 ``BaseParser`` 基本 ini 文件解析器
  - (+) 添加 StrParser接口，用于字符串的简单解析
  - (+) 搭建``RongParser`` 解析器

### v2.0.0/180819

- 重设计项目架构，优化 git 管理工具
- 删除``v1`` 中于*package* 无关的代码以及文档，由 *ini-go* 更名为 *inigo*
- 保证项目无错误

