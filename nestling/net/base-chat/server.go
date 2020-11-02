package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/nestling/net/base-chat/chat"
	"github.com/conero/uymas/str"
	"log"
	"net"
	"net/url"
	"time"
)

var lgger *chat.Logger

//tcp 服务器
func main() {
	lgger = chat.NewLogger("")
	cli := bin.NewCLI()
	cli.RegisterFunc(chatServer, "server", "sv")
	cli.RegisterEmpty(chatServer)
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("普通的 TCP 服务器，命令如下")
		fmt.Println("  server,sv 启动tcp服务器")
		fmt.Printf("     --port,-P=%v         设置端口号 \r\n", chat.DefChatPort)
		fmt.Printf("     --network,-N=%v     设置网络协议 tcp,tcp4,tcp6,unix，unixpacket \r\n", chat.DefChatNetwork)
	}, "help", "?")
	cli.Run()
}

//连接缓存
type ConnCache struct {
	ConnId   string   //链接ID
	conn     net.Conn //资源链接
	IsOnline bool     //是否在线
	Username string
}

//连接池
var ConnPools map[int64]*ConnCache

//服务器
func chatServer(cc *bin.CliCmd) {
	//端口号 [1-65535]; 一般不使用 0-1023
	port := cc.ArgRaw("port", "P")
	if "" == port {
		port = chat.DefChatPort
	}
	network := cc.ArgRaw("network", "N")
	if "" == network {
		network = chat.DefChatNetwork
	}

	lgger.Info("已启动服务器：网络类型 => %v, 端口号 => %v.\r\n", network, port)
	listener, err := net.Listen(network, fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	var connIdex int64
	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		/*go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)*/
		//go handler(conn)
		connIdex += 1
		go func(index int64) {
			ptl := NewProtocol(conn)
			if ptl.IsValid {
				lgger.Info("用户【%v】接入网络，链接ID【%v】.\r\n", ptl.username, index)
				ConnPools[index] = &ConnCache{
					ConnId:   fmt.Sprintf("%v-%v", connIdex, str.RandStr.SafeStr(20)),
					conn:     conn,
					IsOnline: true,
					Username: ptl.username,
				}
				ptl.Handler()
			}
		}(connIdex)
	}
}

//协议处理
type Protocol struct {
	conn     net.Conn
	IsValid  bool
	username string
}

//信息广播
func (c *Protocol) Broadcast(msg string, adr *chat.Address) {
	for idx, pool := range ConnPools {
		if pool.IsOnline && pool.conn != nil {
			v := &url.Values{}
			v.Add("message", msg)
			v.Add("from_user", c.username)
			_, er := pool.conn.Write([]byte(adr.Send(adr.Action, v)))
			if er != nil {
				lgger.Error("Error, 广播发送连接ID为: %v 发生错误，信息: %v.\r\n",
					pool.ConnId, er.Error())
				pool.IsOnline = false
				pool.conn = nil
				ConnPools[idx] = pool
			} else {
				lgger.Info("广播，数据发送方：%v 的消息成功.\r\n", pool.ConnId)
			}
		}
	}
}

func (c *Protocol) ToUser(username, msg string, adr *chat.Address) {
	for _, pool := range ConnPools {
		if pool.IsOnline && pool.conn != nil && username == pool.Username {
			v := &url.Values{}
			v.Add("message", msg)
			v.Add("from_user", c.username)
			_, er := pool.conn.Write([]byte(adr.Send(adr.Action, v)))
			if er != nil {
				lgger.Error("Error, 广播发送连接ID为: %v 发生错误，信息: %v.\r\n",
					pool.ConnId, er.Error())
			} else {
				lgger.Info("发送到用户：%v 的消息成功", username)
			}
		}
	}
}

//获取请求内容
func (c *Protocol) connect() {
	//限定
	overTime := chat.Timer(time.Second * 60)
	for {
		if overTime() {
			lgger.Error("等待客服端请求超时！")
			break
		}
		addr, er := chat.RespondContent(c.conn)
		if er == nil && addr != nil {
			if "native-message" == addr.Protocol && addr.URL != nil && "authorization" == addr.Action {
				username := addr.URL.Query().Get("username")
				if username != "" {
					c.username = username
					v := &url.Values{}
					v.Add("success", "true")
					v.Add("name", username)
					_, er := c.conn.Write([]byte(addr.Send("authorization", v)))
					if er != nil {
						lgger.Error("返回用户数据错误，信息: %v\r\n", er.Error())
						c.username = ""
					}
					break
				}
			}
		} else if er != nil {
			lgger.Error("服务器认证读取客户端数据出错，信息：%v.\r\n", er.Error())
			break
		}
		fmt.Printf("client>> %v.\r\n", addr.Content)
	}

	if c.username == "" {
		c.conn.Close()
		log.Fatal("客服端请求无效，为获取的用户信息")
	}
}

//请求处理
func (c *Protocol) Handler() {
	for {
		adr, er := chat.RespondContent(c.conn)
		if er != nil {
			lgger.Error("认证完成后读取客服端数据错误，信息：%v\r\n", er.Error())
			break
		}
		u := adr.URL
		var value url.Values = nil
		if u != nil {
			value = u.Query()
		}
		switch adr.Action {
		case "broadcast": //广播
			if value != nil {
				c.Broadcast(value.Get("message"), adr)
			} else {
				lgger.Error("广播信息时数据错误，message 数据不存在")
			}
		case "send-message":
			if value != nil {
				c.ToUser(value.Get("username"), value.Get("message"), adr)
			} else {
				lgger.Error("广播信息时数据错误，username, message 数据不存在")
			}
		}
	}
}

//协议初始化
func NewProtocol(conn net.Conn) *Protocol {
	ptl := &Protocol{conn: conn, IsValid: true}
	ptl.connect()
	return ptl
}

func init() {
	ConnPools = map[int64]*ConnCache{}
}
