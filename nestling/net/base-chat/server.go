package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/nestling/net/base-chat/chat"
	"github.com/conero/uymas/str"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

//tcp 服务器
func main() {
	cli := bin.NewCLI()
	cli.RegisterFunc(baseServer, "server", "sv")
	cli.RegisterEmpty(baseServer)
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("普通的 TCP 服务器，命令如下")
		fmt.Println("  server,sv 启动tcp服务器")
		fmt.Println("     --port,-P=7400       设置端口号 ")
		fmt.Println("     --network,-N=tcp     设置网络协议 tcp,tcp4,tcp6,unix，unixpacket ")
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
func baseServer(cc *bin.CliCmd) {
	//端口号 [1-65535]; 一般不使用 0-1023
	port := cc.ArgRaw("port", "P")
	if "" == port {
		port = "7400"
	}
	network := cc.ArgRaw("network", "N")
	if "" == network {
		network = "tcp"
	}

	log.Printf("已启动服务器：网络类型 => %v, 端口号 => %v.\r\n", network, port)
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
				log.Printf("用户【%v】接入网络，链接ID【%v】.\r\n", ptl.username, index)
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

//获取内容
func (c *Protocol) RequestContent() string {
	conn := c.conn
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("read err:", err)
		return ""
	}
	content := string(buf[0:n])
	content = strings.TrimSpace(content)
	return content
}

//信息广播
func (c *Protocol) Broadcast(msg string, adr *chat.Address) {
	for idx, pool := range ConnPools {
		if pool.IsOnline && pool.conn != nil {
			_, er := pool.conn.Write([]byte(msg))
			if er != nil {
				log.Printf("Error, 广播发送连接ID为: %v 发生错误，信息: %v.\r\n",
					pool.ConnId, er.Error())
				pool.IsOnline = false
				pool.conn = nil
				ConnPools[idx] = pool
			}
		}
	}
}

func (c *Protocol) ToUser(username, msg string, adr *chat.Address) {
	for _, pool := range ConnPools {
		if pool.IsOnline && pool.conn != nil && username == pool.Username {
			_, er := pool.conn.Write([]byte(msg))
			if er != nil {
				log.Printf("Error, 广播发送连接ID为: %v 发生错误，信息: %v.\r\n",
					pool.ConnId, er.Error())
			} else {
				_, er = c.conn.Write([]byte(fmt.Sprintf("发送到用户：%v 的消息成功", username)))
				v := &url.Values{}
				v.Add("message", fmt.Sprintf("发送到用户：%v 的消息成功", username))
				_, er = c.conn.Write([]byte(adr.Send(adr.Action, v)))
				if er != nil {
					log.Printf("客服端信息发送失败，信息: %v\r\n", er.Error())
				}
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
			log.Fatal("等待客服端请求超时！")
		}
		content := c.RequestContent()
		if "" != content {
			addr := chat.ParseAddress(content)
			if "native-message" == addr.Protocol && addr.URL != nil && "authorization" == addr.Action {
				username := addr.URL.Query().Get("username")
				if username != "" {
					c.username = username
					v := &url.Values{}
					v.Add("success", "true")
					v.Add("name", username)
					_, er := c.conn.Write([]byte(addr.Send("authorization", v)))
					if er != nil {
						log.Fatalf("返回用户数据错误，信息: %v\r\n", er.Error())
					}
					break
				}
			}
		}
		log.Printf("client>> %v.\r\n", content)
	}

	if c.username == "" {
		c.conn.Close()
		log.Fatal("客服端请求无效，为获取的用户信息")
	}
}

//请求处理
func (c *Protocol) Handler() {
	for {
		content := c.RequestContent()
		adr := chat.ParseAddress(content)
		u := adr.URL
		var value url.Values = nil
		if u != nil {
			value = u.Query()
		}
		switch adr.Action {
		case "broadcast": //广播
			if value != nil {
				c.Broadcast(value.Get("message"), adr)
			}
		case "send-message":
			if value != nil {
				c.ToUser(value.Get("username"), value.Get("message"), adr)
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

//处理连接
func handler(conn net.Conn) {
	_, err := conn.Write([]byte("您已连接成功！"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("read err:", err)
			return
		}
		content := string(buf[0:n])
		log.Printf("client>>\r\n%v", content)
	}
}

func init() {
	ConnPools = map[int64]*ConnCache{}
}
