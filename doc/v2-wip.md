## V2-WIP

> 2024年3月16日 星期六
>
> Joshua Conero







程序编译

```powershell
# 将系统所有包编译到 dist用于编译尺寸比较
go build -o ./dist ./...

# 压缩打包
go build -ldflags "-w -s" -o ./dist/mini ./...

# 使用 tinygo 打包
ls ./cmd/ | foreach{tinygo build -o "./dist/tinygo/$($_.Name).exe" $_}
ls ./example/cli | foreach{tinygo build -o "./dist/tinygo/$($_.Name).exe" $_}
```





### Roadmap

- [x] v2.0.0-alpha.1 发布实现，初步实现主要库的转移
- [x] v2.0.0-alpha.2 移除原包中遗弃的方法
- [x] v2.0.0-alpha.3 cli 下命令行工具功能实现和完善（功能初步稳定）



### Todo

- [ ] 所有库的名称都有可能进行调整 `#240316`
- [ ] 二进制基础库
  - [ ] 命令行解析
    - [ ] 支持 reflect 解析字符串到，不同的模板字符串
    - [ ] 支持 泛型 实现字符串或不同参数的解析



#### 包结构设计

- **app**                  程序应用（应用级别）
- **cli**                     二进制命令行，简单二进制生成，简单可用。没有复杂功能，适合如 tinygo 等编译
  - **evolve**            命令行演化程序，功能更全。支持func/struct
  - **chest**              命令行常用工具
  - **ansi**                命令行颜色码
- **rock**          基础库，基于泛型实现基础的类型
- **fs**                文件系统
- **logger**        日志系统
- **data**            数据处理，基础数据处理。可提供基础的非反射处理
  - **convert**        数据转换，使用 reflect 做数据转换
  - **list**                 列表数据集合
- **culture**      中国文化相关处理包
- **cmd**            当前系统应用包程序
- **util**
  - **tm**             时间助手包，以时间的相关操作为主





simple 与 full 的api尽可能保持一致，可用在cli中声明interface接口。



> **evolve**

- [ ] 根据 struct 的tag解析，生成相关的 doc文档。并将文档缓存为文件，及一般编译后option等是不会改变的





#### reflect

基础类型

- Int
  - int
  - int8
  - int16
  - int32
  - int64
- Uint
  - uint
  - uint8
  - uint16
  - uint32
  - uint64
- Float
  - float32
  - float64
- Bool



> 复杂类型

- struct
  - 通过 TagName 或字段名进行映射
- map
- array



#### 命令行

- [ ] Args 解析，数据
  - [x] 解析输入的字符串为任意用户需要的类型，原始为字符串。`字符串  --> type`
- [x] Option 选项



doc 自动文档设计，使其支持单页动态搜索

```yaml
# 一级
- title
- option       description text,     require,  default

# 支持分组
- group    
  - command      description text
    - option     description text,     require,  default
    - ...
    - subCommand      description text
      - option     description text,     require,  default
      - ...

# 命令以及选项文档
- command      description text
  - option     description text,     require,  default
  - ...
  - subCommand      description text
    - option     description text,     require,  default
    - ...
```



##### 帮助命令

支持帮助命令

```shell
# 全局命令行
$ --help
$ -h

# 查看 command 命令的帮助信息
$ --help command
# 查看 command 命令的帮助信息
$ -h command

#
$ command --help
$ command -h

# ?
$ command --help sub-command
$ command -h sub-command

# ?
$ help sub-command
$ ? sub-command
```

