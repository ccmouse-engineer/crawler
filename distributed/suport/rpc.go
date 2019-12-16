// 开启RPC服务器端服务和连接到RPC服务端服务
package suport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// SeverRPC开启RPC服务器端服务
func SeverRPC(host string, service interface{}) error {
	// 服务端注册rpc服务
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	// 监听服务端口
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	// 接收连接并处理服务调用
	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}
		go jsonrpc.ServeConn(accept)
	}
}

// NewClient连接到RPC服务端服务
func NewClient(host string) (*rpc.Client, error) {
	// 连接到RPC服务端
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	// 创建RPC客户端
	return jsonrpc.NewClient(conn), nil
}
