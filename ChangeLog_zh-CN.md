# 更新日志
> 2018年10月30日 星期二
> 项目



**版本介绍**

- **x.y.z**     	保证的兼容性，可新增功能（用于版本阶段性开发）、修复或调整版本。（待移除使用 `// Deprecated:  descript text` 标记，或说明）
- **x.y**            不保证的兼容性，删除旧版本遗弃的方法
- **x**               重大（颠覆性）的改变，重要里程碑开发



## future

> v2.0.0 开发实现中



### todo

- [x] v2.0 为旧版本v1.x的重大重构版本，包结构等进行重
- [ ] doc 自动文档设计，计划在 v2.0.0-alpha.3 中实现 `#240824`



### v2.0.x/dev

> v2.0



### v2.0.0-alpha.3/dev

> 移除 bin 包，并优化 cli 命令行解析

- del: 移除包 bin 以及，cmd 下原 uymas 以及uymasDemo等应用
- **cli**
  - feat: Application 新增 CommandList 方法用于注册组数（含别名）
  - break: 调整 Fn 为 func(ArgsParser)，固定格式使其更统一
  - break: Application.Command 注册命令时固定为单个参数



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
