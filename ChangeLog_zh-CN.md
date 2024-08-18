# 更新日志
> 2018年10月30日 星期二
> 项目



**版本介绍**

- **x.y.z**     	保证的兼容性，可新增功能（用于版本阶段性开发）、修复或调整版本。（待移除使用 `// Deprecated:  descript text` 标记，或说明）
- **x.y**            不保证的兼容性，删除旧版本遗弃的方法
- **x**               重大（颠覆性）的改变，重要里程碑开发



## future

- **bin**
  - [ ] 是否增加可选的 *cache* 选项，实现对命令行程序缓存，增加程序响应。
- [ ] 包支持 1.18并使用范类重写方法。(maybe v2.0.0）



### todo

- [ ] xini 库的测试，使其可用常规使用
  - [x] 支持指令，如导入文件
  - [ ] 是否支持条件，如三元符号或`if-else`
- [ ] bin 重复注册命令式，提供可选的panic。即提前预知错误（错误检测）
- [ ] v2.0 为旧版本v1.x的重大重构版本，包结构等进行重



### v2.0.x/dev

> v2.0



### v2.0.0-alpha.2/dev

- **app/svn**
  
  - pref: 将 svn 包重命名
  
- **app/storage**
  
  - pref: 将 storage 包重命名
  
- **cli/chest**

  - feat: 将原 butil.InputRequire 相关方法迁移到此包

- **util**
  - del: 删除 InQue， InQueAny等方法，可使用 rock.ListIndex代替。（此方法与 str 重复提供）

- **util/fs**
  - pref: 将 fs 包重命名为 util/fs
  - break: 移除 FsReaderWriter 接口（原实验性的）

- **util/fs/scan**
  - feat: 从 fs 包中分离 DirScanner 作为单独包

- **util/xsql**

  - pref: 将 xsql 包重命名为此包名

- **str**
  - del: 删除 InQue， InQueAny等方法，可使用 rock.ListIndex代替

- **rock**

  - feat: 新增方法 ListRemove，由 str.DelQue 泛型化改进而来

  






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
