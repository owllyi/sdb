package main

import (
	"fmt"
	"github.com/yemingfeng/sdb/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("faild to connect: %+v", err)
	}
	defer conn.Close()

	// 连接服务器
	c := pb.NewSDBClient(conn)
	// 发起 spush 请求
	values := make([][]byte, 100)
	for i := 0; i < 100; i++ {
		values[i] = []byte("h" + fmt.Sprint(i))
		values[i] = []byte("h" + fmt.Sprint(i))
	}
	spushResponse, err := c.SPush(context.Background(),
		&pb.SPushRequest{Key: []byte("h"), Values: values})
	log.Printf("spushResponse: %+v, err: %+v", spushResponse, err)

	// 发起 spop 请求
	values = make([][]byte, 50)
	for i := 0; i < 50; i++ {
		values[i] = []byte("h" + fmt.Sprint(i*2))
	}
	spopResponse, err := c.SPop(context.Background(),
		&pb.SPopRequest{Key: []byte("h"), Values: values})
	log.Printf("spopResponse: %+v, err: %+v", spopResponse, err)

	// 发起 sexist 请求
	sexistResponse, err := c.SExist(context.Background(),
		&pb.SExistRequest{Key: []byte("h"),
			Values: [][]byte{[]byte("h1"), []byte("h2"), []byte("h3"), []byte("h4"), []byte("h5")}})
	log.Printf("sexistResponse: %+v, err: %+v", sexistResponse, err)

	// 发起 scount 请求
	scountResponse, err := c.SCount(context.Background(),
		&pb.SCountRequest{Key: []byte("h")})
	log.Printf("scountResponse: %+v, err: %+v", scountResponse, err)

	// 发起 sdel 请求
	sdelResponse, err := c.SDel(context.Background(),
		&pb.SDelRequest{Key: []byte("h")})
	log.Printf("sdelResponse: %+v, err: %+v", sdelResponse, err)
}
