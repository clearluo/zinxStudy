package znet

import (
	"fmt"
	"net"
	"zinxStudy/zinx/ziface"
)

// IServer接口的实现，定义一个Server的服务模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的ip
	IP string
	// 服务器舰艇的端口
	Port int
}

// 启动服务
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port:%d, is staring\n", s.IP, s.Port)

	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}

		// 2 监听服务器的地址
		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, err)
			return
		}
		fmt.Println("start Zinx server succ, ", s.Name, " succ, Listenning...")

		// 3 阻塞的等待客户端链接，处理客户端链接业务（读写）
		for {
			// 如果有客户端链接过来，阻塞会会返回
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			// 已经与客户端建立链接，做一些业务，做一个最基本的最大512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err:", err)
						continue
					}
					fmt.Printf("recv client buf:%s, cnt:%d\n", buf[:cnt], cnt)
					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err:", err)
						continue
					}
				}
			}()
		}
	}()

}

// 停止服务
func (s *Server) Stop() {
	// TODO 将服务器的一些资源，状态或者一些已经开辟的资源回收
}

// 运行服务
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务之后的额外业务

	// 阻塞状态
	select {}
}

//
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
