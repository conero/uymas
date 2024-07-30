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



### v2.0.0/dev

- [ ] xini 库的测试，使其可用常规使用
  - [x] 支持指令，如导入文件
  - [ ] 是否支持条件，如三元符号或`if-else`
- [ ] bin 重复注册命令式，提供可选的panic。即提前预知错误（错误检测）
- [ ] v2.0 为旧版本v1.x的重大重构版本，包结构等进行重大调整。



- pref!: 将应用由 `gitee.com/conero/uymas` 调整为 `gitee.com/conero/uymas/v2`，使v2与旧版本可并行运行
- pref!: 移除`Deprecated:`标注的代码
- **cli**
  - feat: 命令行包初步搭建
- **data/input**
  - feat: 初步创建字符串输入解析器
- **rock**
  - pref: 将原 `util/rock` 迁移到 `rock`类
- **rock/constraints**
  - pref: 将原 `util/rock/constraints` 迁移到 `rock/constraints`类
