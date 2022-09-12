package bin

//数据类型定义，通过结构体定义命令行程序

// CMD
// the struct to be a map for cli application
// @todo need to do
type CMD struct {
	Title   string //the title of `CMD`
	Command string
	Alias   []string //the alias of `CMD`, support many alias command

	//help information
	Describe    string
	HelpMessage string        //help message
	HelpCall    func(*CliCmd) //help message by `Func`

	//the action by the Func
	Todo    func(*CliCmd) //cli callback
	TodoApp any           //bind from App struct
}
