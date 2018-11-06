# 更新日志
> 2018年10月30日 星期二
> 项目

## 0.2.x

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