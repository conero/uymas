package cli

var (
	glbArgs ArgsParser
)

// OsArgs 获取基于系统命令行的参数解析，基于全局变量
func OsArgs() ArgsParser {
	if glbArgs == nil {
		glbArgs = NewArgs()
	}
	return glbArgs
}
