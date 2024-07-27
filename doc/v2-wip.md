## V2-WIP

> 2024年3月16日 星期六
>
> Joshua Conero





### Todo

- [ ] 所有库的名称都有可能进行调整 `#240316`
- [ ] 二进制基础库
  - [ ] 命令行解析
    - [ ] 支持 reflect 解析字符串到，不同的模板字符串
    - [ ] 支持 泛型 实现字符串或不同参数的解析



#### 包结构设计

- **cli**                     二进制命令行
  - **simple**      简单二进制生成，简单可用。没有复杂功能，适合如 tinygo 等编译
  - **full**            全局工具，更全。
- **rock**          基础库，基于泛型实现基础的类型
- **fs**                文件系统
- **logger**        日志系统
- **data**            数据处理，基础数据处理。可提供基础的非反射处理
  - **convert**        数据转换，使用 reflect 做数据转换
  - **list**                 列表数据集合

- **culture**      中国文化相关处理包
- **cmd**            当前系统应用包程序




simple 与 full 的api尽可能保持一致，可用在cli中声明interface接口。



> **full**

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
  - [ ] 解析输入的字符串为任意用户需要的类型，原始为字符串。`字符串  --> type`
- [ ] Option 选项

