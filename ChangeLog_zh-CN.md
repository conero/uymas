# 更新日志
> 2018年10月30日 星期二
> 项目



**版本介绍**

- **x.y.z**     	保证的兼容性，可新增功能（用于版本阶段性开发）、修复或调整版本。（待移除使用 `// Deprecated:  descript text` 标记，或说明）
- **x.y**            不保证的兼容性，删除旧版本遗弃的方法
- **x**               重大（颠覆性）的改变，重要里程碑开发



## future

> v2.0.0 开发实现中




### v2.0.x/dev

> v2.0




#### todo

- [x] v2.0 为旧版本v1.x的重大重构版本，包结构等进行重
- [x] doc 自动文档设计，计划在 v2.0.0-alpha.3 中实现 `#240824`
- [ ] `#240915` cli.Option 执行 map 类型，gen 亦支持 map 解析
- [ ] `#240924` cmd/uymas2 正式发布后应更名，使其不与 v1 版的 uymas 混淆
- [ ] `#241019` 系统常量或函数等加入关键的文档注释，使其不忘记具体作用或来源原有
- [ ] `#241019` 包 `cli/gen` *ArgsTagData* 等一些列常量命名与时间额的功能含义比较混乱，需对命令进行规范化



### v2.0.0-rc.3/dev

> 主要测试并优化当前以实现的功能，保证功能的有效性

- **parser/xini**
  - pref: （内部）函数 `parseNumber` 修复 ParseFloat 等数制错误
- **cli**
  - feat: cli.Option 新增属性`StructGen`，`StructItems` 并支持选项struct文档标识
  - feat: 新增函数 `ConfigWith` 及 `ConfigSet` 用于读取或设置全局变量
- **cli/gen**
  - feat: 新增函数`MultiArgsMap`用于实现`cli.ArgsParser`到map的值映射
  - pref: 完善 `MultiArgs` 使其赋值兼容性更强
  - pref: 完善对 `ArgsDress` 及 `ArgsDecomposeMust` 对 `StructGen`，`StructItems` 解析支持
- **fs/os**
  - feat: 新增函数 `StdPathList` 实现路径列表标准化处理
  - feat: 实现空函数 `DelEnvPath`，删除指定的路径
  - pref: 函数 `AddEnvPath`等返回参数有 error调整为 string，因实际无法改变环境变量路径
- **util**
  - pref: 新增对象 Object 的对象实例`Obj`，使外部可以直接通过其调用方法

- **util/cloud**
  - feat: 新增 `PortAvailable` /`PortAddress` 放到实现端口可用性检查以及端口转标准地址
  - feat: 新增方法 `PortAsWeb` 用于实现端口转 http地址
  - fix: 修复package表名错误，从v1.4版本移植时
- **example/cli/fast/**
  - feat: 新增命令 port 用于检查当前可用端口




### v2.0.0-rc.2/2024-10-01

> uymas2 二进制程序实现，pinyin相关函数初步与 v1.4.1 进行合并

- **cli**
  - feat: option 新增 isdata cmd 标签用于标注option中的输入数据（非选项）
  - feat: option 新增 `DataOption` 用于标记请求数据
  - feat: Application 新增方法 `RouterBefore(t Fn)`，用于实现路由前的 hook
  - feat: option 新增 Mark 属性用于选项帮助信息生成时值占位符
  - feat: option 新增 FieldName 标识用于记录反射设置值的struct属性名称
  - pref: `Fn` 标注指定名称便于IDE自动生成
- **cli/gen**
  - feat: 新增方法 `StructDress` 实现对struct属性值装扮，支持组合/继承值
  - feat: 新增方法 `ParseStruct` ，`AsCommand` 用于解析 struct 为命令行程序
  - feat: 新增方法 `MultiArgs` 实现对数组的多重解析
  - pref:方法 `ArgsDress` 支持继承的struct进行参数赋值
  - pref: 方法 `ArgsDress` 实现对 isdata 的解析赋值支持
- **cli/evolve**
  - fix: 修复命令默认（index）不定义是为空

- **cli/repl**
  - feat: 新增命令交互模式处理包（合并来自v1.4.1）
- **str**
  - feat: 新增函数 Str.ParseUnicode 用于解析Unicode（合并来自v1.4.1）。
  - pref: 字符串代码机构化调整，将 RandString 移动到单独文件中
- **rock**
  - feat: 新增函数 ListNoRepeat 来自对 v1.4.1 版本的合并
  - feat: 新增函数 `ListGetOr` 使用通过索引获取数组元素的值或使用默认值
  - feat: 新增函数 `FormatTable` 其是对 1.4.1 版本中的重构
- **culture/pinyin**
  - feat: 合并 v1.4.1 版本程序使其实现拼音搜索
- **culture/pinyin/material**
  - chore：mt_pinyin.txt 升级 0.12.0 --> 0.14.0
- **cmd/uymas2**
  - feat: 新增 pinyin, cal，ganz，hash，digit，datediff，base64，scan，repl 命令，来自对 v1.4.1 版本的程序合并以及处理
  - pref: 项目结构优化，cmd 命名为 `cmdX`
- **internal/recipe**
  - feat: 新增内部私有包，常用于内部 cmd或者实验性包等
- **util/tm**
  - feat: 新增函数 ParseDuration 来自对 v1.4.1 版本的合并
- **example/cli/struct**
  - feat: 新增 struct 示例（测试）包，用于实现单struct注册




### v2.0.0-rc.1/2024-09-15

> cli 优化，并整合 cli/evolve 减少冗余。

- doc: example 下加入说明文档，以及godoc文档注释完善
- **cli**
  - feat: 实验性地添加 Register 用于处理 Cli 以及 evolve.Evolve 之间的重复处理
  - feat: 新增 `CommandOptional.OffValid` 、`NoValid()`选项用于配置具体命令的验证开关控制
  - pref: 完整 option 有效性验证，添加对不运行选项的验证支持
  - pref: 优化 hook 使其在 index/help命令中同样有效，命令提示日志更新
  - pref: 删除过渡中间类型 `registerAttr[T any]` 
- **cli/evolve**
  - pref: 删除过渡中间类型 `registerEvolveAttr[T any]` 
- **cli/gen**
  - feat: 新增方法 `ArgsDecomposeMust` 用于解析 struct 为选项
  - pref: 优化 `ArgsDress` 函数使其准确无误地支持基础类型的转换

- **fs**
  - fixed: 修复函数 `RootPath()`，空变量时为不以"/"结尾的标准目录
- **logger**
  - pref!: Logger.Format 函数日志内容为空时不输出
- **logger/lgr**
  - feat: 新增方法 `ErrorIf` 用于调试错误，当出现错误时

- **data/convert**
  - feat: 新增函数 `ToSlice` 及 `IsSlice`用于处理字符串数字切片或判别是否符号
  - pref: `SetByStrSlice` 函数实现字符串转任意 slice类型



### v2.0.0-alpha.4/2024-09-05

> cli 优化，并整合 cli/evolve 减少冗余。因原 rc-1 计划版本发生非兼容新版本临时发布 a4 标签

- **cli/gen**
  - pref: ArgsDress 选项解析时，`-` 标识忽略
  - pref: 注册时加入简单的重复性检测（待完善）
  - fix: 修复 ArgsDress 解析命令选项不全
- **cli/evolve**
  - remove(break): 移除类型 Param，使用 cli.ArgsParser 代替以简化
- **rock**
  - feat: 新增函数 ListEq、ListSubset等判断数组是否包含或相等
- **util/fs**
  - feat: 新增函数 RunDir 用于获取运行目录，优先工作目录、其次所在目录
  - fix: 修复 RootPath 函数 `joins ...string` 参数无效的问题




### v2.0.0-alpha.3/2024-09-01

> 移除 bin 包，并优化 cli 命令行解析

- del: 移除包 bin 以及，cmd 下原 uymas 以及uymasDemo等应用
- **cli**
  - feat: Application 新增 CommandList 方法用于注册组数（含别名）
  - feat: 实现帮助信息自动生成，包括命令行描述和选项描述
  - feat: 实现命令行参数选项可配置的验证
  - feat: Args 读取参数是支持注册命令时选项默认值
  - feat: 新增函数 HelpSub 实现对子命令的帮助信息的注册
  - feat: 新增方法 CommandOptional.SubCommand 用于获取子集 CommandOptional，实现递归
  - feat: 新增方法 Application.RunArgs 实现传入 ArgParse 并执行命令行
  - feat: 新增方法 ArgParse.Raw 用于获取原始输入
  - break: 调整 Fn 为 func(ArgsParser)，固定格式使其更统一
  - break: Application.Command 注册命令时固定为单个参数
  - break: 重构 Cli 结构体，使其支持传统帮助信息等
- **cli/evolve**
  - feat: 添加对文档自动生成，选项自动验证的支持
  - pref: 根据 cli.Application 的调整实现适应性变更
- **cli/gen**
  - feat: 新增包用于实现命令行下数据生成如值转变
  - feat: 新增函数  ArgsDress 实现命令行参数到结构体值的装扮
  - feat: 新增函数 ArgsDecompose 、OptionTagParse实现对struct以及 tag 的解析
- **cli/chest**
  - feat: 新增函数 CmdSearchRun 用于搜索名字并执行它

- ·**str**
  - feat: 新增 QueueMaxLen 函数用于统计数据中长度最大值
  - feat: 新增函数 str.Unescape 用于将转移符号进行处理
- **data/convert**
  - feat: 新增方法 SetByStr、SetByStrSlice实现字符串到其他基础类型的赋值
- **example/cli/evolve**
  - feat: 新增二级命令 test args 用于测试参数自动获取
- **rock**
  - feat: 新增函数 MapKeysString 用于map提取keys为字符串数组
  - feat: 实现 FormatKv 函数逻辑
- **number**
  - feat: 将 util 中原数字相关函数及struct迁移到此（重新整合）
- **logger**
  - feat: 新增函数 Logger.Pref 用于设置日志消息前缀
  - pref: Logger.Format 函数未设置变量时使用非格式化函数




### v2.0.0-alpha.2/2024-08-24

- **app/svn**
  - pref: 将 svn 包重命名
- **app/storage**
  
  - pref: 将 storage 包重命名
- **app/scan**
  
  - feat: 从 fs 包中分离 DirScanner 作为单独包
- **app/calc**
  
  - feat: 将 `str.Calc` 升级为读取的引用包
- **cli/chest**

  - feat: 将原 butil.InputRequire 相关方法迁移到此包
  - feat: 新增 CmdExist/CmdAble 用于判别命令是否存在
- **util**
  - del: 删除 InQue， InQueAny等方法，可使用 rock.ListIndex代替。（此方法与 str 重复提供）
- **util/fs**
  - pref: 将 fs 包重命名为 util/fs
  - break: 移除 FsReaderWriter 接口（原实验性的）
- **util/xsql**

  - pref: 将 xsql 包重命名为此包名
- **util/cloud**
  - pref: 将 netutil 重命名为此包
  
- **util/unit**
  - pref: 将 unit 包命名为此包
  
- **str**

  - del: 删除 InQue， InQueAny等方法，可使用 rock.ListIndex代替
- **rock**

  - feat: 新增方法 ListRemove，由 str.DelQue 泛型化改进而来
  - feat: 新增方法 ListAny，由 str.StrQueueToAny 泛型化改进而来
  - feat: 新增方法 FormatList，由bin.FormatQue 改进而来
- **data/input**

  - feat: 新增方法 Stringer.Bool 用于解析bool数据
  - pref: 方法 Stringer.Int 调整基于 strconv.Atoi
- **logger**
  - feat: 加入对 Trace 级别的支持，并且实现日志颜色码（合并v1.4.1/bf985c）
  - pref: 优化函数 NewLogger 降低 if 语句的层数



### v2.0.0-alpha.1/2024-08-15

- pref!: 将应用由 `gitee.com/conero/uymas` 调整为 `gitee.com/conero/uymas/v2`，使v2与旧版本可并行运行
- pref!: 移除`Deprecated:`标注的代码
- pref!: 调整go最小支持版本为 1.20，使其支持对 window7相关设备的支持
- **cli**
  - feat: 实现基于函数式的最小命令行程序解析
  - feat: 支持自定义帮忙命令
- **cli/evolve**
  - feat: 新增 evolve 包用于表示功能更全的命令解析程序
- **cli/ansi**
  - feat: 在 v1 中的 *bin/color* 基础上重写命令行字体颜色等风格，对原方法进行重写
- **data/input**
  - feat: 初步创建字符串输入解析器
- **rock**
  - pref: 将原 `util/rock` 迁移到 `rock`类
  - feat: 新增方法 `InList` 用于判断值是否存在列表中
  - feat: 新增方式 `ParamIndex` 用于根据索引获取参数
- **rock/constraints**
  - pref: 将原 `util/rock/constraints` 迁移到 `rock/constraints`类
- **str**
  - pref!: 将原如 `Fn(s)` 转化为 `Str(s).Fn()`，使其便于连贯操作
  - feat: 新增方法 UcFirst
- **util/tm**
  - feat: 新增与时间相关的操作处理包
- **uymas**
  - feat: 新增函数 GetBuildInfo 支持  `-ldflags` 参数注入
