package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas/bin"
	chat2 "github.com/conero/uymas/nestling/net/base-chat/chat"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

var lgger chat2.Logger

//客户端
func main() {
	lgger := chat2.NewLogger("")
	cli := bin.NewCLI()
	//call
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		//端口号 [1-65535]; 一般不使用 0-1023
		port := cc.ArgRaw("port", "P")
		if "" == port {
			port = chat2.DefChatPort
		}
		network := cc.ArgRaw("network", "N")
		if "" == network {
			network = chat2.DefChatNetwork
		}
		host := cc.ArgRaw("host", "H")
		if "" == network {
			host = chat2.DefChatHost
		}

		lgger.Info("正在连接……，网络类型 => %v, 端口号 => %v.\r\n", network, port)
		repl(network, port, host)
	}, "call", "c")
	//默认启动请求
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		repl(chat2.DefChatPort, chat2.DefChatNetwork, chat2.DefChatHost)
	})
	//默认服务
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("普通的 TCP 服务器，命令如下")
		fmt.Println("  call,cc 请求tcp服务器")
		fmt.Println("     --port,-P=7400       设置端口号 ")
		fmt.Println("     --network,-N=tcp     设置网络协议 tcp,tcp4,tcp6,unix，unixpacket ")
	}, "help", "?")
	cli.Run()
}

//交互式网络请求
func repl(port, network, host string) {
	hostAdrr := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.Dial(network, hostAdrr)
	if err != nil {
		log.Fatal(err)
	}
	lgger.Info("服务（%v）已连接成功.\r\n", hostAdrr)
	defer conn.Close()
	username := input("请输入您的姓名：", true)
	pwsd := input("请输入密码（123456）：", true)
	if pwsd != "123456" {
		lgger.Error("您的用户密码错误！！")
		return
	}

	NewChat(conn, username)
	/*inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("client-me>> ")
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			return
		}
	}*/
}

func input(tip string, required bool) string {
	fmt.Print(tip)
	inputReader := bufio.NewReader(os.Stdin)
	var input string
	for {
		input, _ = inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if required && "" == input {
			fmt.Println("  数据不可为空！")
			fmt.Print(tip)
			continue
		}
		break
	}
	return input
}

//聊天器
type Chat struct {
	conn net.Conn
}

//构造函数
func NewChat(conn net.Conn, username string) *Chat {
	chat := &Chat{
		conn: conn,
	}
	addr := chat2.Address{
		Protocol: "native-message",
		Action:   "authorization",
	}
	uV := &url.Values{}
	uV.Add("username", username)
	err := addr.SendValue(conn, uV)
	if err != nil {
		lgger.Error("用户请求认证错误，Error: %v.", err)
		return nil
	}

	timer := chat2.Timer(60 * time.Second)
	for {
		if timer() {
			lgger.Fatal("等待服务响应超时")
		}
		addr, err := chat2.RespondContent(conn)
		if err != nil {
			lgger.Fatal("读取服务器数据错误，信息：%v", lgger)
		}
		if addr != nil {
			log.Printf("server>>\r\n%v", addr.Content)
			if addr.Action == "authorization" {
				v := addr.URL.Query()
				if "true" == v.Get("success") {
					lgger.Info("您成功登录服务器！")
					break
				}
			}
		} else {
			lgger.Error("服务器接收的数据为空！")
		}
	}

	cCli := bin.NewCLI()

	//help
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println(" exit              退户程序")
		fmt.Println(" broadcast,bc      广播信息")
		fmt.Println(" user,us <user>    与用户进行聊天")
	}, "help", "?")

	//bc
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		//命令执行
		var input = bufio.NewScanner(os.Stdin)
		fmt.Println("驻留式命令行程序")
		fmt.Printf("$ broadcast/%v>", username)
		for input.Scan() {
			text := input.Text()
			text = strings.TrimSpace(text)
			switch text {
			case "exit":
				fmt.Println("您将退出系统！")
				break
			default:
				if "" != text {
					uV := &url.Values{}
					uV.Add("message", text)
					send := chat2.Address{
						Protocol: "native-message",
						Action:   "broadcast",
					}
					er := send.SendValue(conn, uV)
					if er != nil {
						lgger.Error("发送广播服务错误！信息: %v.\r\n", er.Error())
					} else {
						lgger.Info("广播信息已发送")
					}
				}
			}
			fmt.Println()
			fmt.Printf("$ broadcast/%v>", username)
		}
	}, "bc", "broadcast")

	//user
	cCli.RegisterFunc(func(cc *bin.CliCmd) {
		toUser := cc.SubCommand
		if toUser != "" {
			var input = bufio.NewScanner(os.Stdin)
			fmt.Println("驻留式命令行程序")
			fmt.Printf("$ %v>", username)
			for input.Scan() {
				text := input.Text()
				text = strings.TrimSpace(text)

				switch text {
				case "exit":
					fmt.Println("您将退出系统！")
					break
				default:
					if "" != text {
						uV := &url.Values{}
						uV.Add("message", text)
						uV.Add("username", toUser)
						send := chat2.Address{
							Protocol: "native-message",
							Action:   "send-message",
						}
						er := send.SendValue(conn, uV)
						if er != nil {
							lgger.Error("发送信息到用户(%v)，失败！错误：%v.", toUser, er.Error())
						} else {
							lgger.Info("信息已发送")
						}
					}
				}
				fmt.Println()
				fmt.Printf("$ %v>", username)
			}
		}
	}, "user", "us")

	//命令执行
	var input = bufio.NewScanner(os.Stdin)
	fmt.Println("驻留式命令行程序")
	fmt.Printf("$ %v>", username)
	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		switch text {
		case "exit":
			fmt.Println("您将退出系统！")
			os.Exit(0)
		default:
			cCli.Run(strings.Split(text, " ")...)
		}

		fmt.Println()
		fmt.Printf("$ %v>", username)
	}
	return chat
}
