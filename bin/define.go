package bin

//数据类型定义，通过结构体定义命令行程序

//@todo 需实现
//映射为具体的命令
type CMD struct {
	Title   string   //标题
	Command string   //命令行
	Alias   []string //别名

	//帮助信息
	Describe    string           //描述
	HelpMessage string           //帮助信息
	HelpCall    func(cc *CliCmd) //帮助信息回调

	//回调，可以默认为当前本身
	Todo    func(cc *CliCmd) //命令回调
	TodoApp interface{}      //命令绑定信息
}
