# 更新日志
> 2018年10月30日 星期二
> 项目



**版本介绍**

- **x.y.z**     	保证的兼容性，可新增功能（用于版本阶段性开发）、修复或调整版本。（待移除使用 `// Deprecated:  descript text` 标记，或说明）
- **x.y**            不保证的兼容性，删除旧版本遗弃的方法
- **x**               重大（颠覆性）的改变，重要里程碑开发



## v1.1.x

- [ ] 内置文本使用 embed 包重写，如拼音等。

### 1.1.0/dev

**TODO**

- [ ] 命令行（**`bin`**）支持可选的脚本支持



> 引入数据接收器后二进制编译尺寸大了一倍

- all
  - optimize) 根据 IDE 提示优化备注信息，并进一步将中文切换为英文。

- **bin**
  - +) `CLI.hookBeforeCall` 新增 hook 机制，实现在调用回调前的请求事件。
  - +) `CliCmd.AppendData` 新增数据合并接口，使 CliCmd 支持外部数据导入。
  - +) `CLI.UnLoadScriptSyntax` 支持可选的（`--file,-f`） 语言集合脚本文件，（`--script,-s`）语言文本进行解析。
  - -) 移除遗弃标注项：`FormatStr`、`CliCmd.HasSetting`
  - optimize) `bin.FormatTable` 使用**reflect**扩展器支持任意2维度数组格式化。
  - **parser**
    - +) `DataReceiver` 新增数据接收器，实现对 json/url等数据格式的解析支持。
    - +) `NewScriptFile` 新增用于解析脚本的函数，初步支持 shell语法。
    - +) `ParseLine` 新增行解析，不再简单的根据空格分割。
    - **syntax**  初步增加语言语法支持，常用
      - +) `NumberOpera` 实现字符串的等式的运算
- **fs**
  - -) 移除遗弃标注项：`Struct --> ByteSize`
- **str**
  - -) 移除遗弃标注项：`Func -> UcWords`
  - -) `str.WriterToContent` 移除实验性的词方法，使用语言支持库。
  - +) `GetNotEmpty` 使用从多个字符串中获取非空字符串
- **number**
  - -) 移除遗弃标注项：`Struct -> One`
- **util**
  - -) 移除遗弃标注项：`Func -> Sec、SecCall、SecCallStr`
- **parse/xini**
  - change) `GetFunc` -> `SetFunc`
  - optimize) `Container.HasKey` 扩展支持多个键值检测

#### alpha版本

###### v1.1.0-a1/20210702



## 1.0.x

- [ ] 删除历史版本中标注遗弃的方法
- [ ] godoc 内部文档统一替换为英文
- [ ] bin
  - [x] `--fixed` 与 `-fixed` 的区别，前者指代全拼 *fixed*，后者 *`-f -i -x -e -d `无序化，两者有一个对应列表；*
  
  - [ ] `--full-name, -F` 通过设置，自动生成文档，新增一个对象用于实现。
  
  - [ ] 是否增加系统日志，来调试命令程序？
  
  - [ ] 类似与git查询相识的命令，若命令不存在
  
  - [ ] 语法解析不太准确，可使用正常先处理再分割
  
  - [ ] `--name a1 --name=a2` 处理为数组。支持数组格式
  
  - [ ] 基于os.ENV 的session机制
  
  - [x] ~~使用内置的 history~~（系统提供遂放弃）
  
    

### 1.0.2/2021-07-02

- [x] doc 文档修改为英文。

> v1.0 修复版本

- **number**
  - +)  `BitSize` 增加方法 `Bit/Byte` 以完善可自由获取不同的单位
  - optimize) `BitSize`默认输出采用 1024 制式而非原来的1000
  - +) `Unit` 增加方法 `Unit` 以完善可自由获取不同的单位
  - +) `Bytes` 方法用于字节快速转 `BitSize` 的方法
  
- **bin**
  
  - +) `StructToMap` 和 `StructToMapLStyle` 新增 struct 转 map 的方法
  - +) `CliCmd.AppendData` 新增数据合并接口，使 CliCmd 支持外部数据导入。
  - 优化) `FormatQue` 使用反射调整其支持任一类型的 array/slice数据
  
- **util**
  - +) `InQueAny` 新增方法来支持任一类型的 array/slice数据键值检测
  - +) `SendTimeDiff` 新增方法用于计算更加统计的时间消耗，用于替代旧方法：`Sec、SecCall、SecCallStr`。
  
- **str**
  - +) `IsLatinAlpha` 判断是否为拉丁文字母
  - +) `LowerStyle` 将峰坨式命名规则转为小写风格命名规范
  - +) `CamelCase`  使用 snake case（蛇式）转 CamelCase 实现

- **fs**
  - +) `GetMemStatsRange` 新增方法用于计算程序内存消耗
  - +) `MemUsage` 新增结构体用于快速计算系统的内存消耗情况

- **cmd/uymas**
  
  - optimize) `scan, sc` 增加内存使用监控、时间消耗的监控。
  
    

### 1.0.1/2021-06-11

- **cmd/uymas**
  - +) 使用 Go 1.16 的 embed 特性，添加文本式的命令文档
- **bin**
  - +) 新增方法 *GetHelpEmbed* 用于读取命令行文档格式
  - +) 新增方法 `CliCmd.Context` 获取`CLI`对象的上下文。
  - +) 新增方法 `CliCmd.CmdType` 获取系统 Cmd 类型，目前为分为`App/Func`。
  - optimize) `CliCmd.Next` 参会为空时，获取当前应用的的下一个子命令。
  - optimize) `FormatKvSort` 和 `FormatKv` 使用反射改写是函数支持任何格式的Map类型
- **number**
  - +) 新增类型 `Unit` 并增加方法 `K/W/M/G` 方法，`Unit` 用于替代 `One`。
  - +) 新增 `BitSize` 类型转换处理类，参照 `Unit` 增加子方法，并同时支持 `KiB` 与 `KB`。
- change) go 版本要求 1.16+



### 1.0.0/2021-02-19

**由于，在 0.6.0 上的开发出现非兼容方法，因此发布版本计划进行改变** (~~0.6.0/Next~~)

因版本计划过长，本次直接发布该版本，后续优化体现在 1.0.x 版本中。

*alpha 版本可为功能快照，加快功能迭代，原则上新增的功能将保留在(x.z.)版本中，若需要删除在下一版本中实现*

- **bin**
  - **App**
    - +) *`CheckSetting`* 新增 app 选项是否存其中，支持多个参数
    - +) *`CheckMustKey` 检测必须的键值是否存在*
    - +) `ArgRawLine` 返回原始输入行字符串
  - **CLI**
    - +) 新增 `CallCmd` 用于内部调用存在的命令
    - +) `Inject` 和 `GetInjection` 新增数据注入器，用于实现如 chan 信号控制等
    - +) 新增方法 `CmdExist` 用于检测命令是否存在
    - optimize) `RegisterEmpty` 和 `RegisterUnfind` 支持简化版的注册函数，即 `cc *CliCmd` 非回调函数必须
  - **CliCmd**
    - +) 新增 **CLI** 上下文`context`，使用具有调用命令的能力（`CallCmd`）
    - +) 添加方法 `ArgInt` 用于后去整形数据
    - optimize) `ArgRaw` 添加支持多参数获取单一值得能力
    - optimize) `Arg` 扩展器支持多参数与 `ArgRaw` 参数保持一致
    - optimize) `RegisterUnfind` 参数错误时提供日志级别的提醒
  - **Option**
    - +) 新增选项解析类，用于对 `args` 值得映射
  - (调整) 函数式注册方法，统一新增参数 `a *bin.App` 。 [非兼容性调整]
  - (try) 新增 Option 对象，严格控制option的输入是否正确
- **bin/buitl**
  - +) *新增 `bin util` 包，使其区分 bin 中 `init()`， 后者无该函数*
  - +) 新增 `GetBasedir()` 函数，用于获取应用运行的基础目录地址，增加缓存。
  - +) *新增函数 `StringToArgs()`, 用于将字符串安装args模式切割为数组*
  - +) 新增 `GetPathDir` 函数用于获取基于应用所在目录的地址
- **bin/parser**
  - +) 添加`bin/parser`子包专门用于实现命令行语法解析
- **io**
  - fixed) *io.StdPathName 特定下错误修复* 
- **netutil** 网络请求助手(新增)
  - **Httpu** *http util 方法集合*
- **storage**
	- +) 实验性引入内存数据存储器
- **util**
  - -)  删除 `util.BaseError` 结构体，使用系统的 `errors.New()` 代替
  - +) 添加控制判断，以及控制对比的方法 `NullDefault` 和 `ValueNull`
  - +) 新增 object 对象，初步实现了 map/struct 之间的 assign 操作
- **fs**
  - +) `fs.DirScanner` 添加排序和过滤表达式，且在添加过滤时判断表达式的有效性
- **cmd/uymas**
  - +) 新增 help 命令，以及`scan`添加排除和过滤
  - +) 新增 test 命令，用于测试命令行解析程序情况
- **xsql**
  - optimize) 数据库连接器增加对 `join` 的支持
  - fixed) 修复数据库查询为空时的参数错误。




#### todos

- [x] `bin.Command` 中 `App` 属性无时效性，需要在继承的命令中调用 `bin.GetApp()` 实现，需要优化 `runAppRouter()`



#### alpha5/next

##### todo

- [ ] 所有的方法添加对应的测试脚本，尽可能，以提高代码的可用性。



#### alpha4/200727

**culture/pinyin**

- (+) *将个人其他模块中的`pinyin`生成移动到此包中*

**fs**

- (+) 添加通用的目录扫描工具
- (+) 文件尺寸大小单位转换工具

**number**

- (+) 数字通用大小单位转换工具

**bin**

- (-) 删除旧版本中已经的方法
- (fixed) 命令基本类型或是失败

**parse/xini**

- (+) 新增ini解析库，从个人库[`github.com/conero/inigo`](https://github.com/conero/inigo)迁移过来

**cmd/uymas**

- (+) *添加包实例命令行程序*
- (+) 新增目录扫描工具





#### alpha3/200612

> **bin**

- (+) 添加新的的`bin.CLI`构建方式，函数式设计与原的命令行设计区分。前者更加适合与驻留式命令行程序，引入语言式风格。后期将删除旧的命令行形式
- **bin/parser**
  - (+) 添加`bin/parser`子包专门用于实现命令行语法解析
- **bin/buitl**
  - (+) *新增 `bin util` 包，使其区分 bin 中 `init()`， 后者无该函数*
  - (+) 新增 `GetBasedir()` 函数，用于获取应用运行的基础目录地址
  - (+) *新增函数 `StringToArgs()`, 用于将字符串安装args模式切割为数组*
  - (+) `Clear()`  采用命令行调用清屏幕失败(实验性)

> **storage**

- (+) 实验性引入内存数据存储器





#### alpha2/191230

- *删除 `.idea` IDE 文件* 
- *develop 分支中新增举例目录，在 master 分支中将会删除；也即是其只包含在开发分支中*
  - *新增 python 命令行解析方法*
- *设计文档优化*
- *删除历史版本中 `Deprecated: rename` 标签的代码*
- *bin/buitil 包新增*
- *其他代码优化*




#### alpha1/191025

- **bin**
  - **App**
    - (+) *CheckSetting* 新增 app 选项是否存其中，支持多个参数



## 0.5.x

> **重新定义命令参数解析规则**

### 0.5.2/20191008

- **bin**
  - (+) *新增方法 `FormatKv` 方法用于优化原` FormatStr`，后者标注遗弃状态*
  - **App**
    - (+) *新增代码默认的 arg 参数获取方法, `app.ArgDefault`、`app.ArgRawDefault`*
    - (重命名) *原方法保留，下一个版本中删除。 `app.Args -> app.Arg`, `app.ArgsRaw -> app.ArgRaw`*
- **其他**
  - *项目中代码优化，如添加 [Deprecated]标签*



### 0.5.1/20190827

- **bin**
  - **App**
    - (+) *新增方法 `resetQueue`， 用于重置 app.Queue 队列数据信息*
  - (优化) *Run 启动命令行时 `runAppRouter` 会重置 app.Queue 队列的数据*
  - (更名) *`getArgs()` 更名为 `Args()` 便于外部测试*
  - (优化) *`Init()` 文档添加【deprecated】，表明以后即将删除*
- **fs**
  - (+) *新增方法 `StdPathName` 获取标准路径信息* 
  - (+) *新增方法 `Put` 用于文件写入且覆盖文件原内容*
  - (+) *新增方法 base64 的加解密, `base64_encode/base64_decode`*
  - (优化) *`StdDir` 内部使用 StdPathName 来获取标准信息*
- **str**
  - (修复) *Url.AbsHref 绝对地址生成时端口号出现重复的问题*
  - (+) *新增对象 `RandString` ：实现随机字符串的生成*



### 0.5.0/20190605

- **uymas/bin**
  - (+) *添加方法`FormatQue` 实现切片格式化字符串生成*
  - (+) *添加方法`FormatTable`实现Table类型数据字符串生成*
  - (+) *`bin_test.go` 添加对`FormatQue` 和`FormatTable`的测试*
  - (+) *CmdUitl*
    - 添加对象为 `bin.Command` 提供工具方法
    - *实现统一的方法`BaseSubCAlias`，用于二级命令的别名解析，包含了框架统一的`-h`命令详情*
  - (+) 新增函数式`RegisterFunc` 路由定义，实现轻量级的命令行程序库
  - (+) *函数式命令实现*
    - _添加 *EmptyFunc* 方式使用函数式空接口_
    - _添加 *UnfindFunc* 方式使用函数式未知命令接口_
  - (+) *新增方法 `InjectArgs` 用于开发是 IDE 的测试，通过注入改变内部解析参数*
  - (+) *新增方法 `StrParseData` 使用命令模式下字符串参数的格式解析*
  - (修复) *Bin.GetApp() 等返回 App 类型的实效性，采用返回引用地址*
  - **App**
    - (修复) *app.Data 属性解析错误*
    - (优化) *app.Data 的参数为解析的值，与 app.DataRaw 保持对应的差异性*
    - (+) *新增属性 `DataRaw`， 默认字符串类型*
    - (+) *新增 args 参数的获取方法，即分别从`Data` 和 `DataRaw`获取值。对应方法: `Args` 和 `ArgsRaw`*
  - **Router**
    - (+) *添加函数式 action 命令行实现方式*
  - (优化) *新建文件 format 用于放置格式相关的函数，并将 bin 的对应的函数移入*
- **uymas/fs**
  - (+) *新增方法 `ExistPath` 用于检测文件/目录的存在性*
  - (+) *`Append` 实现文件尾部附加写入*
- **uymas/number**
  - (+) *添加包用于数值运算*
  - (+) *实现`SumQueue`方法用于`interface{}`类型的集合求和*
  - (+) 实现`SumQInt`对int类型集合的求和
- **uymas/util**
  - (+) *新增回调时间运行计数器：`SecCallStr/SecCall`，分别返回数据/字符串*
- **uymas/str**
  - (+) *新增对象 str.Url 并实现方法 `AbsHref` 用于相对地址到绝对地址(实际URL)的抓换；以及 godoc 编写例子和测试用例*
- **其他**
  - (优化) *优化代码，使用尽可能符合 godoc规范*
  - (优化) *godoc 规范优化，所有 package 添加文字说明描述*





## 0.4.x

_<font color="blue">此版本开始，必须在发布大版本时，总结该大版本的具体更变.</font>_

> - 添加 **uymas/svn** 包：svn 命令解析程序
> - 添加**uymas/unit**包：test 单元测试扩展包
> - **uymas/bin**：
>   - 添加对二级命令 “all-key => AllKey” 的映射支持



### 0.4.2/20190502

- **bin**
  - (+) *RegisterFunc 函数式命令解析方法，以及内部注册字典支持*
  - **Router**
    - (+) *属性 `FuncAction` 自定义命令函数方法支持*



### 0.4.1/20181218

- **uymas/svn**
  - (修复) *`XmlLog` 的xml结构标注无效，即`Bridge.Log` 正常*

- (+) **uymas/unit**
  - *`testing`单元测试相关的协助程序包*
  - *实现`StrSingLine` 字符串单通道测试控制，支持自定义测试*

- **uymas/util**
  - (+) *Decimal*
    - *添加`十进制`整数型处理*
    - 实现`十进制转N进制`，以及提供2,8,16,32,36,62等进制的快捷转换
  - 添加快捷方法`DecT36/62`实现对应的机制转换
  - (+) *添加 util_test 单元测试脚本*

- **uymas/str**

  - (+) *添加 `Reverse` 用于翻转字符串，以及添加对其的测试*

- **uymas/bin**

  - (+) *添加方法 `AmendSubC`用于修正二级命令* 
  - (+) *添加`runRouter_test.go` 的测试脚本*
  - (+) *添加`FormatStr` 方法实现命令程序字符串格式化*
  - (优化) *`runAppRouter` 添加对二级命令 “all-key => AllKey” 的映射支持*




### 0.4.0/20181206

- *uymas/svn*
  - (+) *添加 svn 包，实现svn 子命令（二次封装）*
  - (+) *添加方法 `Version` 实现获取svn的命令*
  - (+) *架构`Bridge` 与 svn 命令的桥接口处理*
    - *实现该结构体*
    - (+) *添加方法 `Log` 实现 对` svn log --xml` 的解析；添加与之有关的结构体`XmlInfo`用于解析xml文档*
    - (+) *添加方法 `Info` 实现 对` svn info --xml` 的解析；添加与之相关的结构体`XmlLog `以解析xml文档*
    - (+) *添加测试文件`bridge_test.go` 测试对应的基本方法：Log/Info*
  - (+) *添加方法`Call`实现对操作系统命令行调用svn*
  - (+) *添加`svn_test.go` 测试文件*
- *uymas/bin*
  - **Command**
    - (+) *添加方法`Command.Help` 用于实现默认的帮助文档*
- *其他*
  - (优化) *文档更新，补全历史就的日志内容*
  - (+) *添加仓库文档 doc.md*



## 0.3.x

> - 添加包 _**uymas/str**_        *重命名程序包*

### 0.3.1/20181205

- *uymas/fs*
  - (优化) *`CheckDir` 方法添加返回标准的目录格式*
  - (+) *添加方法 `CopyDir` 实现对目录全复制*
  - (+) *添加方法 `StdDir` 获取标准的目录*
- *uymas/str*
  - (+) *`SplitSafe` 方法用于安全分割字符串为切片*
  - (+) *`ClearSpace` 方法用于清洗字符串中的空格* 
  - (+) `str_test`
    - *添加 str 测试文件，并实现对 SplitSafe/ ClearSpace 的测试*
    - *添加`TestRender` 单元测试*
  - (+) `WriterToContent`
    - *添加类型 `WriterToContent` 用于实现`io.Write` 接口已返回字符串数据*
    - *添加方法 `Content`获取字符串内容*
  - (+) *新增方法`Render` 用于编译 tpl 模板，并返回内容*
- *uymas/bin*
  - *App*
    - (+) *添加方法 `Next` 方法用于下个输入的队列*




### 0.3.0/20181107

> 非兼容调整： *uymas/util* 包中的字符串集合处理转移到 *uymas/util/str 中；以及名字调整。更新由此引起的代码更变*
>
> (+) *uymas/bin 实现参数监听*

- *uymas/util/str*  -> (更名为) -> *uymas/str*
  - (调整) *将 uymas/util中的代码转移至 uymas/str 中*
  - (+) *添加方法 `DelQue` 实现删除字符串中值*
- *uyams/bin*
  - (+) *添加方法 `isVaildCmd(c string) bool` 用于判断是否非法*
  - (优化) *方法`runAppRouter() ` 支持参数监听*
  - **Router**
    - (+) *属性添加`OptionListener` 实现属性监听*



## 0.2.x

> - 添加包 _**uymas/fs**_        *文件系统工具*
> -  _**uymas/bin**_           _程序优化/实现_

### v0.2.1/20181107

> (优化) *uymas/util/str* 实现首字母小写，`Ucfirst 与 Lcfirst 相对应`
>
> (实现) *uymas/bin 实现跨 cmd 请求帮助方法*

- *uymas/bin*
  - (+) *添加方法`CallCmdHelp(key string) bool` 实现框cmd方法帮助方法*
  - (+) *添加方法`Rwd() string` 获取命令行程序所在目录
- *uymas/util*
  - (+) *添加基本的 Error 实现类*
- *uymas/util/str*
  - (+) *添加方法 `Lcfirst` 实现首字母小写*
- *uymas/fs*
  - (+) *添加方法 `CheckDir` 实现自动检测目录存在性不存在则尝试创建*
  - (+) *添加方法`IsDir(dir string)` 检测目录是否存在*



### v0.2.0/20181106

> (优化) *uymas/bin* 实现二级命令分发
>
> (+) 添加 *uymas/fs* 包 

- *uymas/bin*

  - (修复) *`runAppRouter` 中解析 Setting 属性无效*
  - (+) *添加结构体 `SubCmdAlias`*
  - *Command 结构体*
    - (+) *添加SubCmdAlias属性*
    - (+) *`InnerDistribute` 方法实现入口分发(二级命令分发)*
    - (优化) *Init 方法内部支持二级命令别名路由；函数命令有由(comand \*Command) -> (c \*Comand)*

- *uymas/util*

  - (+) *添加方法 `InstruQuei` 实现大小写不敏感的检测*
  - (优化) *从其他就的代码中迁移运行时间花费*

- *uymas/fs*
  - (+) *新增程序包，用于实现文件处理*
  - (+) *添加基于 `io` 读写接口的实现*



## 0.1.x

> 项目基本搭建
>
> - 添加包 _**uymas/bin**_        *命令行协助程序*
> - 添加包 _**uymas/util**_        *项目工具包*
> - 添加包 _**uymas/util/str**_  *字符串工具包*

### v0.1.1/20181105

> 项目优化
>
> (+) *uymas/bin* 支持别名命令行；支持二级属性参数(默认打开)

- *uymas/bin*
  - (+) *`Alias(cmd string, alias ...string)` 方法用于支持别名命令行*
  - (+) *`AliasMany(alias map[string][]string)` 别名批量设置法*
  - (+) *`getCommandByAlias(command string) string` 实现对别名命令解析的支持*
  - (+) *`SubCommand(able bool)` 实现对是否禁止二级命令，默认开启*
  - *App*
    - (+) *`QueueNext(key string) string ` 获取队列右邻值*
    - (优化) *App* 结构体属性优化，删除 Option(更名为 Setting)，添加二级命令以及其他属性
    - (优化) *HasOptions -> HasSetting 配置存在性检测*



### v0.1.0/20181030

> 项目初始化

- *uymas/bin*
  - **router**
    - (+) 实现命令行程序路由，使用 *reflect* 放射机制，通过注册应用实现App路由
    - (+) 初步实现对cmd应用的解析
  - **App**
    - (+) 实现 App 类，提供命令处理的基础方法
  - **Command**
    - (+) 实现Command基类，用于实际应用继承；对命令行程序入口管理
- **uymas/util**
  - (+) *提供切片存在性判断*
- **uymas/util/str**
  - (+) 实现方法 *Ucfirst* 用于对首字母变大写