package main

import (
	"github.com/yemingfeng/sdb/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("faild to connect: %+v", err)
	}
	defer conn.Close()

	// 连接服务器
	c := pb.NewSDBClient(conn)
	// 发起 incr 请求
	incrResponse, err := c.Incr(context.Background(),
		&pb.IncrRequest{Key: []byte("abc"), Delta: 10})
	log.Printf("incrResponse: %+v, err: %+v", incrResponse, err)
	// 发起 get 请求
	getResponse, err := c.Get(context.Background(),
		&pb.GetRequest{Key: []byte("abc")})
	log.Printf("getResponse: %+v, err: %+v", getResponse, err)
	// 发起 delete 请求
	delResponse, err := c.Del(context.Background(),
		&pb.DelRequest{Key: []byte("h")})
	log.Printf("delResponse: %+v, err: %+v", delResponse, err)
	// 发起 set 请求
	setResponse, err := c.Set(context.Background(),
		&pb.SetRequest{Key: []byte("h"), Value: []byte("h")})
	log.Printf("setResponse: %+v, err: %+v", setResponse, err)
	// 发起 get 请求
	getResponse, err = c.Get(context.Background(),
		&pb.GetRequest{Key: []byte("h")})
	log.Printf("getResponse: %+v, err: %+v", getResponse, err)
	for i := 0; i < 100; i++ {
		_, _ = c.Set(context.Background(),
			&pb.SetRequest{Key: []byte("h111" + strconv.Itoa(i)), Value: []byte("h" + strconv.Itoa(i))})
	}
}
