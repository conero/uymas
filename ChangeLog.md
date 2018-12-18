# 更新日志
> 2018年10月30日 星期二
> 项目

## 0.4.x

_<font color="blue">此版本开始，必须在发布大版本时，总结该大版本的具体更变.</font>_

> - 添加 **uymas/svn** 包：svn 命令解析程序
> - 添加**uymas/unit**包：test 单元测试扩展包
> - **uymas/bin**：
>   - 添加对二级命令 “all-key => AllKey” 的映射支持



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