/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月28日 星期六
 * @ini 变量列表
 */

package xini

const (
	SupportNameIni  = "ini"
	SupportNameRong = "rong"
	SupportNameToml = "toml"
)

const (
	Author      = "Joshua Conero" // author 作者
	Name        = "conero/ini"
	Version     = "2.1.0-alpha.4"         // version	版本号
	Release     = "20190617"              // build 发布日期
	Description = "ini parser for golang" // name 名称
	Since       = "20171028"              // start 开始时间
	Copyright   = "@Conero"               // copyright 版权
)

// IniParseSettings ini-parse set base
var IniParseSettings map[string]string = map[string]string{
	"equal":          "=",                    // 等号符
	"comment":        "#|;",                  // 注释符号
	"mcomment1":      "'''",                  // 多行注释 - 开始
	"mcomment2":      "'''",                  // 多行注释 - 结束
	"limiter":        ",",                    // 分隔符
	"scope1":         "{",                    // 作用域 - 开始
	"scope2":         "}",                    // 作用域 - 结束
	"reg_comment":    "^[#;]",                // 注释符号
	"reg_section":    "^\\[[^\\[^\\]}]*\\]$", // 是否为章节正则检测
	"reg_section_sg": "(\\[)|(\\])",          // 章节标点符号处理
	"reg_scope":      "\\{[^\\{^\\}]*\\}",    // 作用域开始于结束正则
	//"reg_scope_sg": "$\\{[^\\{^\\}]*\\}^", // 单行作用域解析
	"reg_scope_sg":    "^\\{.*\\}$", // 单行作用域解析
	"mlstring":        `"|'`,        // 多行字符串
	"reg_clear_mls":   `"|'`,        // 清除多行字符串中的字符
	"reg_has_comment": `#|;`,        // 注释二进制
	//"reg_is_mlstring":  `^[A-Za-z0-9_-]+[=\s]+("|').*[^"^']+$`, //	是多行字符正则开始,  否 key = "ttt" 是 key = " 888
	"reg_is_mlstring":    `^[\w]+[=\s]+("|'){1}[^"']*$`, //	是多行字符正则开始,  否 key = "ttt" 是 key = " 888
	"reg_is_mlstring_nk": `^["']{1}[^"'\,]+$`,           //	是多行字符正则开始（无键值 no key）  "|'
	"reg_mlstring_sta":   `^['"].*`,                     //	多行字符正则开始
	//"reg_mlstring_end": `[^=]*['"]+$`,                          // 多行字符正则结束
	//"reg_mlstring_end": `^[^"'=]*['"]{1}$`, 				//	多行字符正则结束 不支持分隔符
	"reg_mlstring_end": `^[^"'=]*['"\,]+$`, // 多行字符正则结束 支持分隔符
}

// TranStrMap transfer character parsing
var TranStrMap map[string]string = map[string]string{
	`\,`: "_JC__COMMA", // 逗号转移符
	`\{`: "_L__BRACE",  // 左大括弧号
	`\}`: "_R__BRACE",  // 右大括弧号
	`\=`: "_JC__EQUAL", // 等于符号转移替代
}

// TranCommentMap comment character parsing
var TranCommentMap map[string]string = map[string]string{
	`\;`: "_JC__COMMIT1",
	`\#`: "_JC__COMMIT2",
}
